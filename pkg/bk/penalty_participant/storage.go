package penalty_participant

import (
	"github.com/jmoiron/sqlx"

	"blion-worker-consenso/internal/logger"
	"blion-worker-consenso/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesPenaltyParticipantRepository interface {
	create(m *PenaltyParticipant) error
	update(m *PenaltyParticipant) error
	delete(id string) error
	getByID(id string) (*PenaltyParticipant, error)
	getAll() ([]*PenaltyParticipant, error)
	getByWalletID(walletId string) ([]*PenaltyParticipant, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesPenaltyParticipantRepository {
	var s ServicesPenaltyParticipantRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newPenaltyParticipantPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
