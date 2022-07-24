package penalty_participants

import (
	"github.com/jmoiron/sqlx"

	"blion-worker-consenso/internal/logger"
	"blion-worker-consenso/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesPenaltyParticipantsRepository interface {
	create(m *PenaltyParticipants) error
	update(m *PenaltyParticipants) error
	delete(id string) error
	getByID(id string) (*PenaltyParticipants, error)
	getAll() ([]*PenaltyParticipants, error)
	getByWalletID(walletId string) ([]*PenaltyParticipants, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesPenaltyParticipantsRepository {
	var s ServicesPenaltyParticipantsRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newPenaltyParticipantsPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
