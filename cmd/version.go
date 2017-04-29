package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string = "0.0.1"

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of gds-todo",
	Long:  "Print version number of gds-todo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", version)
	},
}
