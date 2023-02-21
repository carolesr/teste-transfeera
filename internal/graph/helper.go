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

func buildFilter(status *string, name *string, keyType *string, key *string) map[string]string {
	filter := make(map[string]string)

	if status != nil {
		filter["status"] = *status
	}
	if name != nil {
		filter["name"] = *name
	}
	if keyType != nil {
		filter["key_type"] = *keyType
	}
	if key != nil {
		filter["key"] = *key
	}

	return filter
}
