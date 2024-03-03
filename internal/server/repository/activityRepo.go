package repository

import "github.com/diSpector/activity.git/pkg/activity/entities"

//go:generate mockery
type ActivityRepository interface {
	Insert(*entities.Activity) (int64, error)
	SelectByName(name string) ([]*entities.Activity, error)
	SelectAll() ([]*entities.Activity, error)
}
