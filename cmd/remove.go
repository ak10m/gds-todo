package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove TASK_ID",
	Short: "Remove TODO Task",
	Long: `Remove TODO task from Google-Datasotre`,
	Run: func(cmd *cobra.Command, args []string) {
		taskId := args[0]
		fmt.Printf("remove task: %v\n", taskId)
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
}
