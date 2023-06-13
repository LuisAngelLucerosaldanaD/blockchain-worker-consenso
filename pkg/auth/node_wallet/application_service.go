package node_wallet

import (
	"fmt"

	"blion-worker-consenso/internal/logger"
	"blion-worker-consenso/internal/models"
	"github.com/asaskevich/govalidator"
	"time"
)

type PortsServerNodeWallet interface {
	CreateNodeWallet(id string, walletId string, name string, ip string, deletedAt *time.time, penaltyAt *time.time) (*NodeWallet, int, error)
	UpdateNodeWallet(id string, walletId string, name string, ip string, deletedAt *time.time, penaltyAt *time.time) (*NodeWallet, int, error)
	DeleteNodeWallet(id string) (int, error)
	GetNodeWalletByID(id string) (*NodeWallet, int, error)
	GetAllNodeWallet() ([]*NodeWallet, error)
}

type service struct {
	repository ServicesNodeWalletRepository
	user       *models.User
	txID       string
}

func NewNodeWalletService(repository ServicesNodeWalletRepository, user *models.User, TxID string) PortsServerNodeWallet {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateNodeWallet(id string, walletId string, name string, ip string, deletedAt *time.time, penaltyAt *time.time) (*NodeWallet, int, error) {
	m := NewNodeWallet(id, walletId, name, ip, deletedAt, penaltyAt)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create NodeWallet :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateNodeWallet(id string, walletId string, name string, ip string, deletedAt *time.time, penaltyAt *time.time) (*NodeWallet, int, error) {
	m := NewNodeWallet(id, walletId, name, ip, deletedAt, penaltyAt)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update NodeWallet :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteNodeWallet(id string) (int, error) {
	if !govalidator.IsUUID(id) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return 15, fmt.Errorf("id isn't uuid")
	}

	if err := s.repository.delete(id); err != nil {
		if err.Error() == "ecatch:108" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s *service) GetNodeWalletByID(id string) (*NodeWallet, int, error) {
	if !govalidator.IsUUID(id) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return nil, 15, fmt.Errorf("id isn't uuid")
	}
	m, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllNodeWallet() ([]*NodeWallet, error) {
	return s.repository.getAll()
}
