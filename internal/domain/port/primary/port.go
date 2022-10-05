package primary

import (
	"github.com/legoulart/poc-go/internal/domain/model"
)

type Port interface {
	Create(user model.User)
}
