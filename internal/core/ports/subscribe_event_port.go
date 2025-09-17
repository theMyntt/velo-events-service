package ports

import "gitlab.com/velo-company/services/events-service/internal/core/entities"

type SubscribeEventPort interface {
	Execute(table entities.UserParticipation) error
}
