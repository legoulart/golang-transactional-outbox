package port

import (
	"github.com/legoulart/poc-go/internal/adapter"
	"github.com/legoulart/poc-go/internal/domain"
)

type Port interface {
	Create(subscription domain.Subscription)
}

type subscriptionUseCase struct {
	dataAccessPort adapter.DataAccessPort
}

func New(port adapter.DataAccessPort) Port {
	return subscriptionUseCase{dataAccessPort: port}
}

func (u subscriptionUseCase) Create(subscription domain.Subscription) {
	u.dataAccessPort.Save(subscription)
}
