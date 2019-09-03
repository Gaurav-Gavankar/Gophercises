package cobra

import (
	"testing"
)

func TestSet(t *testing.T) {
	args := []string{"test_key", "test_value"}
	setCmd.Run(setCmd, args)
}
