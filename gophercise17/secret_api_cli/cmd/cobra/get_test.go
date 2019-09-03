package cobra

import (
	"testing"
)

func TestGet(t *testing.T) {
	arg := []string{"test_key"}
	getCmd.Run(getCmd, arg)
}

func TestGetNegative(t *testing.T) {
	arg := []string{"false_test_key"}
	getCmd.Run(getCmd, arg)
}
