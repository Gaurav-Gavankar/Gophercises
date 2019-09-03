package cmd

import (
	"errors"
	"path/filepath"
	"testing"

	"gophercises/gophercise7/task/db"

	homedir "github.com/mitchellh/go-homedir"
)

func TestList(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "command.db")
	db.Init(dbPath)
	a := []string{}
	listCmd.Run(listCmd, a)
	db.CLoseDB()
}

func TestListNegative(t *testing.T) {
	a := []string{}
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "command.db")
	db.Init(dbPath)

	listCmd.Run(listCmd, a)

	tmp := TaskList
	defer func() {
		TaskList = tmp
	}()
	TaskList = func() ([]db.Task, error) {
		return nil, errors.New("dummy error")
	}

	listCmd.Run(listCmd, a)
	db.CLoseDB()
}

func TestListEmpty(t *testing.T) {
	a := []string{}
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "test5.db")
	db.Init(dbPath)

	listCmd.Run(listCmd, a)
	db.CLoseDB()
}
