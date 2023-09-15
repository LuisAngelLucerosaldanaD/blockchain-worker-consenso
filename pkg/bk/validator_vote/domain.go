package validator_vote

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Model estructura de ValidatorVote
type ValidatorVote struct {
	ID            string    `json:"id" db:"id" valid:"required,uuid"`
	ParticipantId string    `json:"participant_id" db:"participant_id" valid:"required"`
	Hash          string    `json:"hash" db:"hash" valid:"required"`
	Vote          bool      `json:"vote" db:"vote" valid:"required"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

func NewValidatorVotes(id string, participantId string, hash string, vote bool) *ValidatorVote {
	return &ValidatorVote{
		ID:            id,
		ParticipantId: participantId,
		Hash:          hash,
		Vote:          vote,
	}
}

func (m *ValidatorVote) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
