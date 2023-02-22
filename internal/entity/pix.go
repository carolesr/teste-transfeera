package entity

import "fmt"

type PixKeyType string

const (
	CPF       PixKeyType = "CPF"
	CNPJ      PixKeyType = "CNPJ"
	Email     PixKeyType = "EMAIL"
	Phone     PixKeyType = "TELEFONE"
	RandomKey PixKeyType = "CHAVE_ALEATORIA"
)

var mapKeyType = map[string]PixKeyType{"CPF": CPF, "CNPJ": CNPJ, "EMAIL": Email, "TELEFONE": Phone, "CHAVE_ALEATORIA": RandomKey}

type Pix struct {
	KeyType PixKeyType
	Key     string
}

func GetKeyType(keyType string) (PixKeyType, error) {
	t, ok := mapKeyType[keyType]
	if !ok {
		err := fmt.Errorf("Type not found")
		return "", err
	}
	return t, nil
}
