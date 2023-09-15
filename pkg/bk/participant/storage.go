package participant

import (
	"github.com/jmoiron/sqlx"

	"blion-worker-consenso/internal/logger"
	"blion-worker-consenso/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesParticipantRepository interface {
	create(m *Participant) error
	update(m *Participant) error
	delete(id string) error
	getByID(id string) (*Participant, error)
	getAll() ([]*Participant, error)
	getByWalletID(walletId string) (*Participant, error)
	getByLotteryID(lotteryId string) ([]*Participant, error)
	getByWalletIDAndLotteryID(walletId string, lotteryId string) (*Participant, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesParticipantRepository {
	var s ServicesParticipantRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newParticipantPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
