package usecase

import (
	"github.com/teste-transfeera/internal/entity"
	"github.com/teste-transfeera/internal/repository"
)

type ReceiverUseCases interface {
	Create(input *CreateReceiverInput) (*entity.Receiver, error)
	List(filter map[string]string) ([]entity.Receiver, error)
}

type receiverUseCase struct {
	receiverRepository repository.ReceiverRepository
	receivers          []*entity.Receiver
}

func NewReceiverUseCases(repository repository.ReceiverRepository) ReceiverUseCases {
	return &receiverUseCase{
		receiverRepository: repository,
	}
}
