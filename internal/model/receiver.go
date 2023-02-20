package model

import "github.com/teste-transfeera/internal/entity"

type Receiver struct {
	ID         string  `bson:"_id"`
	Identifier string  `bson:"identifier"`
	Name       string  `bson:"name"`
	Email      string  `bson:"email"`
	PixKeyType string  `bson:"pix_key_type"`
	PixKey     string  `bson:"pix_key"`
	Bank       *string `bson:"bank"`
	Agency     *string `bson:"agency"`
	Account    *string `bson:"account"`
	Status     *string `bson:"status"`
}

func (m *Receiver) ToEntity() entity.Receiver {
	return entity.Receiver{
		ID:         m.ID,
		Identifier: m.Identifier,
		Name:       m.Name,
		Email:      m.Email,
		PixKeyType: m.PixKeyType,
		PixKey:     m.PixKey,
		Bank:       m.Bank,
		Agency:     m.Agency,
		Account:    m.Account,
		Status:     m.Status,
	}
}
