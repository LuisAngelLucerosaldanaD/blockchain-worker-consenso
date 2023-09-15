package validator_vote

import (
	"fmt"

	"blion-worker-consenso/internal/logger"
	"blion-worker-consenso/internal/models"
	"github.com/asaskevich/govalidator"
)

type PortsServerValidatorVote interface {
	CreateValidatorVote(id string, participantsId string, hash string, vote bool) (*ValidatorVote, int, error)
	UpdateValidatorVote(id string, participantsId string, hash string, vote bool) (*ValidatorVote, int, error)
	DeleteValidatorVote(id string) (int, error)
	GetValidatorVoteByID(id string) (*ValidatorVote, int, error)
	GetAllValidatorVotes() ([]*ValidatorVote, error)
	GetAllValidatorVoteByLotteryID(lotteryID string) ([]*ValidatorVote, error)
	GetVotesInFavorByLotteryId(id string) (int64, error)
}

type service struct {
	repository ServicesValidatorVoteRepository
	user       *models.User
	txID       string
}

func NewValidatorVoteService(repository ServicesValidatorVoteRepository, user *models.User, TxID string) PortsServerValidatorVote {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateValidatorVote(id string, participantsId string, hash string, vote bool) (*ValidatorVote, int, error) {
	m := NewValidatorVotes(id, participantsId, hash, vote)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create ValidatorVote :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateValidatorVote(id string, participantsId string, hash string, vote bool) (*ValidatorVote, int, error) {
	m := NewValidatorVotes(id, participantsId, hash, vote)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update ValidatorVote :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteValidatorVote(id string) (int, error) {
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

func (s *service) GetValidatorVoteByID(id string) (*ValidatorVote, int, error) {
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

func (s *service) GetAllValidatorVotes() ([]*ValidatorVote, error) {
	return s.repository.getAll()
}

func (s *service) GetAllValidatorVoteByLotteryID(lotteryID string) ([]*ValidatorVote, error) {
	return s.repository.getByLotteryID(lotteryID)
}

func (s *service) GetVotesInFavorByLotteryId(id string) (int64, error) {
	if !govalidator.IsUUID(id) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return 0, fmt.Errorf("id isn't uuid")
	}
	m, err := s.repository.getVotesInFavor(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getVotesInFavor row:", err)
		return 0, err
	}
	return m, nil
}
