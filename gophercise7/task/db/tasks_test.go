package db

import (
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.ibm.com/dash/dash_utils/dashtest"
)

func TestInitNegative(t *testing.T) {
	err := Init("/")
	if err == nil {
		t.Error("Expected error. Got nil.")
	}
}

func TestInit(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "command.db")
	err := Init(dbPath)
	if err != nil {
		t.Error("Expected database connection. Got an Error : ", err)
	}
}

func TestCreateTask(t *testing.T) {
	_, err := CreateTask("Test CreateTask function.")
	if err != nil {
		t.Error("Expecting ID but got an error : ", err)
	}
}

func TestAllTasks(t *testing.T) {
	_, err := AllTasks()
	if err != nil {
		t.Error("Expecting pending tasks. Got an Error : ", err)
	}
}

func TestDeleteTask(t *testing.T) {
	taskNunmber := 1
	err := DeleteTask(taskNunmber)
	if err != nil {
		t.Error("Expecting task ", taskNunmber, " to get deleted.\nBut got an error: ", err)
	}
}

func TestAllTasksNegative(t *testing.T) {
	TaskBucket = []byte{}
	_, err := AllTasks()
	if err == nil {
		t.Error("Expecting an Error. Got ", err)
	}
}

func TestAllTasksNegativeAdded(t *testing.T) {
	db.Close()
	_, err := AllTasks()
	if err == nil {
		t.Error("Expecting an Error. Got ", err)
	}
}

func TestCreateTaskNegative(t *testing.T) {
	db.Close()
	_, err := CreateTask("Test CreateTask function.")
	if err == nil {
		t.Error("Expecting an Error. Got ", err)
	}
}

func TestBtoi(t *testing.T) {
	input := []byte{0, 0, 0, 0, 0, 0, 0, 8}
	output := btoi(input)
	expected := 8
	if output != 8 {
		t.Errorf("Expecting %v, Got %v.", expected, output)
	}
}

func TestItob(t *testing.T) {
	expected := []byte{0, 0, 0, 0, 0, 0, 0, 8}
	input := 8
	output := itob(input)
	if output == nil {
		t.Errorf("Expecting %v, Got %v", expected, output)
	}
}

func TestCloseDB(t *testing.T) {
	CLoseDB()
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
