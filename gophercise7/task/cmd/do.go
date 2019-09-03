package cmd

import (
	"fmt"
	"gophercises/gophercise7/task/db"
	"strconv"

	"github.com/spf13/cobra"
)

var TasksList = db.AllTasks

var RemoveTask = db.DeleteTask

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument.")
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := TasksList()
		if err != nil {
			fmt.Println("Someting went wrong : ", err)
			return
		}
		for _, id := range ids {
			if id < 0 || id > len(tasks) {
				fmt.Println("Invalid task number ", id)
				continue
			}
			task := tasks[id-1]
			err1 := RemoveTask(int(task.Key))
			if err1 != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error : %s \n", id, err1)
			} else {

				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
