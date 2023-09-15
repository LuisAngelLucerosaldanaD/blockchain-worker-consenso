package bk

import (
	"blion-worker-consenso/internal/models"
	"blion-worker-consenso/pkg/bk/block_fee"
	"blion-worker-consenso/pkg/bk/lottery"
	"blion-worker-consenso/pkg/bk/miner_response"
	"blion-worker-consenso/pkg/bk/participant"
	"blion-worker-consenso/pkg/bk/penalty_participant"
	"blion-worker-consenso/pkg/bk/reward"
	"blion-worker-consenso/pkg/bk/validator_vote"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	SrvLottery            lottery.PortsServerLottery
	SrvParticipants       participant.PortsServerParticipant
	SrvReward             reward.PortsServerReward
	SrvValidatorsVote     validator_vote.PortsServerValidatorVote
	SrvMinerResponse      miner_response.PortsServerMinerResponse
	SrvPenaltyParticipant penalty_participant.PortsServerPenaltyParticipant
	SrvBlockFee           block_fee.PortsServerBlockFee
}

func NewServerBk(db *sqlx.DB, user *models.User, txID string) *Server {

	repoLottery := lottery.FactoryStorage(db, user, txID)
	srvLottery := lottery.NewLotteryService(repoLottery, user, txID)

	repoParticipants := participant.FactoryStorage(db, user, txID)
	srvParticipants := participant.NewParticipantService(repoParticipants, user, txID)

	repoReward := reward.FactoryStorage(db, user, txID)
	srvReward := reward.NewRewardService(repoReward, user, txID)

	repoValidatorsVote := validator_vote.FactoryStorage(db, user, txID)
	srvValidatorsVote := validator_vote.NewValidatorVoteService(repoValidatorsVote, user, txID)

	repoMinerResponse := miner_response.FactoryStorage(db, user, txID)
	srvMinerResponse := miner_response.NewMinerResponseService(repoMinerResponse, user, txID)

	repoPenaltyParticipant := penalty_participant.FactoryStorage(db, user, txID)
	srvPenaltyParticipant := penalty_participant.NewPenaltyParticipantService(repoPenaltyParticipant, user, txID)

	repoBlockFee := block_fee.FactoryStorage(db, user, txID)
	srvBlockFee := block_fee.NewBlockFeeService(repoBlockFee, user, txID)

	return &Server{
		SrvLottery:            srvLottery,
		SrvParticipants:       srvParticipants,
		SrvReward:             srvReward,
		SrvValidatorsVote:     srvValidatorsVote,
		SrvMinerResponse:      srvMinerResponse,
		SrvPenaltyParticipant: srvPenaltyParticipant,
		SrvBlockFee:           srvBlockFee,
	}
}
