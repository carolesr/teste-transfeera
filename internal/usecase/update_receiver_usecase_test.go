package usecase_test

import (
	"errors"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/teste-transfeera/internal/entity"
	"github.com/teste-transfeera/internal/usecase"
	"github.com/teste-transfeera/mocks"
)

func Test_ReceiverUseCase_Update_Success(t *testing.T) {
	repository := &mocks.ReceiverRepository{}
	useCase := usecase.NewReceiverUseCases(repository)

	t.Run("Update all fields from receiver successfully", func(t *testing.T) {
		input := usecase.UpdateReceiverInput{
			Id:         "63f8c8d6c6ce914b5b00b88e",
			Identifier: "222.222.222-222",
			Name:       "Receiver 2",
			Email:      "RECEIVER2@GMAIL.COM",
			PixKeyType: "EMAIL",
			PixKey:     "RECEIVER2@GMAIL.COM",
		}
		mockOutput := &entity.Receiver{
			ID:         input.Id,
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			Pix: entity.Pix{
				KeyType: entity.CPF,
				Key:     "111.111.111-11",
			},
			Status: entity.Draft,
		}
		fieldsToUpdate := map[string]string{
			"identifier": input.Identifier,
			"name":       input.Name,
			"email":      input.Email,
			"key_type":   input.PixKeyType,
			"key":        input.PixKey,
		}
		repository.On("FindById", input.Id).Return(mockOutput, nil).Once()
		repository.On("Update", input.Id, fieldsToUpdate).Return(nil).Once()

		err := useCase.Update(&input)

		assert.Equal(t, nil, err)
		repository.AssertExpectations(t)
	})

	t.Run("Update only Pix Key from receiver successfully", func(t *testing.T) {
		input := usecase.UpdateReceiverInput{
			Id:     "63f8c8d6c6ce914b5b00b88e",
			PixKey: "222.222.222-22",
		}
		mockOutput := &entity.Receiver{
			ID:         input.Id,
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			Pix: entity.Pix{
				KeyType: entity.CPF,
				Key:     "111.111.111-11",
			},
			Status: entity.Draft,
		}
		fieldsToUpdate := map[string]string{
			"key": input.PixKey,
		}
		repository.On("FindById", input.Id).Return(mockOutput, nil).Once()
		repository.On("Update", input.Id, fieldsToUpdate).Return(nil).Once()

		err := useCase.Update(&input)

		assert.Equal(t, nil, err)
		repository.AssertExpectations(t)
	})

	t.Run("Update only Email from receiver with Validated status successfully", func(t *testing.T) {
		input := usecase.UpdateReceiverInput{
			Id:    "63f8c8d6c6ce914b5b00b88e",
			Name:  "Receiver 2",
			Email: "RECEIVER2@GMAIL.COM",
		}
		mockOutput := &entity.Receiver{
			ID:         input.Id,
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			Pix: entity.Pix{
				KeyType: entity.CPF,
				Key:     "111.111.111-11",
			},
			Status: entity.Validated,
		}
		fieldsToUpdate := map[string]string{
			"email": input.Email,
		}
		repository.On("FindById", input.Id).Return(mockOutput, nil).Once()
		repository.On("Update", input.Id, fieldsToUpdate).Return(nil).Once()

		err := useCase.Update(&input)

		assert.Equal(t, nil, err)
		repository.AssertExpectations(t)
	})
}

