package validator_votes

import (
	"database/sql"
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

func newValidatorVotesPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *ValidatorVotes) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const psqlInsert = `INSERT INTO bc.validator_votes (id ,lottery_id, participants_id, hash, vote, created_at, updated_at) VALUES (:id ,:lottery_id, :participants_id, :hash, :vote,:created_at, :updated_at) `
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
func (s *psql) update(m *ValidatorVotes) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE bc.validator_votes SET lottery_id = :lottery_id, participants_id = :participants_id, hash = :hash, vote = :vote, updated_at = :updated_at WHERE id = :id `
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
	const psqlDelete = `DELETE FROM bc.validator_votes WHERE id = :id `
	m := ValidatorVotes{ID: id}
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
func (s *psql) getByID(id string) (*ValidatorVotes, error) {
	const psqlGetByID = `SELECT id , lottery_id, participants_id, hash, vote, created_at, updated_at FROM bc.validator_votes WHERE id = $1 `
	mdl := ValidatorVotes{}
	err := s.DB.Get(&mdl, psqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *psql) getAll() ([]*ValidatorVotes, error) {
	var ms []*ValidatorVotes
	const psqlGetAll = ` SELECT id , lottery_id, participants_id, hash, vote, created_at, updated_at FROM bc.validator_votes `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

// GetAll consulta todos los registros de la BD
func (s *psql) getByLotteryID(lotteryID string) ([]*ValidatorVotes, error) {
	var ms []*ValidatorVotes
	const psqlGetAll = ` SELECT id , lottery_id, participants_id, hash, vote, created_at, updated_at FROM bc.validator_votes WHERE lottery_id = $1 `

	err := s.DB.Select(&ms, psqlGetAll, lotteryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *psql) getVotesInFavor(id string) (int64, error) {
	const psqlGetByID = `select sum(case when vt.vote = true then 1 else 0 end) from bc.validator_votes vt WHERE vt.lottery_id = $1 `
	var votes int64
	err := s.DB.Get(&votes, psqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return votes, nil
}
