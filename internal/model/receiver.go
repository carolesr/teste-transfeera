package model

import (
	"time"

	"github.com/teste-transfeera/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Receiver struct {
	ID         primitive.ObjectID `bson:"_id"`
	Identifier string             `bson:"identifier"`
	Name       string             `bson:"name"`
	Email      string             `bson:"email"`
	Pix        Pix                `bson:"pix"`
	Bank       *string            `bson:"bank,omitempty"`
	Agency     *string            `bson:"agency,omitempty"`
	Account    *string            `bson:"account,omitempty"`
	Status     string             `bson:"status"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at,omitempty"`
	DeletedAt  time.Time          `bson:"deleted_at,omitempty"`
}

func (m *Receiver) ToEntity() entity.Receiver {
	return entity.Receiver{
		ID:         m.ID.Hex(),
		Identifier: m.Identifier,
		Name:       m.Name,
		Email:      m.Email,
		Pix: entity.Pix{
			KeyType: entity.PixKeyType(m.Pix.KeyType),
			Key:     m.Pix.Key,
		},
		Bank:    m.Bank,
		Agency:  m.Agency,
		Account: m.Account,
		Status:  (entity.Status)(m.Status),
	}
}
