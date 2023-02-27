package graph

import (
	"github.com/teste-transfeera/internal/entity"
)

const TOTAL_PER_PAGE int = 10

func ToOutput(entity entity.Receiver) *Receiver {
	return &Receiver{
		ID:         entity.ID,
		Name:       entity.Name,
		Email:      entity.Email,
		Identifier: entity.Identifier,
		Pix: &Pix{
			KeyType: string(entity.Pix.KeyType),
			Key:     entity.Pix.Key,
		},
		Bank:    entity.Bank,
		Agency:  entity.Agency,
		Account: entity.Account,
		Status:  (*string)(&entity.Status),
	}
}

func BuildFilter(status *string, name *string, keyType *string, key *string) map[string]string {
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
