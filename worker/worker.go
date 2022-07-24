package worker

import (
	"blion-worker-consenso/internal/env"
	"blion-worker-consenso/internal/grpc/accounting_proto"
	"blion-worker-consenso/internal/grpc/auth_proto"
	"blion-worker-consenso/internal/grpc/mine_proto"
	"blion-worker-consenso/internal/grpc/wallet_proto"
	"blion-worker-consenso/internal/logger"
	"blion-worker-consenso/pkg/bk"
	"blion-worker-consenso/pkg/bk/lotteries"
	"blion-worker-consenso/pkg/bk/miner_response"
	"blion-worker-consenso/pkg/bk/participants"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpcMetadata "google.golang.org/grpc/metadata"
	"math/rand"
	"time"
)

type Worker struct {
	Srv *bk.Server
}

func NewWorker(srv *bk.Server) IWorker {
	return &Worker{Srv: srv}
}

func (w Worker) Execute() {
	c := cron.New(cron.WithSeconds())
	cfg := env.NewConfiguration()
	_, err := c.AddFunc(cfg.App.TimerInterval, func() { w.doWork() })
	if err != nil {
		logger.Error.Println("error execute: ", err)
	}
	c.Start()
	defer c.Stop()
	select {}
}

func (w Worker) doWork() {

	e := env.NewConfiguration()

	connBk, err := grpc.Dial(e.BlockService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error.Printf("error conectando con el servicio auth de blockchain: %s", err)
		return
	}
	defer connBk.Close()

	connAuth, err := grpc.Dial(e.AuthService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error.Printf("error conectando con el servicio auth de blockchain: %s", err)
		return
	}
	defer connAuth.Close()

	clientMine := mine_proto.NewMineBlockServicesBlocksClient(connBk)
	clientAuth := auth_proto.NewAuthServicesUsersClient(connAuth)
	clientAccount := accounting_proto.NewAccountingServicesAccountingClient(connAuth)
	clientWallet := wallet_proto.NewWalletServicesWalletClient(connAuth)

	token, err := login(clientAuth)
	if err != nil {
		logger.Error.Printf("No se pudo obtener el token de autorización, error: %v", err)
		return
	}

	ctx := grpcMetadata.AppendToOutgoingContext(context.Background(), "authorization", token)

	block, err := getBlockToMine(clientMine, ctx)
	if err != nil {
		logger.Error.Printf("No se pudo obtener el bloque a minar, error: %v", err)
		return
	}

	lotteryActive, _, err := w.Srv.SrvLottery.GetLotteryActive()
	if err != nil {
		logger.Error.Printf("error trayendo una loteria activa: %s", err)
		return
	}

	if lotteryActive != nil {
		w.isLotteryActive(lotteryActive)
		return
	}

	lotteryMined, _, err := w.Srv.SrvLottery.GetLotteryActiveForMined()
	if err != nil {
		logger.Error.Printf("error trayendo una loteria lista para minar: %s", err)
		return
	}

	if lotteryMined != nil {
		resHash, _, err := w.Srv.SrvMinerResponse.GetMinerResponseRegister(lotteryMined.ID)
		if err != nil {
			logger.Error.Printf("error trayendo el hash del minero: %s", err)
			return
		}
		votesInFavor := 0
		votes, err := w.Srv.SrvValidatorsVote.GetAllValidatorVotesByLotteryID(lotteryMined.ID)
		if err != nil {
			logger.Error.Printf("error trayendo los votos de los validadores: %s", err)
			return
		}

		if len(votes) != e.App.MaxValidator {
			logger.Error.Printf("Se requiere la totalidad de los votos designado en la loteria")
			return
		}

		for _, vote := range votes {
			if vote.Vote {
				votesInFavor++
			}
		}

		participantsLottery, _, err := w.Srv.SrvParticipants.GetParticipantsByLotteryID(lotteryMined.ID)
		if err != nil {
			logger.Error.Printf("error trayendo los participantes de la loteria: %s", err)
			return
		}

		if (votesInFavor*100)/len(votes) > 51 {
			w.isApproved(clientAccount, clientMine, ctx, block, resHash, clientWallet, lotteryMined, participantsLottery)
			return
		}

		w.isNotApproved(resHash, lotteryMined, participantsLottery, ctx, clientWallet)
		return
	}

	_, _, err = w.Srv.SrvLottery.CreateLottery(uuid.New().String(), block.Id, time.Now(), nil, nil, nil, nil, 25)
	if err != nil {
		logger.Error.Printf("No se pudo obtener el bloque a minar, error: %s", err)
	}

	return
}

