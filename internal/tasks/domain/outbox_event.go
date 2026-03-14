package domain

import "time"

type OutboxEvent struct {
	ID          uint64
	Aggregate   string
	AggregateID string
	Type        string
	Payload     string
	CreatedAt   time.Time
	Processed   bool
}