func Test_ReceiverUseCase_Update_Error(t *testing.T) {
	repository := &mocks.ReceiverRepository{}
	useCase := usecase.NewReceiverUseCases(repository)

	t.Run("Update receiver returns error from repository on Update", func(t *testing.T) {
		input := usecase.UpdateReceiverInput{
			Id:     "63f8c8d6c6ce914b5b00b88e",
			PixKey: "222.222.222-22",
		}
		mockOutput := &entity.Receiver{
			ID:         input.Id,
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			Pix: entity.Pix{
				KeyType: entity.CPF,
				Key:     "111.111.111-11",
			},
			Status: entity.Draft,
		}
		fieldsToUpdate := map[string]string{
			"key": input.PixKey,
		}
		expectedError := errors.New("error")
		repository.On("FindById", input.Id).Return(mockOutput, nil).Once()
		repository.On("Update", input.Id, fieldsToUpdate).Return(errors.New("error")).Once()

		err := useCase.Update(&input)

		assert.Equal(t, expectedError, err)
		repository.AssertExpectations(t)
	})

	t.Run("Update zero fields from receiver returns error", func(t *testing.T) {
		input := usecase.UpdateReceiverInput{
			Id: "63f8c8d6c6ce914b5b00b88e",
		}
		mockOutput := &entity.Receiver{
			ID:         input.Id,
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			Pix: entity.Pix{
				KeyType: entity.CPF,
				Key:     "111.111.111-11",
			},
			Status: entity.Draft,
		}
		expectedError := errors.New("Required at least one field to be updated")
		repository.On("FindById", input.Id).Return(mockOutput, nil).Once()

		err := useCase.Update(&input)

		assert.Equal(t, expectedError, err)
		repository.AssertExpectations(t)
	})

	t.Run("Update only Pix Key Type from receiver returns error", func(t *testing.T) {
		input := usecase.UpdateReceiverInput{
			Id:         "63f8c8d6c6ce914b5b00b88e",
			PixKeyType: "EMAIL",
		}
		mockOutput := &entity.Receiver{
			ID:         input.Id,
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			Pix: entity.Pix{
				KeyType: entity.CPF,
				Key:     "111.111.111-11",
			},
			Status: entity.Draft,
		}
		expectedError := errors.New("Updating Pix Key Type requires also updating Pix Key")
		repository.On("FindById", input.Id).Return(mockOutput, nil).Once()

		err := useCase.Update(&input)

		assert.Equal(t, expectedError, err)
		repository.AssertExpectations(t)
	})

	t.Run("Update Pix Key from receiver returns validation error", func(t *testing.T) {
		input := usecase.UpdateReceiverInput{
			Id:     "63f8c8d6c6ce914b5b00b88e",
			PixKey: "111.111.111-1",
		}
		mockOutput := &entity.Receiver{
			ID:         input.Id,
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			Pix: entity.Pix{
				KeyType: entity.CPF,
				Key:     "111.111.111-11",
			},
			Status: entity.Draft,
		}
		expectedError := errors.New("Invalid Pix Key for CPF Key Type")
		repository.On("FindById", input.Id).Return(mockOutput, nil).Once()

		err := useCase.Update(&input)

		assert.Equal(t, expectedError, err)
		repository.AssertExpectations(t)
	})

	t.Run("Update Pix from receiver returns validation error for Pix Key", func(t *testing.T) {
		input := usecase.UpdateReceiverInput{
			Id:         "63f8c8d6c6ce914b5b00b88e",
			PixKeyType: "EMAIL",
			PixKey:     "RECEIVER2@",
		}
		mockOutput := &entity.Receiver{
			ID:         input.Id,
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			Pix: entity.Pix{
				KeyType: entity.CPF,
				Key:     "111.111.111-11",
			},
			Status: entity.Draft,
		}
		expectedError := errors.New("Invalid Pix Key for EMAIL Key Type")
		repository.On("FindById", input.Id).Return(mockOutput, nil).Once()

		err := useCase.Update(&input)

		assert.Equal(t, expectedError, err)
		repository.AssertExpectations(t)
	})

	t.Run("Update Pix from receiver returns validation error for Pix Key Type", func(t *testing.T) {
		input := usecase.UpdateReceiverInput{
			Id:         "63f8c8d6c6ce914b5b00b88e",
			PixKeyType: "email",
			PixKey:     "RECEIVER2@GMAIL.COM",
		}
		mockOutput := &entity.Receiver{
			ID:         input.Id,
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			Pix: entity.Pix{
				KeyType: entity.CPF,
				Key:     "111.111.111-11",
			},
			Status: entity.Draft,
		}
		expectedError := errors.New("Invalid Pix Key Type")
		repository.On("FindById", input.Id).Return(mockOutput, nil).Once()

		err := useCase.Update(&input)

		assert.Equal(t, expectedError, err)
		repository.AssertExpectations(t)
	})

	t.Run("Update receiver returns error from repository on FindById", func(t *testing.T) {
		input := usecase.UpdateReceiverInput{
			Id:         "63f8c8d6c6ce914b5b00b88e",
			PixKeyType: "EMAIL",
			PixKey:     "RECEIVER2@GMAIL.COM",
		}
		expectedError := errors.New("error")
		repository.On("FindById", input.Id).Return(nil, errors.New("error")).Once()

		err := useCase.Update(&input)

		assert.Equal(t, expectedError, err)
		repository.AssertExpectations(t)
	})

	t.Run("Update receiver returns error validation error for Email", func(t *testing.T) {
		input := usecase.UpdateReceiverInput{
			Id:    "63f8c8d6c6ce914b5b00b88e",
			Email: "a",
		}
		expectedError := errors.New(`Key: 'UpdateReceiverInput.Email' Error:Field validation for 'Email' failed on the 'validateEmail' tag`)

		err := useCase.Update(&input)

		assert.Equal(t, expectedError.Error(), err.Error())
		repository.AssertExpectations(t)
	})

	t.Run("Update receiver returns error validation error for Identifier", func(t *testing.T) {
		input := usecase.UpdateReceiverInput{
			Id:         "63f8c8d6c6ce914b5b00b88e",
			Identifier: "a",
		}
		expectedError := errors.New(`Key: 'UpdateReceiverInput.Identifier' Error:Field validation for 'Identifier' failed on the 'validateIdentifier' tag`)

		err := useCase.Update(&input)

		assert.Equal(t, expectedError.Error(), err.Error())
		repository.AssertExpectations(t)
	})
}
