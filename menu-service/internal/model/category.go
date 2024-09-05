package model

import "time"

type CategoryStatus string

var (
	CategoryStatusActive   CategoryStatus = "ACTIVE"
	CategoryStatusInactive CategoryStatus = "INACTIVE"
)

func (c CategoryStatus) String() string {
	return string(c)
}

type Category struct {
	ID          int64
	Name        string
	Description string
	Status      CategoryStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
