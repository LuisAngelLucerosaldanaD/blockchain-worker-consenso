package worker

import (
	"blion-worker-consenso/internal/env"
	"blion-worker-consenso/internal/grpc/auth_proto"
	"blion-worker-consenso/internal/grpc/mine_proto"
	"blion-worker-consenso/internal/grpc/wallet_proto"
	"blion-worker-consenso/internal/logger"
	"blion-worker-consenso/pkg/bk"
	"context"
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
	clientWallet := wallet_proto.NewWalletServicesWalletClient(connAuth)

	resLogin, err := clientAuth.Login(context.Background(), &auth_proto.LoginRequest{
		Email:    nil,
		Nickname: &e.App.UserLogin,
		Password: e.App.UserPassword,
	})
	if err != nil {
		logger.Error.Printf("No se pudo obtener el token de autenticacion: %s", err)
		return
	}

	if resLogin == nil {
		logger.Error.Printf("No se pudo obtener el token de autenticacion")
		return
	}

	if resLogin.Error {
		logger.Error.Printf(resLogin.Msg)
		return
	}

	ctx := grpcMetadata.AppendToOutgoingContext(context.Background(), "authorization", resLogin.Data.AccessToken)

	resBkMine, err := clientMine.GetBlockToMine(ctx, &mine_proto.GetBlockToMineRequest{})
	if err != nil {
		logger.Error.Printf("No se pudo obtener el bloque a minar, error: %s", err)
	}

	if resBkMine == nil {
		logger.Error.Printf("No se pudo obtener el bloque a minar")
		return
	}

	if resBkMine.Error {
		logger.Error.Printf(resBkMine.Msg)
		return
	}

	if resBkMine.Data == nil {
		logger.Info.Println(resBkMine.Msg)
		return
	}

	block := resBkMine.Data

	lotteryActive, _, err := w.Srv.SrvLottery.GetLotteryActive()
	if err != nil {
		logger.Error.Printf("error trayendo una loteria activa: %s", err)
		return
	}

	if lotteryActive != nil {
		lifeLottery := time.Now().Sub(lotteryActive.RegistrationStartDate).Seconds()
		if int(lifeLottery) > (e.App.SubscriptionTime * 1000) {
			endDate := time.Now()
			lotteryUpdated, _, err := w.Srv.SrvLottery.UpdateLottery(lotteryActive.ID, lotteryActive.BlockId, lotteryActive.RegistrationStartDate, &endDate, &endDate, nil, nil, 26)
			if err != nil {
				logger.Error.Printf("error iniciando la loteria: %s", err)
				return
			}

			participants, _, err := w.Srv.SrvParticipants.GetParticipantsByLotteryID(lotteryActive.ID)
			if err != nil {
				logger.Error.Printf("error trayendo los participantes: %s", err)
				return
			}

			participantsLottery := make([]string, len(participants))

			for i, participant := range participants {
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

			for _, participant := range participants {
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
		votesAgainst := 0
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
			} else {
				votesAgainst++
			}
		}

		isApproved := false

		if (votesInFavor*100)/len(votes) > 51 {
			isApproved = true
		}

		participants, _, err := w.Srv.SrvParticipants.GetParticipantsByLotteryID(lotteryMined.ID)
		if err != nil {
			logger.Error.Printf("error trayendo los participantes de la loteria: %s", err)
			return
		}

		if isApproved {
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

			for _, participant := range participants {
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
			}

			return
		}

		_, _, err = w.Srv.SrvMinerResponse.UpdateMinerResponse(resHash.ID, lotteryMined.ID, resHash.ParticipantsId, resHash.Hash, 30, resHash.Nonce, resHash.Difficulty)
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

		for _, participant := range participants {
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
		}

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
