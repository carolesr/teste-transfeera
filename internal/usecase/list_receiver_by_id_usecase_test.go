package usecase_test

import (
	"errors"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/teste-transfeera/internal/entity"
	"github.com/teste-transfeera/internal/usecase"
	"github.com/teste-transfeera/mocks"
)

func Test_ReceiverUseCase_ListById_Success(t *testing.T) {
	repository := &mocks.ReceiverRepository{}
	useCase := usecase.NewReceiverUseCases(repository)

	t.Run("List receiver by id successfully", func(t *testing.T) {
		input := usecase.ListReceiverByIdInput{
			Id: "63f8c8d6c6ce914b5b00b88e",
		}
		expectedResult := &entity.Receiver{
			ID:         input.Id,
			Identifier: "111.111.111-11",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			Status:     entity.Draft,
			Pix: entity.Pix{
				KeyType: entity.CPF,
				Key:     "111.111.111-11",
			},
		}
		repository.On("FindById", input.Id).Return(expectedResult, nil).Once()

		result, err := useCase.ListById(&input)

		assert.Equal(t, expectedResult, result)
		assert.Equal(t, nil, err)
		repository.AssertExpectations(t)
	})
}

func Test_ReceiverUseCase_ListById_Error(t *testing.T) {
	repository := &mocks.ReceiverRepository{}
	useCase := usecase.NewReceiverUseCases(repository)

	t.Run("List receiver by id returns error from repository", func(t *testing.T) {
		input := usecase.ListReceiverByIdInput{
			Id: "63f8c8d6c6ce914b5b00b88e",
		}
		expectedError := errors.New("error")
		repository.On("FindById", input.Id).Return(nil, errors.New("error")).Once()

		result, err := useCase.ListById(&input)

		assert.Equal(t, (*entity.Receiver)(nil), result)
		assert.Equal(t, expectedError, err)
		repository.AssertExpectations(t)
	})

	t.Run("List receiver by id returns validation error for id", func(t *testing.T) {
		input := usecase.ListReceiverByIdInput{}
		expectedError := errors.New(`Key: 'ListReceiverByIdInput.Id' Error:Field validation for 'Id' failed on the 'required' tag`)
		result, err := useCase.ListById(&input)

		assert.Equal(t, (*entity.Receiver)(nil), result)
		assert.Equal(t, expectedError.Error(), err.Error())
		repository.AssertExpectations(t)
	})

}
