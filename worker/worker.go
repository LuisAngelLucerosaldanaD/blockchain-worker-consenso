package worker

import (
	"blion-worker-consenso/internal/env"
	"blion-worker-consenso/internal/grpc/accounting_proto"
	"blion-worker-consenso/internal/grpc/auth_proto"
	"blion-worker-consenso/internal/grpc/mine_proto"
	"blion-worker-consenso/internal/grpc/transactions_proto"
	"blion-worker-consenso/internal/grpc/wallet_proto"
	"blion-worker-consenso/internal/logger"
	"blion-worker-consenso/pkg/auth"
	"blion-worker-consenso/pkg/bk"
	"blion-worker-consenso/pkg/bk/lotteries"
	"blion-worker-consenso/pkg/bk/miner_response"
	"blion-worker-consenso/pkg/bk/participants"
	"blion-worker-consenso/pkg/cfg"
	"blion-worker-consenso/pkg/cfg/blockchain"
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpcMetadata "google.golang.org/grpc/metadata"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	minersAccepted     = map[string]string{}
	validatorsAccepted = map[string]string{}
)

type Worker struct {
	SrvBK         *bk.Server
	SrvAuth       *auth.Server
	SrvCfg        *cfg.Server
	ClientMine    mine_proto.MineBlockServicesBlocksClient
	ClientAuth    auth_proto.AuthServicesUsersClient
	ClientAccount accounting_proto.AccountingServicesAccountingClient
	ClientTrx     transactions_proto.TransactionsServicesClient
	ClientWallet  wallet_proto.WalletServicesWalletClient
	Block         *mine_proto.DataBlockMine
	Ctx           context.Context
	ConfigBk      *blockchain.Blockchain
}

func NewWorker(srvBk *bk.Server, srvAuth *auth.Server, srvCfg *cfg.Server) IWorker {
	return &Worker{SrvBK: srvBk, SrvAuth: srvAuth, SrvCfg: srvCfg}
}

func (w *Worker) Execute() {

	c := env.NewConfiguration()
	for {
		w.WorkerChanel()
		time.Sleep(time.Duration(c.App.WorkerInterval) * time.Second)
	}
}

func (w *Worker) WorkerChanel() {

	err := w.initGrpcServices()
	if err != nil {
		return
	}

	lottery, _, err := w.SrvBK.SrvLottery.GetLotteryActiveOrReadyMined()
	if err != nil {
		logger.Error.Printf("error trayendo una lotería activa: %s", err)
		return
	}

	if lottery == nil {
		lottery, err = w.createLottery()
		if err != nil {
			return
		}
	}

	color.Blue("Iniciando worker loteria")
	workChan := make(chan *lotteries.Lottery, 1)
	var wg sync.WaitGroup
	workChan <- lottery
	wg.Add(1)
	go func() {
		defer wg.Done()
		for true {
			workItem, more := <-workChan
			if more {
				w.doWork(workItem)
			} else {
				log.Printf("No more work to do")
				break
			}
		}
		log.Printf("Worker Terminated")
	}()

	close(workChan)
	wg.Wait()
}

func (w *Worker) doWork(workItem *lotteries.Lottery) {
	if workItem.ProcessStatus == 25 {
		err := w.processLotteryActive(workItem)
		if err != nil {
			logger.Error.Printf("No se pudo procesar la loteria activa, err: %s", err.Error())
			return
		}
		return
	}

	err := w.ProcessLotteryOfMined(workItem)
	if err != nil {
		logger.Error.Printf("No se pudo procesar la loteria lista para minar, err: %s", err.Error())
		return
	}

	return
}

