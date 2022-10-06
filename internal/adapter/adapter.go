package adapter

import (
	"github.com/legoulart/poc-go/internal/domain"
)

type DataAccessPort interface {
	Save(subscription domain.Subscription)
}
