package server

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/diSpector/activity.git/internal/server/usecase"
	"github.com/diSpector/activity.git/pkg/activity/entities"
	pb "github.com/diSpector/activity.git/pkg/activity/grpc"
	"github.com/pkg/errors"
)

type Server struct {
	url string
	use usecase.ActivityUseCase
	pb.UnimplementedActivityApiServer
}

func New(url string, use usecase.ActivityUseCase) *Server {
	return &Server{
		url: url,
		use: use,
	}
}

func (s *Server) GetActivity(ctx context.Context, empty *pb.Empty) (*pb.Activity, error) {
	req, err := http.NewRequest("GET", s.url, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var act entities.Activity
	err = json.Unmarshal(body, &act)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	actId, err := s.use.Save(&act)
	if err != nil {
		log.Println(`activity not saved in db:`, err)
	} else {
		log.Println(`saved new activity:`, actId)
	}

	return &pb.Activity{
		Activity:     act.Activity,
		Participants: act.Participants,
	}, nil
}

func (s *Server) GetActivityStream(empty *pb.Empty, stream pb.ActivityApi_GetActivityStreamServer) error {
	req, err := http.NewRequest("GET", s.url, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	var act entities.Activity
	err = json.Unmarshal(body, &act)
	if err != nil {
		return errors.WithStack(err)
	}

	words := strings.Split(act.Activity, ` `)

	for _, word := range words {
		err := stream.Send(&pb.Description{Text: word})
		if err != nil {
			return err
		}
		time.Sleep(200 * time.Millisecond)
	}

	return nil
}

func (s *Server) AddActivity(ctx context.Context, activity *pb.Activity) (*pb.Empty, error) {
	act := &entities.Activity{
		Activity:     activity.Activity,
		Participants: activity.Participants,
	}

	actId, err := s.use.Save(act)
	if err != nil {
		return nil, err
	}

	if err != nil {
		log.Println(`activity not saved in db:`, err)
	} else {
		log.Println(`saved new activity:`, actId)
	}

	return &pb.Empty{}, nil
}

func (s *Server) SearchActivities(desc *pb.Description, stream pb.ActivityApi_SearchActivitiesServer) error {
	acts, err := s.use.Search(desc.Text)
	if err != nil {
		return err
	}

	for i := range acts {
		act := pb.Activity{
			Activity:     acts[i].Activity,
			Participants: acts[i].Participants,
		}
		err := stream.Send(&act)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) ListActivities(_ *pb.Empty, stream pb.ActivityApi_ListActivitiesServer) error {
	acts, err := s.use.List()
	if err != nil {
		return err
	}

	for i := range acts {
		act := pb.Activity{
			Activity:     acts[i].Activity,
			Participants: acts[i].Participants,
		}
		err := stream.Send(&act)
		if err != nil {
			return err
		}
	}

	return nil
}
