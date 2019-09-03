package cmd

import (
	"testing"

	"github.ibm.com/dash/dash_utils/dashtest"
)

func TestRoot(t *testing.T) {
	RootCmd.Execute()
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
