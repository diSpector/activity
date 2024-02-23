package server

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/diSpector/activity.git/pkg/activity"
	pb "github.com/diSpector/activity.git/pkg/activity/grpc"
	"github.com/pkg/errors"
)

type Server struct {
	Url string
	pb.UnimplementedActivityApiServer
}

func New(url string) *Server {
	return &Server{Url: url}
}

func (s *Server) GetActivity(ctx context.Context, empty *pb.Empty) (*pb.Activity, error) {
	req, err := http.NewRequest("GET", s.Url, nil)
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

	var act activity.Activity
	err = json.Unmarshal(body, &act)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &pb.Activity{
		Activity:      act.Activity,
		Type:          act.Type,
		Participants:  act.Participants,
		Price:         act.Price,
		Link:          act.Link,
		Key:           act.Key,
		Accessibility: act.Accessibility,
	}, nil

}
