package usecase

import (
	"github.com/legoulart/poc-go/internal/domain/model"
	"github.com/legoulart/poc-go/internal/domain/port/primary"
	"github.com/legoulart/poc-go/internal/domain/port/secondary"
)

type UserUseCase struct {
	dataAccessPort secondary.DataAccessPort
}

func NewUserUseCase(port secondary.DataAccessPort) primary.Port {
	return UserUseCase{dataAccessPort: port}
}

func (u UserUseCase) Create(user model.User) {
	u.dataAccessPort.Save(user)
	u.dataAccessPort.SaveEvent(model.EventFromUser(user))
}
