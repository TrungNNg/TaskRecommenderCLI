package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
  Use:   "add",
  Short: "add task",
  Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("add task")	
  },
}

func init() {
	rootCmd.AddCommand(addCmd)
}