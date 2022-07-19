package bk

import (
	"blion-worker-consenso/internal/models"
	"blion-worker-consenso/pkg/bk/lotteries"
	"blion-worker-consenso/pkg/bk/miner_response"
	"blion-worker-consenso/pkg/bk/participants"
	"blion-worker-consenso/pkg/bk/rewards"
	"blion-worker-consenso/pkg/bk/validator_votes"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	SrvLottery        lotteries.PortsServerLottery
	SrvParticipants   participants.PortsServerParticipants
	SrvReward         rewards.PortsServerRewards
	SrvValidatorsVote validator_votes.PortsServerValidatorVotes
	SrvMinerResponse  miner_response.PortsServerMinerResponse
}

func NewServerBk(db *sqlx.DB, user *models.User, txID string) *Server {

	repoLottery := lotteries.FactoryStorage(db, user, txID)
	srvLottery := lotteries.NewLotteryService(repoLottery, user, txID)

	repoParticipants := participants.FactoryStorage(db, user, txID)
	srvParticipants := participants.NewParticipantsService(repoParticipants, user, txID)

	repoReward := rewards.FactoryStorage(db, user, txID)
	srvReward := rewards.NewRewardsService(repoReward, user, txID)

	repoValidatorsVote := validator_votes.FactoryStorage(db, user, txID)
	srvValidatorsVote := validator_votes.NewValidatorVotesService(repoValidatorsVote, user, txID)

	repoMinerResponse := miner_response.FactoryStorage(db, user, txID)
	srvMinerResponse := miner_response.NewMinerResponseService(repoMinerResponse, user, txID)

	return &Server{
		SrvLottery:        srvLottery,
		SrvParticipants:   srvParticipants,
		SrvReward:         srvReward,
		SrvValidatorsVote: srvValidatorsVote,
		SrvMinerResponse:  srvMinerResponse,
	}
}
