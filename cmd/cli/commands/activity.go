package commands

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/diSpector/activity.git/pkg/activity/grpc"
)

// activityCmd represents the activity command
var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "get activity",
	Long: `Get random activity from api`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Println(`err dial grpc:`, err)
			os.Exit(1)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		client := pb.NewActivityApiClient(conn)

		activity, err := client.GetActivity(ctx, &pb.Empty{})
		if err != nil {
			log.Println(`api returns err:`, err)
			return
		}

		log.Println(activity)
	},
}

func init() {
	rootCmd.AddCommand(activityCmd)
}
