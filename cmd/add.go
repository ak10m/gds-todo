package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/ak10m/gds-todo/todo"
)

var addCmd = &cobra.Command{
	Use:   "add TASK_DESCRIPTION",
	Short: "Add TODO Task",
	Long: `Store new TODO task to Google-Datasotre`,
	Run: func(cmd *cobra.Command, args []string) {
		desc := args[0]
		task := todo.CreateTask(desc)
		log.Printf(">>> add task\n%v\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
