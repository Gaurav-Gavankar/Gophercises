package cmd

import (
	"fmt"
	"gophercises/gophercise7/task/db"
	"strings"

	"github.com/spf13/cobra"
)

//createTask := db.CreateTask

var GenerateTask = db.CreateTask

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		//_, err := createTask(task)
		_, err := GenerateTask(task)
		if err != nil {
			fmt.Println("Someting went wrong: ", err.Error())
			return
		}
		fmt.Printf("Added \"%s\" to your list.", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
