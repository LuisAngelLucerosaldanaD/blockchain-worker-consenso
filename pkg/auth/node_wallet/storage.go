package node_wallet

import (
	"github.com/jmoiron/sqlx"

	"blion-worker-consenso/internal/logger"
	"blion-worker-consenso/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesNodeWalletRepository interface {
	create(m *NodeWallet) error
	update(m *NodeWallet) error
	delete(id string) error
	getByID(id string) (*NodeWallet, error)
	getAll() ([]*NodeWallet, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesNodeWalletRepository {
	var s ServicesNodeWalletRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newNodeWalletPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
