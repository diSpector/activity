package commands

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/diSpector/activity.git/pkg/activity/grpc"
)

// activityCmd represents the activity command
var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "get activity",
	Long:  `Get random activity from api. Use flag "--stream" (-s) from stream output`,
	Run: func(cmd *cobra.Command, args []string) {
		creds, err := credentials.NewClientTLSFromFile(`/etc/keys/grpc/grpc_server.crt`, `mercator`)
		if err != nil {
			log.Println(`err create client with credentials:`, err)
			os.Exit(1)
		}

		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(creds),
		}

		conn, err := grpc.Dial("localhost:50053", opts...)
		if err != nil {
			log.Println(`err dial grpc:`, err)
			os.Exit(1)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		client := pb.NewActivityApiClient(conn)

		if !streamFlag {
			activity, err := client.GetActivity(ctx, &pb.Empty{})
			if err != nil {
				log.Println(`api returns err:`, err)
				return
			}

			log.Println(activity)
		} else {
			stream, err := client.GetActivityStream(ctx, nil)
			if err != nil {
				log.Println(`err get stream`, err)
				return
			}

			for {
				letter, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatalln("err read letter from stream:", err)
				}
				log.Printf(letter.Text)
			}
		}
	},
}
