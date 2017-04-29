package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show TODO Tasks",
	Long: `Display TODO tasks from Google-Datasotre`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list tasks")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
