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

func Test_ReceiverUseCase_List_Success(t *testing.T) {
	repository := &mocks.ReceiverRepository{}
	useCase := usecase.NewReceiverUseCases(repository)

	t.Run("List all receivers successfully", func(t *testing.T) {
		input := map[string]string{}
		expectedResult := []entity.Receiver{
			{
				ID:         uuid.New().String(),
				Identifier: "111.111.111-11",
				Name:       "Receiver 1",
				Email:      "receiver1@gmail.com",
				Status:     entity.Draft,
				Pix: entity.Pix{
					KeyType: entity.CPF,
					Key:     "111.111.111-11",
				},
			},
			{
				ID:         uuid.New().String(),
				Identifier: "222.222.222-22",
				Name:       "Receiver 2",
				Email:      "receiver2@gmail.com",
				Status:     entity.Validated,
				Pix: entity.Pix{
					KeyType: entity.Email,
					Key:     "receiver2@gmail.com",
				},
			},
		}
		repository.On("List", input).Return(expectedResult, nil).Once()

		result, err := useCase.List(input)

		assert.Equal(t, expectedResult, result)
		assert.Equal(t, nil, err)
		repository.AssertExpectations(t)
	})
}

func Test_ReceiverUseCase_List_Error(t *testing.T) {
	repository := &mocks.ReceiverRepository{}
	useCase := usecase.NewReceiverUseCases(repository)

	t.Run("List all receivers returns error from repository", func(t *testing.T) {
		input := map[string]string{}
		repository.On("List", input).Return(nil, errors.New("error")).Once()
		expectedError := errors.New("error")

		result, err := useCase.List(input)

		assert.Equal(t, []entity.Receiver(nil), result)
		assert.Equal(t, expectedError, err)
		repository.AssertExpectations(t)
	})
}
