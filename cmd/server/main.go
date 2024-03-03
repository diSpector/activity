package main

import (
	"log"
	"net"

	"github.com/diSpector/activity.git/internal/server"
	"github.com/diSpector/activity.git/internal/server/repository"
	"github.com/diSpector/activity.git/internal/server/usecase"
	"google.golang.org/grpc"

	pb "github.com/diSpector/activity.git/pkg/activity/grpc"
)

const API_URL = `https://www.boredapi.com/api/activity`
const PORT = `50053`

func main() {
	log.Println(`server run`)

	actRepo, err := repository.NewActivitySqlLiteRepo(`../../storage/sqlite.db`)
	if err != nil {
		log.Fatalln(`err create repo:`, err)
	}

	actUseCase := usecase.NewActivityUseCaseImpl(actRepo)

	serv := server.New(API_URL, actUseCase)

	lis, err := net.Listen(`tcp`, `localhost:`+PORT)
	if err != nil {
		log.Fatalln(`err listen:`, err)
	}

	s := grpc.NewServer()
	pb.RegisterActivityApiServer(s, serv)

	log.Println(`grpc server is listening`)

	if err := s.Serve(lis); err != nil {
		log.Fatalln(`failed serve:`, err)
	}
}
