package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func GreetUser(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	if name == "" {
		name = "Forge User"
	}
	dir, _ := os.Getwd()
	fmt.Printf("Hello, I am here:  %s!\n", dir)
}
