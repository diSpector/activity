package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var streamFlag bool
var description string
var personsCnt int

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Root command",
	Long:  `This is root command. Everything starts here`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(activityCmd)

	activityCmd.Flags().BoolVarP(&streamFlag, `stream`, `s`, false, `use stream`)

	activityCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&description, `description`, `d`, ``, `add your description for activity`)
	addCmd.Flags().IntVarP(&personsCnt, `persons`, `p`, 0, `add persons count for your activity`)

	activityCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&description, `description`, `d`, ``, `search activity by word or phrase`)

	activityCmd.AddCommand(listCmd)
}
