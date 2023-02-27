package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/teste-transfeera/internal/entity"
	"github.com/teste-transfeera/pkg/validation"
)

type CreateReceiverInput struct {
	Identifier string `validate:"required,validateIdentifier"`
	Name       string `validate:"required"`
	Email      string `validate:"required,max=250,validateEmail"`
	PixKeyType string `validate:"required,validatePixType"`
	PixKey     string `validate:"required,validatePixKey"`
}

func (u *receiverUseCase) Create(input *CreateReceiverInput) (*entity.Receiver, error) {
	validator := validator.New()
	validator.RegisterValidation("validateIdentifier", validation.ValidatorIdentifier)
	validator.RegisterValidation("validateEmail", validation.ValidatorEmail)
	validator.RegisterValidation("validatePixType", validation.ValidatorPixType)
	validator.RegisterValidation("validatePixKey", validation.ValidatorPixKey)

	err := validator.Struct(input)
	if err != nil {
		return nil, err
	}

	keyType, _ := entity.GetKeyType(input.PixKeyType)
	receiver := entity.Receiver{
		Identifier: input.Identifier,
		Name:       input.Name,
		Email:      input.Email,
		Pix: entity.Pix{
			KeyType: keyType,
			Key:     input.PixKey,
		},
		Status: entity.Draft,
	}

	newReceiver, err := u.receiverRepository.Create(receiver)
	if err != nil {
		return nil, err
	}

	return newReceiver, nil
}
