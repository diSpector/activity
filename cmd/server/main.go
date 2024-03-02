package main

import (
	"crypto/tls"
	"log"
	"net"

	"github.com/diSpector/activity.git/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/diSpector/activity.git/pkg/activity/grpc"
)

const API_URL = `https://www.boredapi.com/api/activity`
const PORT = `50053`

func main() {
	log.Println(`server run`)

	cert, err := tls.LoadX509KeyPair(`/etc/keys/grpc/grpc_server.crt`, `/etc/keys/grpc/grpc_server.key`)
	if err != nil {
		log.Fatalln(`err load x509 key pair:`, err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	
	serv := server.New(API_URL)

	s := grpc.NewServer(opts...)
	pb.RegisterActivityApiServer(s, serv)

	lis, err := net.Listen(`tcp`, `localhost:`+PORT)
	if err != nil {
		log.Fatalln(`err listen:`, err)
	}

	log.Println(`grpc server is listening`)

	if err := s.Serve(lis); err != nil {
		log.Fatalln(`failed serve:`, err)
	}
}
