package usecase

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/teste-transfeera/internal/entity"
)

func validatorIdentifier(fl validator.FieldLevel) bool {
	return validateIdentifier(fl.Field().String())
}

func validatorEmail(fl validator.FieldLevel) bool {
	return validateEmail(fl.Field().String())
}

func validatorPixType(fl validator.FieldLevel) bool {
	return validatePixType(fl.Field().String())
}

func validatorPixKey(fl validator.FieldLevel) bool {
	return validatePixKey(fl.Field().String(), fl.Parent().FieldByName("PixKeyType").String())
}

func validateIdentifier(identifier string) bool {
	pattern := regexp.MustCompile(`^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}|[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`)
	return pattern.MatchString(identifier)
}

func validateEmail(email string) bool {
	pattern := regexp.MustCompile(`^[A-Z0-9+_.-]+@[A-Z0-9.-]+$`)
	return pattern.MatchString(email)
}

func validatePixType(keyType string) bool {
	if _, err := entity.GetKeyType(keyType); err != nil {
		return false
	}

	return true
}

func validatePixKey(key string, keyTypeStr string) bool {
	keyType, err := entity.GetKeyType(keyTypeStr)
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
		return validateEmail(key)

	case entity.Phone:
		pattern := regexp.MustCompile(`^((?:\+?55)?)([1-9][0-9])(9[0-9]{8})$`)
		return pattern.MatchString(key)

	case entity.RandomKey:
		pattern := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i`)
		return pattern.MatchString(key)
	}

	return false
}
