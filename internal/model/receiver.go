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
	Bank       *string            `bson:"bank"`
	Agency     *string            `bson:"agency"`
	Account    *string            `bson:"account"`
	Status     string             `bson:"status"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	DeletedAt  time.Time          `bson:"deleted_at"`
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
