package participants

import (
	"database/sql"
	"fmt"

	"blion-worker-consenso/internal/models"
	"github.com/jmoiron/sqlx"
	"time"
)

// psql estructura de conexión a la BD de postgresql
type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newParticipantsPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *Participants) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const psqlInsert = `INSERT INTO bc.participants (id ,lottery_id, wallet_id, amount, accepted, type_charge, returned, created_at, updated_at) VALUES (:id ,:lottery_id, :wallet_id, :amount, :accepted, :type_charge, :returned,:created_at, :updated_at) `
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
func (s *psql) update(m *Participants) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE bc.participants SET lottery_id = :lottery_id, wallet_id = :wallet_id, amount = :amount, accepted = :accepted, type_charge = :type_charge, returned = :returned, updated_at = :updated_at WHERE id = :id `
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
	const psqlDelete = `DELETE FROM bc.participants WHERE id = :id `
	m := Participants{ID: id}
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
func (s *psql) getByID(id string) (*Participants, error) {
	const psqlGetByID = `SELECT id , lottery_id, wallet_id, amount, accepted, type_charge, returned, created_at, updated_at FROM bc.participants WHERE id = $1 `
	mdl := Participants{}
	err := s.DB.Get(&mdl, psqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetByID consulta un registro por su ID
func (s *psql) getByWalletID(walletId string) (*Participants, error) {
	const psqlGetByWalletID = `SELECT id , lottery_id, wallet_id, amount, accepted, type_charge, returned, created_at, updated_at FROM bc.participants WHERE wallet_id = $1 order by id desc limit 1`
	mdl := Participants{}
	err := s.DB.Get(&mdl, psqlGetByWalletID, walletId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s *psql) getByWalletIDAndLotteryID(walletId string, lotteryId string) (*Participants, error) {
	const psqlGetByWalletID = `SELECT id , lottery_id, wallet_id, amount, accepted, type_charge, returned, created_at, updated_at FROM bc.participants WHERE wallet_id = $1 AND lottery_id = $2 limit 1`
	mdl := Participants{}
	err := s.DB.Get(&mdl, psqlGetByWalletID, walletId, lotteryId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetByID consulta un registro por su ID
func (s *psql) getByLotteryID(lotteryId string) ([]*Participants, error) {
	var ms []*Participants
	const psqlGetAll = ` SELECT id , lottery_id, wallet_id, amount, accepted, type_charge, returned, created_at, updated_at FROM bc.participants WHERE wallet_id = '%s'`

	err := s.DB.Select(&ms, fmt.Sprintf(psqlGetAll, lotteryId))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

// GetAll consulta todos los registros de la BD
func (s *psql) getAll() ([]*Participants, error) {
	var ms []*Participants
	const psqlGetAll = ` SELECT id , lottery_id, wallet_id, amount, accepted, type_charge, returned, created_at, updated_at FROM bc.participants `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
