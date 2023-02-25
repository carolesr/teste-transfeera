package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/teste-transfeera/internal/entity"
)

type ListReceiverByIdInput struct {
	Id string `validate:"required"`
}

func (u *receiverUseCase) ListById(input *ListReceiverByIdInput) (*entity.Receiver, error) {
	err := validator.New().Struct(input)
	if err != nil {
		return nil, err
	}
	receiver, err := u.receiverRepository.FindById(input.Id)
	if err != nil {
		return nil, err
	}

	return receiver, nil
}
