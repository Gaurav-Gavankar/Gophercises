package cmd

import (
	"fmt"
	"gophercises/gophercise7/task/db"

	"github.com/spf13/cobra"
)

var TaskList = db.AllTasks

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := TaskList()
		if err != nil {
			fmt.Println("Something went wrong : ", err.Error())
			// os.Exit(1)
			return
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete!")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d: %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
