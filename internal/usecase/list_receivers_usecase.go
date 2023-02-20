package usecase

import (
	"github.com/teste-transfeera/internal/entity"
)

func (u *receiverUseCase) List() ([]entity.Receiver, error) {
	receivers, err := u.receiverRepository.List()
	if err != nil {
		return nil, err
	}

	return receivers, nil
}
