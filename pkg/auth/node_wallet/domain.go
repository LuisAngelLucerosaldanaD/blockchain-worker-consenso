package node_wallet

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// NodeWallet  Model struct NodeWallet
type NodeWallet struct {
	ID        string     `json:"id" db:"id" valid:"required,uuid"`
	WalletId  string     `json:"wallet_id" db:"wallet_id" valid:"required"`
	Name      string     `json:"name" db:"name" valid:"required"`
	Ip        string     `json:"ip" db:"ip" valid:"required"`
	DeletedAt *time.time `json:"deleted_at" db:"deleted_at" valid:"required"`
	PenaltyAt *time.time `json:"penalty_at" db:"penalty_at" valid:"required"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
}

func NewNodeWallet(id string, walletId string, name string, ip string, deletedAt *time.time, penaltyAt *time.time) *NodeWallet {
	return &NodeWallet{
		ID:        id,
		WalletId:  walletId,
		Name:      name,
		Ip:        ip,
		DeletedAt: deletedAt,
		PenaltyAt: penaltyAt,
	}
}

func (m *NodeWallet) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
