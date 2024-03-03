package commands

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/diSpector/activity.git/internal/cli/validators"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/diSpector/activity.git/pkg/activity/grpc"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all saved activities",
	Long:  `list all saved activities from DB`,
	Run: func(cmd *cobra.Command, args []string) {
		isValid, err := validators.ValidateList(args)
		if !isValid {
			log.Println(err)
			os.Exit(1)
		}

		conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Println(`err dial grpc:`, err)
			os.Exit(1)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		client := pb.NewActivityApiClient(conn)

		stream, err := client.ListActivities(ctx, &pb.Empty{})
		if err != nil {
			log.Println(`err get stream:`, err)
			return
		}

		var i int
		for {
			activity, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("err read activity from stream:", err)
			}
			log.Printf("found activity `%s` for %d persons\n", activity.Activity, activity.Participants)
			i++
		}

		if i == 0 {
			log.Printf("activities for phrase `%s` NOT found\n", description)
		}
	},
}
