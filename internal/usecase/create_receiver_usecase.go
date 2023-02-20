package usecase

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/teste-transfeera/internal/entity"
)

type CreateReceiverInput struct {
	ID         string  `json:"id"`
	Identifier string  `json:"identifier"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	PixKeyType string  `json:"pix_key_type"`
	PixKey     string  `json:"pix_key"`
	Bank       *string `json:"bank"`
	Agency     *string `json:"agency"`
	Account    *string `json:"account"`
	Status     *string `json:"status"`
}

func (u *receiverUseCase) Create(input *CreateReceiverInput) (*entity.Receiver, error) {
	fmt.Println("create receiver usecase")
	rand, _ := rand.Int(rand.Reader, big.NewInt(100))
	receiver := &entity.Receiver{
		ID:         fmt.Sprintf("T%d", rand),
		Identifier: input.Identifier,
		Name:       input.Name,
		Email:      input.Email,
		PixKeyType: input.PixKeyType,
		PixKey:     input.PixKey,
	}
	u.receivers = append(u.receivers, receiver)

	newReceiver, err := u.receiverRepository.Create(*receiver)
	if err != nil {
		return nil, err
	}

	return newReceiver, nil
}
