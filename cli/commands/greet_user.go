package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GreetUser(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	if name == "" {
		name = "Forge User"
	}
	fmt.Printf("Hello, %s!\n", name)
}
