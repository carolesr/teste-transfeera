package usecase

import (
	"github.com/teste-transfeera/internal/entity"
	"github.com/teste-transfeera/internal/repository"
)

type ReceiverUseCases interface {
	Create(input *CreateReceiverInput) (*entity.Receiver, error)
	List(filter map[string]string) ([]entity.Receiver, error)
	ListById(input *ListReceiverByIdInput) (*entity.Receiver, error)
	Update(input *UpdateReceiverInput) error
	Delete(input *DeleteReceiverInput) error
}

type receiverUseCase struct {
	receiverRepository repository.ReceiverRepository
}

func NewReceiverUseCases(repository repository.ReceiverRepository) ReceiverUseCases {
	return &receiverUseCase{
		receiverRepository: repository,
	}
}
