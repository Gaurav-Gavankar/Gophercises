package cmd

import (
	"errors"
	"path/filepath"
	"testing"

	"gophercises/gophercise7/task/db"

	homedir "github.com/mitchellh/go-homedir"
)

func TestAdd(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "command.db")
	db.Init(dbPath)
	a := []string{"Go", "to", "home"}
	b := []string{"Go", "to", "room"}

	addCmd.Run(addCmd, a)
	addCmd.Run(addCmd, b)
	db.CLoseDB()
}

func TestAddNegative(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "command.db")
	db.Init(dbPath)
	a := []string{"Go", "to", "home"}

	tmp := GenerateTask
	defer func() {
		GenerateTask = tmp
	}()

	GenerateTask = func(task string) (int, error) {
		return 100, errors.New("DUMMY ERROR")
	}

	addCmd.Run(addCmd, a)
	db.CLoseDB()
}
