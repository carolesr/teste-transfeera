package usecase

import (
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
	receiver := entity.Receiver{
		Identifier: input.Identifier,
		Name:       input.Name,
		Email:      input.Email,
		Pix: entity.Pix{
			KeyType: entity.PixKeyType(input.PixKeyType),
			Key:     input.PixKey,
		},
	}

	newReceiver, err := u.receiverRepository.Create(receiver)
	if err != nil {
		return nil, err
	}

	return newReceiver, nil
}
