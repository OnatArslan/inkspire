package model

import "time"

// Domain model
type User struct {
	ID        int64
	Email     string
	Password  string
	CreatedAt time.Time
}
