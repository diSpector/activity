package usecase

import (
	"github.com/diSpector/activity.git/internal/server/repository"
	"github.com/diSpector/activity.git/pkg/activity/entities"
)

type ActivityUseCaseImpl struct {
	repo repository.ActivityRepository
}

func NewActivityUseCaseImpl(repo repository.ActivityRepository) *ActivityUseCaseImpl {
	return &ActivityUseCaseImpl{
		repo: repo,
	}
}

func (s *ActivityUseCaseImpl) Generate() (*entities.Activity, error) {
	return nil, nil
}

func (s *ActivityUseCaseImpl) Save(activity *entities.Activity) (int64, error) {
	return s.repo.Insert(activity)
}

func (s *ActivityUseCaseImpl) Search(name string) ([]*entities.Activity, error) {
	return s.repo.SelectByName(name)
}

func (s *ActivityUseCaseImpl) List() ([]*entities.Activity, error) {
	return s.repo.SelectAll()
}
