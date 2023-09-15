package penalty_participant

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"blion-worker-consenso/internal/models"
	"github.com/jmoiron/sqlx"
)

// psql estructura de conexión a la BD de postgresql
type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newPenaltyParticipantPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *PenaltyParticipant) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const psqlInsert = `INSERT INTO bc.penalty_participant (id, participant_id, amount, penalty_percentage, created_at, updated_at) VALUES (:id, :participant_id, :amount, :penalty_percentage,:created_at, :updated_at) `
	rs, err := s.DB.NamedExec(psqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psql) update(m *PenaltyParticipant) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE bc.penalty_participant SET participant_id = :participant_id, amount = :amount, penalty_percentage = :penalty_percentage, updated_at = :updated_at WHERE id = :id `
	rs, err := s.DB.NamedExec(psqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *psql) delete(id string) error {
	const psqlDelete = `DELETE FROM bc.penalty_participant WHERE id = :id `
	m := PenaltyParticipant{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *psql) getByID(id string) (*PenaltyParticipant, error) {
	const psqlGetByID = `SELECT id, participant_id, amount, penalty_percentage, created_at, updated_at FROM bc.penalty_participant WHERE id = $1 `
	mdl := PenaltyParticipant{}
	err := s.DB.Get(&mdl, psqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetByID consulta un registro por su ID
func (s *psql) getByWalletID(walletId string) ([]*PenaltyParticipant, error) {
	const psqlGetByWalletID = `SELECT pp.id, pp.participant_id, pp.amount, pp.penalty_percentage, pp.created_at, pp.updated_at FROM bc.penalty_participant pp
join bc.participant p on(p.id = pp.participant_id) WHERE p.wallet_id = '%s'`
	var ms []*PenaltyParticipant
	err := s.DB.Select(&ms, fmt.Sprintf(psqlGetByWalletID, walletId))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

// GetAll consulta todos los registros de la BD
func (s *psql) getAll() ([]*PenaltyParticipant, error) {
	var ms []*PenaltyParticipant
	const psqlGetAll = ` SELECT id, participant_id, amount, penalty_percentage, created_at, updated_at FROM bc.penalty_participant `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
