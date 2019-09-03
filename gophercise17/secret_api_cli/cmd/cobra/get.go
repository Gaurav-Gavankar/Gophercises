package cobra

import (
	"fmt"
	"gophercises/gophercise17/secret_api_cli"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a secret from your secret store.",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret_api_cli.File(encodingKey, secretsPath())
		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			fmt.Println("No value set for Key- ", key)
			return
		}
		fmt.Printf("%s = %s\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
