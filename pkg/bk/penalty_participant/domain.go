package penalty_participant

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Model estructura de PenaltyParticipant
type PenaltyParticipant struct {
	ID                string    `json:"id" db:"id" valid:"required,uuid"`
	ParticipantId     string    `json:"participant_id" db:"participant_id" valid:"required"`
	Amount            float64   `json:"amount" db:"amount" valid:"required"`
	PenaltyPercentage float64   `json:"penalty_percentage" db:"penalty_percentage" valid:"required"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

func NewPenaltyParticipants(id string, participantsId string, amount float64, penaltyPercentage float64) *PenaltyParticipant {
	return &PenaltyParticipant{
		ID:                id,
		ParticipantId:     participantsId,
		Amount:            amount,
		PenaltyPercentage: penaltyPercentage,
	}
}

func (m *PenaltyParticipant) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
