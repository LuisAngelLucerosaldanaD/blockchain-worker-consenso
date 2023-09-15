package penalty_participant

import (
	"fmt"

	"blion-worker-consenso/internal/logger"
	"blion-worker-consenso/internal/models"
	"github.com/asaskevich/govalidator"
)

type PortsServerPenaltyParticipant interface {
	CreatePenaltyParticipant(id string, participantsId string, amount float64, penaltyPercentage float64) (*PenaltyParticipant, int, error)
	UpdatePenaltyParticipant(id string, participantsId string, amount float64, penaltyPercentage float64) (*PenaltyParticipant, int, error)
	DeletePenaltyParticipant(id string) (int, error)
	GetPenaltyParticipantByID(id string) (*PenaltyParticipant, int, error)
	GetAllPenaltyParticipants() ([]*PenaltyParticipant, error)
	GetAllPenaltyParticipantByWalletID(walletID string) ([]*PenaltyParticipant, error)
}

type service struct {
	repository ServicesPenaltyParticipantRepository
	user       *models.User
	txID       string
}

func NewPenaltyParticipantService(repository ServicesPenaltyParticipantRepository, user *models.User, TxID string) PortsServerPenaltyParticipant {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreatePenaltyParticipant(id string, participantsId string, amount float64, penaltyPercentage float64) (*PenaltyParticipant, int, error) {
	m := NewPenaltyParticipants(id, participantsId, amount, penaltyPercentage)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create PenaltyParticipant :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdatePenaltyParticipant(id string, participantsId string, amount float64, penaltyPercentage float64) (*PenaltyParticipant, int, error) {
	m := NewPenaltyParticipants(id, participantsId, amount, penaltyPercentage)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update PenaltyParticipant :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeletePenaltyParticipant(id string) (int, error) {
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

func (s *service) GetPenaltyParticipantByID(id string) (*PenaltyParticipant, int, error) {
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

func (s *service) GetAllPenaltyParticipants() ([]*PenaltyParticipant, error) {
	return s.repository.getAll()
}

func (s *service) GetAllPenaltyParticipantByWalletID(walletID string) ([]*PenaltyParticipant, error) {
	return s.repository.getByWalletID(walletID)
}
