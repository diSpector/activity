package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var streamFlag bool

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
}
