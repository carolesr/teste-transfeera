package usecase

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type DeleteReceiverInput struct {
	Ids []string `validate:"required"`
}

func (u *receiverUseCase) Delete(input *DeleteReceiverInput) error {
	err := validator.New().Struct(input)
	if err != nil {
		return err
	}

	if len(input.Ids) == 0 {
		return errors.New("At leat one id is required to delete receiver")
	}

	err = u.receiverRepository.Delete(input.Ids)
	if err != nil {
		return err
	}

	return nil
}
