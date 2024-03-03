package entities

import "time"

type Activity struct {
	Id           int64      `json:"id" db:"id"`
	Activity     string     `json:"activity" db:"description"`
	Participants int32      `json:"participants" db:"persons"`
	Created      *time.Time `json:"created,omitempty" db:"created"`
}
