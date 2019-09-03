package main

import (
	"gophercises/gophercise17/secret_api_cli/cmd/cobra"
)

//in memory version
/*func main() {
	v := secret_api_cli.Memory("fake-key")
	err := v.Set("demo-key", "demo value")
	if err != nil {
		panic(err)
	}
	plain, err := v.Get("demo-key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain text:", plain)
}
*/

//file backed version
/*func main() {
	v := secret_api_cli.File("my-fake-key", "C:/go-work/src/gophercises/gophercise17/secret_api_cli/cmd/secrets")
	err := v.Set("demo_key1", "some crazy value")
	if err != nil {
		panic(err)
	}
	err = v.Set("demo_key2", "456 some crazy value")
	if err != nil {
		panic(err)
	}
	err = v.Set("demo_key3", "789 some crazy value")
	if err != nil {
		panic(err)
	}
	plain, err := v.Get("demo_key1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)
	plain, err = v.Get("demo_key2")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)
	plain, err = v.Get("demo_key3")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)
}
*/

//CLI version
func main() {
	cobra.RootCmd.Execute()

}
