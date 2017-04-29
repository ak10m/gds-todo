package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done TASK_ID",
	Short: "Change status TODO Task",
	Long: `Change status TODO task on Google-Datasotre`,
	Run: func(cmd *cobra.Command, args []string) {
		taskId := args[0]
		fmt.Printf("done task: %v\n", taskId)
	},
}

func init() {
	RootCmd.AddCommand(doneCmd)
}