func shuffle(data []string, maxParticipants int) []string {
	lenSlice := maxParticipants
	if len(data) < maxParticipants {
		lenSlice = len(data)
	}
	ret := make([]string, lenSlice)
	idxs := rand.Perm(len(data))
	for i := 0; i < lenSlice; i++ {
		ret[i] = data[idxs[i]]
	}
	return ret
}

func contains(data []string, value string) bool {
	for _, v := range data {
		if v == value {
			return true
		}
	}

	return false
}

func findOneParticipant(participants []*participants.Participants, value string) *participants.Participants {
	for _, participant := range participants {
		if participant.ID == value {
			return participant
		}
	}
	return nil
}

func getBlockToMine(clientMine mine_proto.MineBlockServicesBlocksClient, ctx context.Context) (*mine_proto.DataBlockMine, error) {
	resBkMine, err := clientMine.GetBlockToMine(ctx, &mine_proto.GetBlockToMineRequest{})
	if err != nil {
		return nil, err
	}

	if resBkMine == nil {
		return nil, fmt.Errorf("no se pudo obtener el bloque a minar")
	}

	if resBkMine.Error {
		return nil, fmt.Errorf(resBkMine.Msg)
	}

	if resBkMine.Data == nil {
		return nil, fmt.Errorf(resBkMine.Msg)
	}

	return resBkMine.Data, nil
}

func login(clientAuth auth_proto.AuthServicesUsersClient) (string, error) {
	e := env.NewConfiguration()
	resLogin, err := clientAuth.Login(context.Background(), &auth_proto.LoginRequest{
		Email:    nil,
		Nickname: &e.App.UserLogin,
		Password: e.App.UserPassword,
	})
	if err != nil {
		logger.Error.Printf("No se pudo obtener el token de autenticacion: %s", err)
		return "", err
	}

	if resLogin == nil {
		logger.Error.Printf("No se pudo obtener el token de autenticacion")
		return "", fmt.Errorf("no se pudo obtener el token de autenticacion")
	}

	if resLogin.Error {
		logger.Error.Printf(resLogin.Msg)
		return "", fmt.Errorf(resLogin.Msg)
	}

	return resLogin.Data.AccessToken, nil
}

func (w Worker) isLotteryActive(lotteryActive *lotteries.Lottery) {
	e := env.NewConfiguration()
	lifeLottery := time.Now().Sub(lotteryActive.RegistrationStartDate).Seconds()
	if int(lifeLottery) > (e.App.SubscriptionTime * 1000) {
		endDate := time.Now()
		lotteryUpdated, _, err := w.Srv.SrvLottery.UpdateLottery(lotteryActive.ID, lotteryActive.BlockId, lotteryActive.RegistrationStartDate, &endDate, &endDate, nil, nil, 26)
		if err != nil {
			logger.Error.Printf("error iniciando la loteria: %s", err)
			return
		}

		participantsActive, _, err := w.Srv.SrvParticipants.GetParticipantsByLotteryID(lotteryActive.ID)
		if err != nil {
			logger.Error.Printf("error trayendo los participantes: %s", err)
			return
		}

		participantsLottery := make([]string, len(participantsActive))

		for i, participant := range participantsActive {
			participantsLottery[i] = participant.ID
		}

		rand.Shuffle(len(participantsLottery), func(i, j int) {
			participantsLottery[i], participantsLottery[j] = participantsLottery[j], participantsLottery[i]
		})

		miners := shuffle(participantsLottery, e.App.MaxMiners)

		newParticipants := make([]string, len(participantsLottery)-e.App.MaxMiners)
		counter := 0
		for _, participant := range participantsLottery {
			if !contains(miners, participant) {
				newParticipants[counter] = participant
				counter++
			}
		}

		rand.Shuffle(len(newParticipants), func(i, j int) {
			newParticipants[i], newParticipants[j] = newParticipants[j], newParticipants[i]
		})

		validators := shuffle(newParticipants, e.App.MaxValidator)

		for _, participant := range participantsActive {
			if contains(miners, participant.ID) {
				_, _, err = w.Srv.SrvParticipants.UpdateParticipants(participant.ID, lotteryActive.ID, participant.WalletId, participant.Amount, true, 23, false)
				if err != nil {
					logger.Error.Printf("error actulizando el registro del participante: %s", err)
					return
				}
			} else if contains(validators, participant.ID) {
				_, _, err = w.Srv.SrvParticipants.UpdateParticipants(participant.ID, lotteryActive.ID, participant.WalletId, participant.Amount, true, 24, false)
				if err != nil {
					logger.Error.Printf("error actulizando el registro del participante: %s", err)
					return
				}
			}
		}

		endDate = time.Now()
		_, _, err = w.Srv.SrvLottery.UpdateLottery(lotteryUpdated.ID, lotteryUpdated.BlockId, lotteryUpdated.RegistrationStartDate, lotteryUpdated.RegistrationEndDate, lotteryUpdated.LotteryStartDate, &endDate, nil, 27)
		if err != nil {
			logger.Error.Printf("error iniciando la loteria: %s", err)
			return
		}
	}
	return
}

