package model

import "time"

type Comment struct {
	Id         string
	Content    string
	Created_at time.Time
}
