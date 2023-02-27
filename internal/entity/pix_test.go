package entity_test

import (
	"errors"
	"testing"

	"github.com/teste-transfeera/internal/entity"
	"gopkg.in/stretchr/testify.v1/assert"
)

func Test_Pix_KeyType_Success(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should get pix CPF Key Type successfully", func(t *testing.T) {
		keyType, err := entity.GetKeyType("CPF")
		assert.Empty(err)
		assert.Equal(keyType, entity.CPF)
	})

	t.Run("Should get pix CNPJ Key Type successfully", func(t *testing.T) {
		keyType, err := entity.GetKeyType("CNPJ")
		assert.Empty(err)
		assert.Equal(keyType, entity.CNPJ)
	})
	t.Run("Should get pix Email Key Type successfully", func(t *testing.T) {
		keyType, err := entity.GetKeyType("EMAIL")
		assert.Empty(err)
		assert.Equal(keyType, entity.Email)
	})
	t.Run("Should get pix Phone Key Type successfully", func(t *testing.T) {
		keyType, err := entity.GetKeyType("TELEFONE")
		assert.Empty(err)
		assert.Equal(keyType, entity.Phone)
	})
	t.Run("Should get pix RandomKey Key Type successfully", func(t *testing.T) {
		keyType, err := entity.GetKeyType("CHAVE_ALEATORIA")
		assert.Empty(err)
		assert.Equal(keyType, entity.RandomKey)
	})
}

func Test_Pix_KeyType_Error(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should get error with wrong key type input", func(t *testing.T) {
		keyType, err := entity.GetKeyType("cpf")
		assert.Equal(err, errors.New("Type not found"))
		assert.Equal(keyType, entity.PixKeyType(""))
	})
}