func (w *Worker) processLotteryActive(lotteryActive *lotteries.Lottery) error {
	participantsIds := map[string]string{}
	lifeLottery := time.Now().Sub(lotteryActive.RegistrationStartDate).Seconds()
	if int(lifeLottery) > (w.ConfigBk.LotteryTtl * 1000) {

		participantsActive, _, err := w.SrvBK.SrvParticipants.GetParticipantsByLotteryID(lotteryActive.ID)
		if err != nil {
			logger.Error.Printf("error trayendo los participantes: %s", err)
			return err
		}

		if participantsActive == nil || len(participantsActive) < (w.ConfigBk.MaxValidators+w.ConfigBk.MaxMiners) {
			return nil
		}

		endDate := time.Now()
		lotteryUpdated, _, err := w.SrvBK.SrvLottery.UpdateLottery(lotteryActive.ID, lotteryActive.BlockId, lotteryActive.RegistrationStartDate, &endDate, &endDate, nil, nil, 26)
		if err != nil {
			logger.Error.Printf("error iniciando la loteria: %s", err)
			return err
		}

		rand.Shuffle(len(participantsActive), func(i, j int) {
			participantsActive[i], participantsActive[j] = participantsActive[j], participantsActive[i]
		})

		for _, p := range participantsActive {
			tickets := int64(math.Round(p.Amount / float64(w.ConfigBk.TicketsPrice)))
			for j := int64(0); j < tickets; j++ {
				participantsIds[p.ID+"/"+strconv.FormatInt(j, 10)] = p.WalletId
			}
		}

		GetMinersAndValidators(participantsIds, w.ConfigBk.MaxMiners, w.ConfigBk.MaxValidators)

		for key := range minersAccepted {
			participant := findOneParticipant(participantsActive, strings.Split("/", key)[0])
			_, _, err = w.SrvBK.SrvParticipants.UpdateParticipants(participant.ID, lotteryActive.ID, participant.WalletId, participant.Amount, true, 23, false)
			if err != nil {
				logger.Error.Printf("error actulizando el registro del participante: %s", err)
				return err
			}
		}

		for key := range validatorsAccepted {
			participant := findOneParticipant(participantsActive, strings.Split("/", key)[0])
			_, _, err = w.SrvBK.SrvParticipants.UpdateParticipants(participant.ID, lotteryActive.ID, participant.WalletId, participant.Amount, true, 24, false)
			if err != nil {
				logger.Error.Printf("error actulizando el registro del participante: %s", err)
				return err
			}
		}

		minersAccepted = make(map[string]string)
		validatorsAccepted = make(map[string]string)

		endDate = time.Now()
		_, _, err = w.SrvBK.SrvLottery.UpdateLottery(lotteryUpdated.ID, lotteryUpdated.BlockId, lotteryUpdated.RegistrationStartDate, lotteryUpdated.RegistrationEndDate, lotteryUpdated.LotteryStartDate, &endDate, nil, 27)
		if err != nil {
			logger.Error.Printf("error finalizando la loteria: %s", err)
			return err
		}
	}
	return nil
}

func (w *Worker) ProcessLotteryOfMined(lottery *lotteries.Lottery) error {
	resHash, _, err := w.SrvBK.SrvMinerResponse.GetMinerResponseRegister(lottery.ID)
	if err != nil {
		logger.Error.Printf("error trayendo el hash del minero: %s", err)
		return err
	}
	votesInFavor := 0
	votes, err := w.SrvBK.SrvValidatorsVote.GetAllValidatorVotesByLotteryID(lottery.ID)
	if err != nil {
		logger.Error.Printf("error trayendo los votos de los validadores: %s", err)
		return err
	}

	if len(votes) != w.ConfigBk.MaxValidators {
		logger.Info.Printf("Se requiere la totalidad de los votos designado en la loteria")
		return nil
	}

	for _, vote := range votes {
		if vote.Vote {
			votesInFavor++
		}
	}

	participantsLottery, _, err := w.SrvBK.SrvParticipants.GetParticipantsByLotteryID(lottery.ID)
	if err != nil {
		logger.Error.Printf("error trayendo los participantes de la loteria: %s", err)
		return err
	}

	if (votesInFavor*100)/len(votes) > 51 {
		return w.isApproved(w.Block.Id, resHash, lottery, participantsLottery)
	}

	return w.isNotApproved(resHash, lottery, participantsLottery)
}

