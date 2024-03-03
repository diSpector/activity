package usecase

import "github.com/diSpector/activity.git/pkg/activity/entities"

type ActivityUseCase interface {
	Generate() (*entities.Activity, error)
	Save(*entities.Activity) (int64, error)
	Search(text string) ([]*entities.Activity, error)
	List() ([]*entities.Activity, error)
}
