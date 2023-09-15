package validator_vote

import (
	"github.com/jmoiron/sqlx"

	"blion-worker-consenso/internal/logger"
	"blion-worker-consenso/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesValidatorVoteRepository interface {
	create(m *ValidatorVote) error
	update(m *ValidatorVote) error
	delete(id string) error
	getByID(id string) (*ValidatorVote, error)
	getAll() ([]*ValidatorVote, error)
	getByLotteryID(lotteryID string) ([]*ValidatorVote, error)
	getVotesInFavor(id string) (int64, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesValidatorVoteRepository {
	var s ServicesValidatorVoteRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newValidatorVotePsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
