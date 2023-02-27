package usecase

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/teste-transfeera/internal/entity"
	"github.com/teste-transfeera/pkg/validation"
)

type UpdateReceiverInput struct {
	Id         string `validate:"required"`
	Identifier string `validate:"omitempty,validateIdentifier"`
	Name       string `validate:"omitempty"`
	Email      string `validate:"omitempty,validateEmail"`
	PixKeyType string `validate:"omitempty"`
	PixKey     string `validate:"omitempty"`
}

func (u *receiverUseCase) Update(input *UpdateReceiverInput) error {
	validator := validator.New()
	validator.RegisterValidation("validateIdentifier", validation.ValidatorIdentifier)
	validator.RegisterValidation("validateEmail", validation.ValidatorEmail)

	err := validator.Struct(input)
	if err != nil {
		return err
	}

	receiver, err := u.receiverRepository.FindById(input.Id)
	if err != nil {
		return err
	}

	if err := validatePix(input, receiver); err != nil {
		return err
	}

	fieldsToUpdate := buildUpdateByStatus(receiver.Status, input)
	if len(fieldsToUpdate) == 0 {
		return errors.New("Required at least one field to be updated")
	}

	err = u.receiverRepository.Update(input.Id, fieldsToUpdate)
	if err != nil {
		return err
	}

	return nil
}

func validatePix(input *UpdateReceiverInput, receiver *entity.Receiver) error {
	bothKeyAndTypeChanged := input.PixKeyType != "" && input.PixKey != ""
	if bothKeyAndTypeChanged {
		if !validation.ValidatePixType(input.PixKeyType) {
			return errors.New("Invalid Pix Key Type")
		}
		if !validation.ValidatePixKey(input.PixKey, input.PixKeyType) {
			return errors.New(fmt.Sprintf("Invalid Pix Key for %s Key Type", input.PixKeyType))
		}
		return nil
	}

	onlyKeyChanged := input.PixKey != ""
	if onlyKeyChanged {
		if !validation.ValidatePixKey(input.PixKey, string(receiver.Pix.KeyType)) {
			return errors.New(fmt.Sprintf("Invalid Pix Key for %s Key Type", receiver.Pix.KeyType))
		}
		return nil
	}

	onlyTypeChanged := input.PixKeyType != ""
	if onlyTypeChanged {
		return errors.New("Updating Pix Key Type requires also updating Pix Key")
	}
	return nil
}

func buildUpdateByStatus(status entity.Status, input *UpdateReceiverInput) map[string]string {
	fieldsToUpdate := make(map[string]string)
	switch status {
	case entity.Validated:
		if input.Email != "" {
			fieldsToUpdate["email"] = input.Email
		}
	case entity.Draft:
		if input.Identifier != "" {
			fieldsToUpdate["identifier"] = input.Identifier
		}
		if input.Name != "" {
			fieldsToUpdate["name"] = input.Name
		}
		if input.Email != "" {
			fieldsToUpdate["email"] = input.Email
		}
		if input.PixKey != "" {
			fieldsToUpdate["key"] = input.PixKey
		}
		if input.PixKeyType != "" {
			fieldsToUpdate["key_type"] = input.PixKeyType
		}
	}
	return fieldsToUpdate
}
