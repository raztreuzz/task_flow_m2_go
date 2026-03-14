package domain

import "time"

type Project struct {
	ID        uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
