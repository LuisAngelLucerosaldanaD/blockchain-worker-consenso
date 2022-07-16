package bk

import (
	"blion-worker-consenso/internal/models"
	"blion-worker-consenso/pkg/bk/lottery_table"
	"blion-worker-consenso/pkg/bk/participants_table"
	"blion-worker-consenso/pkg/bk/reward_table"
	"blion-worker-consenso/pkg/bk/validation_table"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	SrvLottery      lottery_table.PortsServerLotteryTable
	SrvParticipants participants_table.PortsServerParticipantsTable
	SrvReward       reward_table.PortsServerRewardTable
	SrvValidation   validation_table.PortsServerValidationTable
}

func NewServerBk(db *sqlx.DB, user *models.User, txID string) *Server {

	repoLottery := lottery_table.FactoryStorage(db, user, txID)
	srvLottery := lottery_table.NewLotteryTableService(repoLottery, user, txID)

	repoParticipants := participants_table.FactoryStorage(db, user, txID)
	srvParticipants := participants_table.NewParticipantsTableService(repoParticipants, user, txID)

	repoReward := reward_table.FactoryStorage(db, user, txID)
	srvReward := reward_table.NewRewardTableService(repoReward, user, txID)

	repoValidation := validation_table.FactoryStorage(db, user, txID)
	srvValidation := validation_table.NewValidationTableService(repoValidation, user, txID)

	return &Server{
		SrvLottery:      srvLottery,
		SrvParticipants: srvParticipants,
		SrvReward:       srvReward,
		SrvValidation:   srvValidation,
	}
}