func (w Worker) isApproved(clientAccount accounting_proto.AccountingServicesAccountingClient, clientMine mine_proto.MineBlockServicesBlocksClient, ctx context.Context, block *mine_proto.DataBlockMine, resHash *miner_response.MinerResponse, clientWallet wallet_proto.WalletServicesWalletClient, lotteryMined *lotteries.Lottery, participantsLottery []*participants.Participants) {

	e := env.NewConfiguration()
	resMine, err := clientMine.MineBlock(ctx, &mine_proto.RequestMineBlock{
		Id:         block.Id,
		Hash:       resHash.Hash,
		Nonce:      resHash.Nonce,
		Difficulty: int32(resHash.Difficulty),
	})
	if err != nil {
		logger.Error.Printf("error agregando el bloque temporal a la cadena de bloques: %s", err)
		return
	}

	if resMine == nil {
		logger.Error.Printf("error agregando el bloque temporal a la cadena de bloques")
		return
	}

	if resMine.Error {
		logger.Error.Printf(resMine.Msg)
		return
	}

	_, _, err = w.Srv.SrvMinerResponse.UpdateMinerResponse(resHash.ID, lotteryMined.ID, resHash.ParticipantsId, resHash.Hash, 31, resHash.Nonce, resHash.Difficulty)
	if err != nil {
		logger.Error.Printf("error actualizando el hash del minero como aceptado: %s", err)
		return
	}
	endProcess := time.Now()
	_, _, err = w.Srv.SrvLottery.UpdateLottery(lotteryMined.ID, lotteryMined.BlockId, lotteryMined.RegistrationStartDate, lotteryMined.RegistrationEndDate, lotteryMined.LotteryStartDate, lotteryMined.LotteryEndDate, &endProcess, 28)
	if err != nil {
		logger.Error.Printf("error actualizando la loteria: %s", err)
		return
	}

	feeBlock, _, err := w.Srv.SrvBlockFee.GetBlockFeeByBlockID(block.Id)
	if err != nil {
		logger.Error.Printf("error trayendo el fee del bloque: %v", err)
		return
	}

	feeMiner := (feeBlock.Fee * float64(e.App.FeeMine)) / 100

	feeValidators := (feeBlock.Fee * float64(e.App.FeeValidators)) / 100

	for _, participant := range participantsLottery {
		resUnfreeze, err := clientWallet.UnFreezeMoney(ctx, &wallet_proto.RqUnFreezeMoney{WalletId: participant.WalletId, LotteryId: lotteryMined.ID})
		if err != nil {
			logger.Error.Printf("error descongelando el dinero: %s", err)
			continue
		}

		if resUnfreeze == nil {
			logger.Error.Printf("error descongelando el dinero")
			continue
		}

		if resUnfreeze.Error {
			logger.Error.Printf(resUnfreeze.Msg)
			continue
		}

		_, _, err = w.Srv.SrvParticipants.UpdateParticipants(participant.ID, participant.LotteryId, participant.WalletId, participant.Amount, participant.Accepted, participant.TypeCharge, true)
		if err != nil {
			logger.Error.Printf("error actualizando el participante", err)
			continue
		}

		if participant.Accepted && (participant.ID == resHash.ParticipantsId || participant.TypeCharge == 24) {
			resAccount, err := clientAccount.GetAccountingByWalletById(ctx, &accounting_proto.RequestGetAccountingByWalletId{Id: participant.WalletId})
			if err != nil {
				logger.Error.Printf("error obteniendo la cuenta del participante, error: %v", err)
				continue
			}
			if resAccount == nil {
				logger.Error.Printf("error obteniendo la cuenta del participante")
				continue
			}
			if resAccount.Error {
				logger.Error.Printf(resAccount.Msg)
				continue
			}

			amount := feeMiner

			if participant.TypeCharge == 24 {
				amount = feeValidators / float64(e.App.MaxValidator)
			}

			resReward, err := clientAccount.SetAmountToAccounting(ctx, &accounting_proto.RequestSetAmountToAccounting{
				WalletId: participant.WalletId,
				Amount:   resAccount.Data.Amount + amount,
				IdUser:   resAccount.Data.IdUser,
			})
			if err != nil {
				logger.Error.Printf("error actualizando el dinero de la cuenta, error: %v", err)
				continue
			}
			if resReward == nil {
				logger.Error.Printf("error actualizando el dinero de la cuenta")
				continue
			}
			if resReward.Error {
				logger.Error.Printf(resAccount.Msg)
				continue
			}
		}
	}

	return
}

