package usecase

import (
	"github.com/teste-transfeera/internal/entity"
)

func (u *receiverUseCase) List(filter map[string]string) ([]entity.Receiver, error) {
	receivers, err := u.receiverRepository.List(filter)
	if err != nil {
		return nil, err
	}

	return receivers, nil
}
