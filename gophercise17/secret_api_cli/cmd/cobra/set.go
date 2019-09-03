package cobra

import (
	"fmt"
	"gophercises/gophercise17/secret_api_cli"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret store.",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret_api_cli.File(encodingKey, secretsPath())
		key := args[0]
		value := args[1]
		err := v.Set(key, value)
		if err != nil {
			return
		}
		fmt.Println("Value set successfully.")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