func (w Worker) isNotApproved(resHash *miner_response.MinerResponse, lotteryMined *lotteries.Lottery, participantsLottery []*participants.Participants, ctx context.Context, clientWallet wallet_proto.WalletServicesWalletClient) {
	_, _, err := w.Srv.SrvMinerResponse.UpdateMinerResponse(resHash.ID, lotteryMined.ID, resHash.ParticipantsId, resHash.Hash, 30, resHash.Nonce, resHash.Difficulty)
	if err != nil {
		logger.Error.Printf("error actualizando el hash del minero como aceptado: %s", err)
		return
	}

	endProcess := time.Now()
	_, _, err = w.Srv.SrvLottery.UpdateLottery(lotteryMined.ID, lotteryMined.BlockId, lotteryMined.RegistrationStartDate, lotteryMined.RegistrationEndDate, lotteryMined.LotteryStartDate, lotteryMined.LotteryEndDate, &endProcess, 32)
	if err != nil {
		logger.Error.Printf("error actualizando la loteria: %s", err)
		return
	}

	miner := findOneParticipant(participantsLottery, resHash.ParticipantsId)
	walletPenalties, err := w.Srv.SrvPenaltyParticipant.GetAllPenaltyParticipantsByWalletID(miner.WalletId)
	if err != nil {
		logger.Error.Printf("error trayendo las multas del minero: %s", err)
		return
	}

	var amountPenalty = 0.0
	var penaltyPercentage = 0.0

	resFreezeMoney, err := clientWallet.GetFrozenMoney(ctx, &wallet_proto.RqGetFrozenMoney{WalletId: miner.WalletId})
	if err != nil {
		logger.Error.Printf("error trayendo el dinero congelado: %s", err)
		return
	}

	if resFreezeMoney == nil {
		logger.Error.Printf("error trayendo el dinero congelado")
		return
	}

	if resFreezeMoney.Error {
		logger.Error.Printf(resFreezeMoney.Msg)
		return
	}
	freezeMoney := resFreezeMoney.Data

	if walletPenalties != nil && len(walletPenalties) > 0 {
		if len(walletPenalties) == 1 {
			penaltyPercentage = 25
			amountPenalty = float64((freezeMoney * 25) / 100)
		} else if len(walletPenalties) == 2 {
			penaltyPercentage = 50
			amountPenalty = float64((freezeMoney * 50) / 100)
		} else if len(walletPenalties) >= 3 {
			penaltyPercentage = 100
			amountPenalty = float64(freezeMoney)
		}
	} else {
		amountPenalty = float64(freezeMoney * 10 / 100)
		penaltyPercentage = 10
	}

	_, _, err = w.Srv.SrvPenaltyParticipant.CreatePenaltyParticipants(uuid.New().String(), lotteryMined.ID, resHash.ParticipantsId, amountPenalty, penaltyPercentage)
	if err != nil {
		logger.Error.Printf("error multando al minero por fraude: %v", err)
		return
	}

	resUnfreeze, err := clientWallet.UnFreezeMoney(ctx, &wallet_proto.RqUnFreezeMoney{WalletId: miner.WalletId, LotteryId: lotteryMined.ID, Penalty: amountPenalty})
	if err != nil {
		logger.Error.Printf("error descongelando el dinero: %s", err)
		return
	}

	if resUnfreeze == nil {
		logger.Error.Printf("error descongelando el dinero")
		return
	}

	if resUnfreeze.Error {
		logger.Error.Printf(resUnfreeze.Msg)
		return
	}

	_, _, err = w.Srv.SrvParticipants.UpdateParticipants(miner.ID, miner.LotteryId, miner.WalletId, miner.Amount, miner.Accepted, miner.TypeCharge, true)
	if err != nil {
		logger.Error.Printf("error actualizando el participante", err)
		return
	}

	for _, participant := range participantsLottery {
		if participant.ID != resHash.ParticipantsId {
			resUnfreeze, err := clientWallet.UnFreezeMoney(ctx, &wallet_proto.RqUnFreezeMoney{WalletId: participant.WalletId, LotteryId: lotteryMined.ID, Penalty: 0})
			if err != nil {
				logger.Error.Printf("error descongelando el dinero: %s", err)
				continue
			}

			if resUnfreeze == nil {
				logger.Error.Printf("error descongelando el dinero")
				continue
			}

			if resUnfreeze.Error {
				logger.Error.Printf(resUnfreeze.Msg)
				continue
			}

			_, _, err = w.Srv.SrvParticipants.UpdateParticipants(participant.ID, participant.LotteryId, participant.WalletId, participant.Amount, participant.Accepted, participant.TypeCharge, true)
			if err != nil {
				logger.Error.Printf("error actualizando el participante", err)
				continue
			}
		}
	}
}
