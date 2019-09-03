package cmd

import (
	"errors"
	"gophercises/gophercise7/task/db"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func TestDoRemove(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "command.db")
	db.Init(dbPath)
	input1 := []string{"1"}
	input2 := []string{"-1", "2", "100"}
	input3 := []string{"-1", "hi", "100"}

	tmp2 := RemoveTask
	defer func() {
		RemoveTask = tmp2
	}()

	RemoveTask = func(key int) error {
		return errors.New("dummy ERROR")
	}

	doCmd.Run(doCmd, input1)

	doCmd.Run(doCmd, input2)

	doCmd.Run(doCmd, input3)
	db.CLoseDB()
}

func TestDoNegative(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "command.db")
	db.Init(dbPath)
	input1 := []string{"1"}
	input2 := []string{"-1", "2", "100"}
	input3 := []string{"-1", "hi", "100"}
	doCmd.Run(doCmd, input1)

	tmp := TasksList
	tmp1 := TasksList
	tmp2 := RemoveTask
	defer func() {
		TaskList = tmp
		TaskList = tmp1
		RemoveTask = tmp2
	}()

	RemoveTask = func(key int) error {
		return errors.New("dummy ERROR")
	}

	doCmd.Run(doCmd, input1)
	/*
		defer func() {
			TasksList = tmp
		}()*/
	TasksList = func() ([]db.Task, error) {
		return nil, errors.New("dummy Error")
	}
	/*
		tmp1 := TasksList
		defer func() {
			TasksList = tmp1
		}()*/
	TasksList = func() ([]db.Task, error) {
		return nil, errors.New("dummy Error")
	}

	doCmd.Run(doCmd, input2)

	doCmd.Run(doCmd, input3)
	db.CLoseDB()
}

func TestDo(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "command.db")
	db.Init(dbPath)
	input1 := []string{"1"}
	input2 := []string{"1", "hi", "2", "100"}
	doCmd.Run(doCmd, input1)
	doCmd.Run(doCmd, input2)

	db.CLoseDB()
}
