package participant

import (
	"fmt"

	"blion-worker-consenso/internal/logger"
	"blion-worker-consenso/internal/models"
	"github.com/asaskevich/govalidator"
)

type PortsServerParticipant interface {
	CreateParticipant(id string, lotteryId string, walletId string, amount float64, accepted bool, typeCharge int, returned bool) (*Participant, int, error)
	UpdateParticipant(id string, lotteryId string, walletId string, amount float64, accepted bool, typeCharge int, returned bool) (*Participant, int, error)
	DeleteParticipant(id string) (int, error)
	GetParticipantByID(id string) (*Participant, int, error)
	GetAllParticipants() ([]*Participant, error)
	GetParticipantByWalletID(walletId string) (*Participant, int, error)
	GetParticipantsByLotteryID(lotteryID string) ([]*Participant, int, error)
	GetParticipantByWalletAndLottery(walletId string, lotteryID string) (*Participant, int, error)
}

type service struct {
	repository ServicesParticipantRepository
	user       *models.User
	txID       string
}

func NewParticipantService(repository ServicesParticipantRepository, user *models.User, TxID string) PortsServerParticipant {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateParticipant(id string, lotteryId string, walletId string, amount float64, accepted bool, typeCharge int, returned bool) (*Participant, int, error) {
	m := NewParticipant(id, lotteryId, walletId, amount, accepted, typeCharge, returned)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Participant :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateParticipant(id string, lotteryId string, walletId string, amount float64, accepted bool, typeCharge int, returned bool) (*Participant, int, error) {
	m := NewParticipant(id, lotteryId, walletId, amount, accepted, typeCharge, returned)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Participant :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteParticipant(id string) (int, error) {
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

func (s *service) GetParticipantByID(id string) (*Participant, int, error) {
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

func (s *service) GetParticipantByWalletID(walletId string) (*Participant, int, error) {
	if !govalidator.IsUUID(walletId) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("wallet id isn't uuid"))
		return nil, 15, fmt.Errorf("wallet id isn't uuid")
	}
	m, err := s.repository.getByWalletID(walletId)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByWalletID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetParticipantByWalletAndLottery(walletId string, lotteryID string) (*Participant, int, error) {
	m, err := s.repository.getByWalletIDAndLotteryID(walletId, lotteryID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByWalletID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetParticipantsByLotteryID(lotteryID string) ([]*Participant, int, error) {
	if !govalidator.IsUUID(lotteryID) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("lottery id isn't uuid"))
		return nil, 15, fmt.Errorf("lottery id isn't uuid")
	}
	m, err := s.repository.getByLotteryID(lotteryID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByWalletID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllParticipants() ([]*Participant, error) {
	return s.repository.getAll()
}
