package ports

type CancelSubscriptionPort interface {
	Execute(eventId int) error
}
