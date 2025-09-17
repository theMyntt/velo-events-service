package entities

type UserParticipation struct {
	ID                  *int                    `json:"id"`
	UserID              *int                    `json:"user_id"`
	EventID             *int                    `json:"event_id"`
	ParticipationStatus UserParticipationStatus `json:"participation_status"`
	Events              []Event                 `json:"events"`
}

type UserParticipationStatus int

const (
	Registered UserParticipationStatus = iota
	Participated
	Cancelled
)
