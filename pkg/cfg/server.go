package cfg

import (
	"blion-worker-consenso/internal/models"
	"blion-worker-consenso/pkg/cfg/blockchain"
	"blion-worker-consenso/pkg/cfg/messages"

	"github.com/jmoiron/sqlx"
)

type Server struct {
	SrvMessage    messages.PortsServerMessages
	SrvBlockchain blockchain.PortsServerBlockchain
}

func NewServerCfg(db *sqlx.DB, user *models.User, txID string) *Server {

	repoMessage := messages.FactoryStorage(db, user, txID)
	srvMessage := messages.NewMessagesService(repoMessage, user, txID)

	repoBlockchain := blockchain.FactoryStorage(db, user, txID)
	srvBlockchain := blockchain.NewBlockchainService(repoBlockchain, user, txID)

	return &Server{
		SrvMessage:    srvMessage,
		SrvBlockchain: srvBlockchain,
	}
}
