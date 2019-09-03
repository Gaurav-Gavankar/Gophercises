package cobra

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "Secret is an API key manager",
}

var encodingKey string

func init() {
	RootCmd.PersistentFlags().StringVarP(&encodingKey, "key", "k", "", "The key to use when encoding and decoding the secrets.")
}

func secretsPath() string {
	home, _ := homedir.Dir()
	path := filepath.Join(home, "secrets")
	return path
}
