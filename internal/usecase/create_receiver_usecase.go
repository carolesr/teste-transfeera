package usecase

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/teste-transfeera/internal/entity"
)

type CreateReceiverInput struct {
	Identifier string `validate:"required,validateIdentifier"`
	Name       string `validate:"required"`
	Email      string `validate:"required,validateEmail"`
	PixKeyType string `validate:"required,validatePixType"`
	PixKey     string `validate:"required,validatePixKey"`
	Bank       *string
	Agency     *string
	Account    *string
}

func (u *receiverUseCase) Create(input *CreateReceiverInput) (*entity.Receiver, error) {

	validator := validator.New()
	validator.RegisterValidation("validateIdentifier", validateIdentifier)
	validator.RegisterValidation("validateEmail", validateEmail)
	validator.RegisterValidation("validatePixType", validatePixType)
	validator.RegisterValidation("validatePixKey", validatePixKey)

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
		Bank:    input.Bank,
		Agency:  input.Agency,
		Account: input.Account,
		Status:  entity.Draft,
	}

	newReceiver, err := u.receiverRepository.Create(receiver)
	if err != nil {
		return nil, err
	}

	return newReceiver, nil
}

func validateIdentifier(fl validator.FieldLevel) bool {
	pattern := regexp.MustCompile(`^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}|[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`)
	return pattern.MatchString(fl.Field().String())
}

func validateEmail(fl validator.FieldLevel) bool {
	pattern := regexp.MustCompile(`^[A-Z0-9+_.-]+@[A-Z0-9.-]+$`)
	return pattern.MatchString(fl.Field().String())
}

func validatePixType(fl validator.FieldLevel) bool {
	if _, err := entity.GetKeyType(fl.Field().String()); err != nil {
		return false
	}

	return true
}

func validatePixKey(fl validator.FieldLevel) bool {
	key := fl.Field().String()
	keyType, err := entity.GetKeyType(fl.Parent().FieldByName("PixKeyType").String())
	if err != nil {
		return false
	}

	switch keyType {
	case entity.CPF:
		pattern := regexp.MustCompile(`^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}$`)
		return pattern.MatchString(key)

	case entity.CNPJ:
		pattern := regexp.MustCompile(`^[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`)
		return pattern.MatchString(key)

	case entity.Email:
		pattern := regexp.MustCompile(`^[A-Z0-9+_.-]+@[A-Z0-9.-]+$`)
		return pattern.MatchString(key)

	case entity.Phone:
		pattern := regexp.MustCompile(`^((?:\+?55)?)([1-9][0-9])(9[0-9]{8})$`)
		return pattern.MatchString(key)

	case entity.RandomKey:
		pattern := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i`)
		return pattern.MatchString(key)
	}

	return false
}
