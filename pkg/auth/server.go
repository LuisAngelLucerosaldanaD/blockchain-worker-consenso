package auth

import (
	"blion-worker-consenso/internal/models"
	"blion-worker-consenso/pkg/auth/node_wallet"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	SrvNodeWallet node_wallet.PortsServerNodeWallet
}

func NewServerAuth(db *sqlx.DB, user *models.User, txID string) *Server {

	repoNode := node_wallet.FactoryStorage(db, user, txID)
	srvNodeWallet := node_wallet.NewNodeWalletService(repoNode, user, txID)

	return &Server{
		SrvNodeWallet: srvNodeWallet,
	}
}