func (w *Worker) isApproved(block int64, resHash *miner_response.MinerResponse, lotteryMined *lotteries.Lottery, participantsLottery []*participants.Participants) error {

	resMine, err := w.ClientMine.MineBlock(w.Ctx, &mine_proto.RequestMineBlock{
		Id:         block,
		Hash:       resHash.Hash,
		Nonce:      resHash.Nonce,
		Difficulty: int32(resHash.Difficulty),
		MinerId:    findOneParticipant(participantsLottery, resHash.ParticipantsId).WalletId,
	})
	if err != nil {
		logger.Error.Printf("error agregando el bloque temporal a la cadena de bloques: %s", err)
		return err
	}

	if resMine == nil {
		logger.Error.Printf("error agregando el bloque temporal a la cadena de bloques")
		return fmt.Errorf("error agregando el bloque temporal a la cadena de bloques")
	}

	if resMine.Error {
		logger.Error.Printf(resMine.Msg)
		return fmt.Errorf(resMine.Msg)
	}

	_, _, err = w.SrvBK.SrvMinerResponse.UpdateMinerResponse(resHash.ID, lotteryMined.ID, resHash.ParticipantsId, resHash.Hash, 31, resHash.Nonce, resHash.Difficulty)
	if err != nil {
		logger.Error.Printf("error actualizando el hash del minero como aceptado: %s", err)
		return err
	}
	endProcess := time.Now()
	_, _, err = w.SrvBK.SrvLottery.UpdateLottery(lotteryMined.ID, lotteryMined.BlockId, lotteryMined.RegistrationStartDate, lotteryMined.RegistrationEndDate, lotteryMined.LotteryStartDate, lotteryMined.LotteryEndDate, &endProcess, 28)
	if err != nil {
		logger.Error.Printf("error actualizando la loteria: %s", err)
		return err
	}

	feeBlock, _, err := w.SrvBK.SrvBlockFee.GetBlockFeeByBlockID(block)
	if err != nil {
		logger.Error.Printf("error trayendo el fee del bloque: %v", err)
		return err
	}

	err = w.processTransactions(lotteryMined.BlockId, w.Ctx)
	if err != nil {
		logger.Error.Printf("error procesando las transacciones: %v", err)
		return err
	}

	feeMiner := (feeBlock.Fee * float64(w.ConfigBk.FeeMiner)) / 100
	feeValidators := (feeBlock.Fee * float64(w.ConfigBk.FeeValidator)) / 100
	feeNodes := (feeBlock.Fee * float64(w.ConfigBk.FeeNode)) / 100

	for _, participant := range participantsLottery {
		resUnfreeze, err := w.ClientWallet.UnFreezeMoney(w.Ctx, &wallet_proto.RqUnFreezeMoney{WalletId: participant.WalletId, LotteryId: lotteryMined.ID})
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

		_, _, err = w.SrvBK.SrvParticipants.UpdateParticipants(participant.ID, participant.LotteryId, participant.WalletId, participant.Amount, participant.Accepted, participant.TypeCharge, true)
		if err != nil {
			logger.Error.Printf("error actualizando el participante", err)
			continue
		}

		if participant.Accepted && (participant.ID == resHash.ParticipantsId || participant.TypeCharge == 24) {
			resAccount, err := w.ClientAccount.GetAccountingByWalletById(w.Ctx, &accounting_proto.RequestGetAccountingByWalletId{Id: participant.WalletId})
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
				amount = feeValidators / float64(w.ConfigBk.MaxValidators)
			}

			resReward, err := w.ClientAccount.SetAmountToAccounting(w.Ctx, &accounting_proto.RequestSetAmountToAccounting{
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

			_, _, err = w.SrvBK.SrvReward.CreateReward(uuid.New().String(), lotteryMined.ID, participant.WalletId, amount)
			if err != nil {
				logger.Error.Printf("error creando el registro de recompenzas, error: %v", err)
				continue
			}
		}
	}

	w.ProcessNode(feeNodes)

	resAccountMain, err := w.ClientAccount.GetAccountingByWalletById(w.Ctx, &accounting_proto.RequestGetAccountingByWalletId{Id: w.ConfigBk.WalletMain})
	if err != nil {
		logger.Error.Printf("error obteniendo la cuenta principal, error: %v", err)
		return err
	}
	if resAccountMain == nil {
		logger.Error.Printf("error obteniendo la cuenta principal")
		return fmt.Errorf("error obteniendo la cuenta principal")
	}
	if resAccountMain.Error {
		logger.Error.Printf(resAccountMain.Msg)
		return fmt.Errorf(resAccountMain.Msg)
	}

	fee := feeBlock.Fee - (feeMiner + feeValidators)
	resMain, err := w.ClientAccount.SetAmountToAccounting(w.Ctx, &accounting_proto.RequestSetAmountToAccounting{
		WalletId: w.ConfigBk.WalletMain,
		Amount:   resAccountMain.Data.Amount + fee,
		IdUser:   resAccountMain.Data.IdUser,
	})
	if err != nil {
		logger.Error.Printf("error actualizando el dinero de la cuenta principal, error: %v", err)
		return err
	}
	if resMain == nil {
		logger.Error.Printf("error actualizando el dinero de la cuenta principal")
		return fmt.Errorf("error actualizando el dinero de la cuenta principal")
	}
	if resMain.Error {
		logger.Error.Printf(resMain.Msg)
		return fmt.Errorf(resMain.Msg)
	}

	return nil
}

func (w *Worker) isNotApproved(resHash *miner_response.MinerResponse, lotteryMined *lotteries.Lottery, participantsLottery []*participants.Participants) error {
	_, _, err := w.SrvBK.SrvMinerResponse.UpdateMinerResponse(resHash.ID, lotteryMined.ID, resHash.ParticipantsId, resHash.Hash, 30, resHash.Nonce, resHash.Difficulty)
	if err != nil {
		logger.Error.Printf("error actualizando el hash del minero como aceptado: %s", err)
		return err
	}

	endProcess := time.Now()
	_, _, err = w.SrvBK.SrvLottery.UpdateLottery(lotteryMined.ID, lotteryMined.BlockId, lotteryMined.RegistrationStartDate, lotteryMined.RegistrationEndDate, lotteryMined.LotteryStartDate, lotteryMined.LotteryEndDate, &endProcess, 32)
	if err != nil {
		logger.Error.Printf("error actualizando la loteria: %s", err)
		return err
	}

	miner := findOneParticipant(participantsLottery, resHash.ParticipantsId)
	walletPenalties, err := w.SrvBK.SrvPenaltyParticipant.GetAllPenaltyParticipantsByWalletID(miner.WalletId)
	if err != nil {
		logger.Error.Printf("error trayendo las multas del minero: %s", err)
		return err
	}

	resFreezeMoney, err := w.ClientWallet.GetFrozenMoney(w.Ctx, &wallet_proto.RqGetFrozenMoney{WalletId: miner.WalletId})
	if err != nil {
		logger.Error.Printf("error trayendo el dinero congelado: %s", err)
		return err
	}

	if resFreezeMoney == nil {
		logger.Error.Printf("error trayendo el dinero congelado")
		return fmt.Errorf("error trayendo el dinero congelado")
	}

	if resFreezeMoney.Error {
		logger.Error.Printf(resFreezeMoney.Msg)
		return fmt.Errorf(resFreezeMoney.Msg)
	}
	freezeMoney := resFreezeMoney.Data

	amountPenalty := freezeMoney * 10 / 100
	penaltyPercentage := 10.0

	if walletPenalties != nil && len(walletPenalties) > 0 {
		if len(walletPenalties) == 1 {
			penaltyPercentage = 25
			amountPenalty = (freezeMoney * 25) / 100
		} else if len(walletPenalties) == 2 {
			penaltyPercentage = 50
			amountPenalty = (freezeMoney * 50) / 100
		} else if len(walletPenalties) >= 3 {
			penaltyPercentage = 100
			amountPenalty = freezeMoney
		}
	}

	_, _, err = w.SrvBK.SrvPenaltyParticipant.CreatePenaltyParticipants(uuid.New().String(), lotteryMined.ID, resHash.ParticipantsId, amountPenalty, penaltyPercentage)
	if err != nil {
		logger.Error.Printf("error multando al minero por fraude: %v", err)
		return err
	}

	resUnfreeze, err := w.ClientWallet.UnFreezeMoney(w.Ctx, &wallet_proto.RqUnFreezeMoney{WalletId: miner.WalletId, LotteryId: lotteryMined.ID, Penalty: amountPenalty})
	if err != nil {
		logger.Error.Printf("error descongelando el dinero: %s", err)
		return err
	}

	if resUnfreeze == nil {
		logger.Error.Printf("error descongelando el dinero")
		return fmt.Errorf("error descongelando el dinero")
	}

	if resUnfreeze.Error {
		logger.Error.Printf(resUnfreeze.Msg)
		return fmt.Errorf(resUnfreeze.Msg)
	}

	_, _, err = w.SrvBK.SrvParticipants.UpdateParticipants(miner.ID, miner.LotteryId, miner.WalletId, miner.Amount, miner.Accepted, miner.TypeCharge, true)
	if err != nil {
		logger.Error.Printf("error actualizando el participante", err)
		return err
	}

	for _, participant := range participantsLottery {
		if participant.ID != resHash.ParticipantsId {
			resUnfreeze, err = w.ClientWallet.UnFreezeMoney(w.Ctx, &wallet_proto.RqUnFreezeMoney{WalletId: participant.WalletId, LotteryId: lotteryMined.ID, Penalty: 0})
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

			_, _, err = w.SrvBK.SrvParticipants.UpdateParticipants(participant.ID, participant.LotteryId, participant.WalletId, participant.Amount, participant.Accepted, participant.TypeCharge, true)
			if err != nil {
				logger.Error.Printf("error actualizando el participante", err)
				continue
			}
		}
	}

	return nil
}

func (w *Worker) getBlockToMine(ctx context.Context) (*mine_proto.DataBlockMine, error) {
	resBkMine, err := w.ClientMine.GetBlockToMine(ctx, &mine_proto.GetBlockToMineRequest{})
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

func (w *Worker) processTransactions(blockId int64, ctx context.Context) error {

	resWsAllTrx, err := w.ClientTrx.GetTransactionsByBlockId(ctx, &transactions_proto.RqGetTransactionByBlock{BlockId: blockId})
	if err != nil {
		logger.Error.Printf("no se pudo obtener todas las transacciones para minar, error: %V", err)
		return err
	}

	if resWsAllTrx == nil {
		logger.Error.Printf("no se pudo obtener todas las transacciones para minar")
		return fmt.Errorf("no se pudo obtener todas las transacciones para minar")
	}

	if resWsAllTrx.Error {
		logger.Error.Printf(resWsAllTrx.Msg)
		return fmt.Errorf(resWsAllTrx.Msg)
	}

	transactions := resWsAllTrx.Data

	for _, transaction := range transactions {
		resAccountTo, err := w.ClientAccount.GetAccountingByWalletById(ctx, &accounting_proto.RequestGetAccountingByWalletId{Id: transaction.To})
		if err != nil {
			logger.Error.Printf("couldn't get account to wallet by id_wallet: %v", err)
			return err
		}

		if resAccountTo == nil {
			logger.Error.Printf("couldn't get account to wallet by id_wallet: %v", err)
			return fmt.Errorf("couldn't get account to wallet by id_wallet")
		}

		if resAccountTo.Error {
			logger.Error.Printf(resAccountTo.Msg)
			return fmt.Errorf(resAccountTo.Msg)
		}

		accountTo := resAccountTo.Data

		accountToAmounted, err := w.ClientAccount.SetAmountToAccounting(ctx, &accounting_proto.RequestSetAmountToAccounting{
			WalletId: transaction.To,
			Amount:   accountTo.Amount + accountTo.Amount,
			IdUser:   accountTo.IdUser,
		})
		if err != nil {
			logger.Error.Printf("couldn't update amount from user: %v", err)
			return err
		}

		if accountToAmounted == nil {
			logger.Error.Printf("couldn't update amount from user")
			return fmt.Errorf("couldn't update amount from use")
		}

		if accountToAmounted.Error {
			logger.Error.Printf(accountToAmounted.Msg)
			return fmt.Errorf(accountToAmounted.Msg)
		}
	}

	return nil
}

func (w *Worker) login() (string, error) {
	e := env.NewConfiguration()
	resLogin, err := w.ClientAuth.Login(context.Background(), &auth_proto.LoginRequest{
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

func (w *Worker) createLottery() (*lotteries.Lottery, error) {

	lottery, _, err := w.SrvBK.SrvLottery.CreateLottery(uuid.New().String(), w.Block.Id, time.Now(), nil, nil, nil, nil, 25)
	if err != nil {
		logger.Error.Printf("No se pudo obtener el bloque a minar, error: %s", err)
		return nil, err
	}

	return lottery, nil
}

func (w *Worker) initGrpcServices() error {
	c := env.NewConfiguration()
	connBk, err := grpc.Dial(c.BlockService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error.Printf("error conectando con el servicio auth de blockchain: %s", err)
		return err
	}
	defer connBk.Close()

	connAuth, err := grpc.Dial(c.AuthService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error.Printf("error conectando con el servicio auth de blockchain: %s", err)
		return err
	}
	defer connAuth.Close()

	connTrx, err := grpc.Dial(c.TransactionsService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error.Printf("error conectando con el servicio auth de blockchain: %s", err)
		return err
	}
	defer connAuth.Close()

	w.ClientAuth = auth_proto.NewAuthServicesUsersClient(connAuth)
	w.ClientMine = mine_proto.NewMineBlockServicesBlocksClient(connBk)
	w.ClientAccount = accounting_proto.NewAccountingServicesAccountingClient(connAuth)
	w.ClientWallet = wallet_proto.NewWalletServicesWalletClient(connAuth)
	w.ClientTrx = transactions_proto.NewTransactionsServicesClient(connTrx)

	token, err := w.login()
	if err != nil {
		logger.Error.Printf("No se pudo obtener el token de autorización, error: %v", err)
		return err
	}
	w.Ctx = grpcMetadata.AppendToOutgoingContext(context.Background(), "authorization", token)

	w.Block, err = w.getBlockToMine(w.Ctx)
	if err != nil {
		logger.Error.Printf("No se pudo obtener el bloque a minar, error: %v", err)
		return nil
	}

	configBk, err := w.SrvCfg.SrvBlockchain.GetLasted()
	if err != nil {
		logger.Error.Printf("No se pudo cargar la configuración de la bk, err: ", err)
		return err
	}

	w.ConfigBk = configBk

	return nil
}

func (w *Worker) ProcessNode(feeNodes float64) {
	nodes, err := w.SrvAuth.SrvNodeWallet.GetAllNodeWallet()
	if err != nil {
		logger.Error.Printf("No se pudo obtener los nodos: %v", err)
		return
	}

	if len(nodes) <= 0 {
		return
	}

	for _, node := range nodes {
		resAccountNode, err := w.ClientAccount.GetAccountingByWalletById(w.Ctx, &accounting_proto.RequestGetAccountingByWalletId{Id: node.WalletId})
		if err != nil {
			logger.Error.Printf("error obteniendo la cuenta del nodo, error: %v", err)
			continue
		}
		if resAccountNode == nil {
			logger.Error.Printf("error obteniendo la cuenta del nodo")
			continue
		}
		if resAccountNode.Error {
			logger.Error.Printf(resAccountNode.Msg)
			continue
		}

		resNode, err := w.ClientAccount.SetAmountToAccounting(w.Ctx, &accounting_proto.RequestSetAmountToAccounting{
			WalletId: node.WalletId,
			Amount:   resAccountNode.Data.Amount + (feeNodes / float64(len(nodes))),
			IdUser:   resAccountNode.Data.IdUser,
		})
		if err != nil {
			logger.Error.Printf("error actualizando el dinero del nodo, error: %v", err)
			continue
		}
		if resNode == nil {
			logger.Error.Printf("error actualizando el dinero del nodo")
			continue
		}
		if resNode.Error {
			logger.Error.Printf(resNode.Msg)
			continue
		}
	}
}

func GetMinersAndValidators(participantsIds map[string]string, maxMiners, maxValidators int) {

	numberAccepted := rand.Intn(60-1) + 1
	counter := 0
	participantID := ""
	for key, participant := range participantsIds {
		if counter == numberAccepted {
			if len(minersAccepted) == 0 || len(minersAccepted) == len(validatorsAccepted) {
				minersAccepted[key] = participant
			} else if len(validatorsAccepted) < len(minersAccepted) {
				validatorsAccepted[key] = participant
			}
			participantID = strings.Split("/", key)[0]
			delete(participantsIds, key)
			break
		}
		counter++
	}

	for key := range participantsIds {
		if participantID == strings.Split("/", key)[0] {
			delete(participantsIds, key)
		}
	}

	if len(minersAccepted) == maxMiners && len(validatorsAccepted) == maxValidators {
		return
	}

	GetMinersAndValidators(participantsIds, maxMiners, maxValidators)
}

func findOneParticipant(participants []*participants.Participants, value string) *participants.Participants {
	for _, participant := range participants {
		if participant.ID == value {
			return participant
		}
	}
	return nil
}
