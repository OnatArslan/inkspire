package model

import "time"

// Domain model
type User struct {
	ID        string
	Email     string
	Password  string
	CreatedAt time.Time
}
