package main

import (
	"fmt"
	"gophercises/gophercise7/task/cmd"
	"gophercises/gophercise7/task/db"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	DbPath := filepath.Join(home, "command.db")
	must(db.Init(DbPath))
	must(cmd.RootCmd.Execute())
	db.CLoseDB()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
