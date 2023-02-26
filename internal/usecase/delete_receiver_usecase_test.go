package usecase_test

import (
	"errors"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/teste-transfeera/internal/usecase"
	"github.com/teste-transfeera/mocks"
)

func Test_ReceiverUseCase_Delete_Success(t *testing.T) {
	repository := &mocks.ReceiverRepository{}
	useCase := usecase.NewReceiverUseCases(repository)

	t.Run("Delete receiver by id successfully", func(t *testing.T) {
		input := usecase.DeleteReceiverInput{
			Ids: []string{"63f8c8d6c6ce914b5b00b88e"},
		}
		repository.On("Delete", input.Ids).Return(nil).Once()

		err := useCase.Delete(&input)

		assert.Equal(t, nil, err)
		repository.AssertExpectations(t)
	})
}

func Test_ReceiverUseCase_Delete_Error(t *testing.T) {
	repository := &mocks.ReceiverRepository{}
	useCase := usecase.NewReceiverUseCases(repository)

	t.Run("Delete receiver by id returns error from repository", func(t *testing.T) {
		input := usecase.DeleteReceiverInput{
			Ids: []string{"63f8c8d6c6ce914b5b00b88e"},
		}
		expectedError := errors.New("error")
		repository.On("Delete", input.Ids).Return(errors.New("error")).Once()

		err := useCase.Delete(&input)

		assert.Equal(t, expectedError, err)
		repository.AssertExpectations(t)
	})

	t.Run("Delete receiver by id returns validation error for id", func(t *testing.T) {
		input := usecase.DeleteReceiverInput{
			Ids: []string{},
		}
		expectedError := errors.New("At leat one id is required to delete receiver")
		err := useCase.Delete(&input)

		assert.Equal(t, expectedError.Error(), err.Error())
		repository.AssertExpectations(t)
	})

	t.Run("Delete receiver by id returns validation error for id", func(t *testing.T) {
		input := usecase.DeleteReceiverInput{}
		expectedError := errors.New("Key: 'DeleteReceiverInput.Ids' Error:Field validation for 'Ids' failed on the 'required' tag")
		err := useCase.Delete(&input)

		assert.Equal(t, expectedError.Error(), err.Error())
		repository.AssertExpectations(t)
	})

}
