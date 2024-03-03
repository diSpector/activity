package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/diSpector/activity.git/internal/server/repository"
	"github.com/diSpector/activity.git/internal/server/usecase"
	"github.com/diSpector/activity.git/pkg/activity/entities"
) 

func TestGetActivity(t *testing.T) {
	actRepo, err := repository.NewActivitySqlLiteRepo(`../../storage/sqlite.db`)
	if err != nil {
		log.Fatalln(`err create repo:`, err)
	}

	actUseCase := usecase.NewActivityUseCaseImpl(actRepo)

	testHttpServ := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		act := entities.Activity{
			Activity: `some fake activity`,
			Participants: 3,
		}
		byteAct, err := json.Marshal(act)
		if err != nil {
			t.Error(`err marshal activity:`, err)
		}
		fmt.Fprintf(w, "%s", string(byteAct))
	}))
	defer testHttpServ.Close()

	serv := New(testHttpServ.URL, actUseCase)
	ctx := context.Background()
	activity, err := serv.GetActivity(ctx, nil)
	if err != nil {
		t.Error(`err is not nil:`, err)
	}

	t.Log(activity)
}