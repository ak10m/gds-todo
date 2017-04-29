package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add TASK_DESCRIPTION",
	Short: "Add TODO Task",
	Long: `Store new TODO task to Google-Datasotre`,
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		fmt.Printf("add task: %v\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
