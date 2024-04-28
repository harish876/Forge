package main

import (
	"fmt"
	"os"

	"github.com/harish876/forge/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "hello",
		Short: "A simple CLI tool to greet the user",
		Run:   commands.GreetUser,
	}

	var createStepCmd = &cobra.Command{
		Use:   "create_step",
		Short: "Create a new ETL Step",
		Long:  `Create a new ETL Step specifying the type and the name of the step.`,
		Run:   commands.CreateStep,
	}

	createStepCmd.Flags().StringP("type", "t", "", "Type of the step (e.g., extractor, transformer, loader, reporter, etc.)")
	createStepCmd.Flags().StringP("name", "n", "", "Name of the new step")

	createStepCmd.MarkFlagRequired("type")
	createStepCmd.MarkFlagRequired("name")

	rootCmd.Flags().StringP("name", "n", "", "Name of the person to greet")
	rootCmd.AddCommand(createStepCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
