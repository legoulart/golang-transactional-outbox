package secondary

import (
	"github.com/legoulart/poc-go/internal/domain/model"
)

type DataAccessPort interface {
	Save(user model.User)
	SaveEvent(event model.Event)
}
