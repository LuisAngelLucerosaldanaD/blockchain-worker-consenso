package reward

import (
	"github.com/jmoiron/sqlx"

	"blion-worker-consenso/internal/logger"
	"blion-worker-consenso/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesRewardRepository interface {
	create(m *Reward) error
	update(m *Reward) error
	delete(id string) error
	getByID(id string) (*Reward, error)
	getAll() ([]*Reward, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesRewardRepository {
	var s ServicesRewardRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newRewardsPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
