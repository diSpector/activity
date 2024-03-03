package commands

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/diSpector/activity.git/internal/cli/validators"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/diSpector/activity.git/pkg/activity/grpc"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   `add --description="Some description" --persons=5`,
	Short: "Add activity manually",
	Long:  `Add activity manually with description and persons amount`,
	Run: func(cmd *cobra.Command, args []string) {
		isValid, err := validators.ValidateAdd(description, personsCnt, args)
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

		_, err = client.AddActivity(ctx, &pb.Activity{
			Activity:     description,
			Participants: int32(personsCnt),
		})

		if err != nil {
			log.Println(`api returns err:`, err)
			return
		}

		log.Printf("activity `%s` successfully added\n", description)
	},
}
