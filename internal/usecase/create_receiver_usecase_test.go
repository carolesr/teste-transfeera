package usecase_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/magiconair/properties/assert"
	"github.com/teste-transfeera/internal/entity"
	"github.com/teste-transfeera/internal/usecase"
	"github.com/teste-transfeera/mocks"
)

func Test_ReceiverUseCase_Create_Success(t *testing.T) {
	repository := &mocks.ReceiverRepository{}
	useCase := usecase.NewReceiverUseCases(repository)

	t.Run("Create receiver successfully", func(t *testing.T) {
		input := usecase.CreateReceiverInput{
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			PixKeyType: "CPF",
			PixKey:     "111.111.111-11",
		}
		mockInput := entity.Receiver{
			Identifier: input.Identifier,
			Name:       input.Name,
			Email:      input.Email,
			Status:     entity.Draft,
			Pix: entity.Pix{
				KeyType: entity.PixKeyType(input.PixKeyType),
				Key:     input.PixKey,
			},
		}
		expectedResult := &entity.Receiver{
			ID:         uuid.New().String(),
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			Status:     entity.Draft,
			Pix: entity.Pix{
				KeyType: entity.CPF,
				Key:     "111.111.111-11",
			},
		}
		repository.On("Create", mockInput).Return(expectedResult, nil).Once()

		result, err := useCase.Create(&input)

		assert.Equal(t, expectedResult, result)
		assert.Equal(t, nil, err)
		repository.AssertExpectations(t)
	})
}

func Test_ReceiverUseCase_Create_Error(t *testing.T) {
	repository := &mocks.ReceiverRepository{}
	useCase := usecase.NewReceiverUseCases(repository)

	t.Run("Create receiver returns error from repository", func(t *testing.T) {
		input := usecase.CreateReceiverInput{
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			PixKeyType: "CPF",
			PixKey:     "111.111.111-11",
		}
		mockInput := entity.Receiver{
			Identifier: input.Identifier,
			Name:       input.Name,
			Email:      input.Email,
			Status:     entity.Draft,
			Pix: entity.Pix{
				KeyType: entity.PixKeyType(input.PixKeyType),
				Key:     input.PixKey,
			},
		}
		expectedError := errors.New("error")
		repository.On("Create", mockInput).Return(nil, errors.New("error")).Once()

		result, err := useCase.Create(&input)

		assert.Equal(t, (*entity.Receiver)(nil), result)
		assert.Equal(t, expectedError, err)
		repository.AssertExpectations(t)
	})

	t.Run("Create receiver returns validation error for pix key", func(t *testing.T) {
		input := usecase.CreateReceiverInput{
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			PixKeyType: "EMAIL",
			PixKey:     "a@a",
		}
		expectedError := errors.New(`Key: 'CreateReceiverInput.PixKey' Error:Field validation for 'PixKey' failed on the 'validatePixKey' tag`)
		result, err := useCase.Create(&input)

		assert.Equal(t, (*entity.Receiver)(nil), result)
		assert.Equal(t, expectedError.Error(), err.Error())
		repository.AssertExpectations(t)
	})

	t.Run("Create receiver returns validation error for pix key type", func(t *testing.T) {
		input := usecase.CreateReceiverInput{
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			PixKeyType: "CPFF",
			PixKey:     "111.111.111-11",
		}
		expectedError := errors.New("Key: 'CreateReceiverInput.PixKeyType' Error:Field validation for 'PixKeyType' failed on the 'validatePixType' tag\nKey: 'CreateReceiverInput.PixKey' Error:Field validation for 'PixKey' failed on the 'validatePixKey' tag")
		result, err := useCase.Create(&input)

		assert.Equal(t, (*entity.Receiver)(nil), result)
		assert.Equal(t, expectedError.Error(), err.Error())
		repository.AssertExpectations(t)
	})

	t.Run("Create receiver returns validation error for email", func(t *testing.T) {
		input := usecase.CreateReceiverInput{
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "receiver1@gmail.com",
			PixKeyType: "EMAIL",
			PixKey:     "A@A",
		}
		expectedError := errors.New(`Key: 'CreateReceiverInput.Email' Error:Field validation for 'Email' failed on the 'validateEmail' tag`)
		result, err := useCase.Create(&input)

		assert.Equal(t, (*entity.Receiver)(nil), result)
		assert.Equal(t, expectedError.Error(), err.Error())
		repository.AssertExpectations(t)
	})

	t.Run("Create receiver returns validation error for identifier", func(t *testing.T) {
		input := usecase.CreateReceiverInput{
			Identifier: "111.111.111-1",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			PixKeyType: "EMAIL",
			PixKey:     "A@A",
		}
		expectedError := errors.New(`Key: 'CreateReceiverInput.Identifier' Error:Field validation for 'Identifier' failed on the 'validateIdentifier' tag`)
		result, err := useCase.Create(&input)

		assert.Equal(t, (*entity.Receiver)(nil), result)
		assert.Equal(t, expectedError.Error(), err.Error())
		repository.AssertExpectations(t)
	})

	t.Run("Create receiver returns validation error for name", func(t *testing.T) {
		input := usecase.CreateReceiverInput{
			Identifier: "111.111.111-11",
			Email:      "RECEIVER1@GMAIL.COM",
			PixKeyType: "EMAIL",
			PixKey:     "A@A",
		}
		expectedError := errors.New(`Key: 'CreateReceiverInput.Name' Error:Field validation for 'Name' failed on the 'required' tag`)
		result, err := useCase.Create(&input)

		assert.Equal(t, (*entity.Receiver)(nil), result)
		assert.Equal(t, expectedError.Error(), err.Error())
		repository.AssertExpectations(t)
	})

}
