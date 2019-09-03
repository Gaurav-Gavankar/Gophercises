package secret_api_cli

import (
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.ibm.com/dash/dash_utils/dashtest"
)

func secretsPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, "secrets-test-file")
}
func TestSet(t *testing.T) {
	encodingKey := "random-key"
	v := File(encodingKey, secretsPath())
	key, value := "vault-key", "vault-val"
	err := v.Set(key, value)
	if err != nil {
		t.Error("Expecting values to set successfully, but got ", err)
	}
}

func TestSetNegative(t *testing.T) {
	encodingKey := "random-key"
	v := File(encodingKey, "/")
	key, value := "vault-key", "vault-val"
	err := v.Set(key, value)
	if err == nil {
		t.Error("Expecting error, but got ", err)
	}
}

func TestGet(t *testing.T) {
	encodingKey := "random-key"
	v := File(encodingKey, secretsPath())
	key := "vault-key"
	expectedVal := "vault-val"
	val, err := v.Get(key)
	if err != nil {
		t.Errorf("Expecting values to be \"%s\", but got %s", expectedVal, err)
	} else if val != expectedVal {
		t.Errorf("Expecting values to be \"%s\", but got %s", expectedVal, val)
	}
}

func TestGetNegative(t *testing.T) {
	encodingKey := "random-key"
	v := File(encodingKey, secretsPath())
	key := "fake-vault-key"
	_, err := v.Get(key)
	if err.Error() != "No value found for the Key." {
		t.Errorf("Expecting no value for given key. But got wrong value.")
	}
}

func TestLoad(t *testing.T) {
	encodingKey := "random-key"
	v := File(encodingKey, "/")
	v.load()
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
