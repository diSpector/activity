package usecase

import (
	"testing"

	"github.com/diSpector/activity.git/internal/server/repository/mocks"
	"github.com/diSpector/activity.git/pkg/activity/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestActivitySave(t *testing.T) {
	mockRepo := mocks.NewMockActivityRepository(t)

	call := mockRepo.EXPECT().Insert(mock.Anything).Return(51, nil)

	useCase := NewActivityUseCaseImpl(mockRepo)
	id, err := useCase.Save(&entities.Activity{})

	require.Nil(t, err)
	assert.Equal(t, int64(51), id)

	call.Unset()
	mockRepo.EXPECT().Insert(mock.Anything).Return(52, nil)
	id, err = useCase.Save(&entities.Activity{})

	require.Nil(t, err)
	assert.Equal(t, int64(52), id)
}

func TestActivityList(t *testing.T) {
	mockRepo := mocks.NewMockActivityRepository(t)
	mockRepo.EXPECT().SelectAll().Return([]*entities.Activity{
		{
			Id:           1,
			Activity:     `Learn testing`,
			Participants: 5,
		},
		{
			Id:           2,
			Activity:     `Learn Golang`,
			Participants: 2,
		},
	}, nil)

	usecase := NewActivityUseCaseImpl(mockRepo)
	acts, err := usecase.List()

	require.Nil(t, err)
	assert.Equal(t, 2, len(acts))
}

func TestActivitySelectByName(t *testing.T) {
	mockRepo := mocks.NewMockActivityRepository(t)
	mockRepo.EXPECT().SelectByName(`clean`).Return([]*entities.Activity{
		{
			Id:           1,
			Activity:     `Clean out the room`,
			Participants: 5,
		},
	}, nil)

	mockRepo.EXPECT().SelectByName(`learn`).Return([]*entities.Activity{
		{
			Id:           2,
			Activity:     `Learn Golang`,
			Participants: 2,
		},
	}, nil)

	usecase := NewActivityUseCaseImpl(mockRepo)

	acts, err := usecase.Search(`clean`)
	require.Nil(t, err)
	assert.Equal(t, 1, len(acts))

	acts, err = usecase.Search(`learn`)

	require.Nil(t, err)
	assert.Equal(t, int32(2), acts[0].Participants)

	// acts, err = usecase.Search(`rrrrrrrrrrr`)
	// require.Nil(t, err)
	// assert.Equal(t, int32(2), acts[0].Participants)
}
