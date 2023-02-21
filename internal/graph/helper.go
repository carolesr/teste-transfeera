package graph

import (
	"encoding/base64"

	"github.com/teste-transfeera/internal/entity"
)

const TOTAL_PER_PAGE int = 10

func toOutput(entity entity.Receiver) *Receiver {
	return &Receiver{
		ID:         entity.ID,
		Name:       entity.Name,
		Email:      entity.Email,
		Identifier: entity.Identifier,
		Pix: &Pix{
			KeyType: string(entity.Pix.KeyType),
			Key:     entity.Pix.Key,
		},
	}
}

func decodeBase64(cursor string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func encodeBase64(cursor []byte) string {
	return base64.StdEncoding.EncodeToString(cursor)
}
