package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "task",
  Short: "CLI TODO task manager",
  Long: "TODO task manager using Cobra and BoltDB.",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println(`
Usage: task [command]

Available Commands:
  add         Add a new task to your TODO list
  do          Mark a task on your TODO list as complete
  list        List all of your incomplete tasks

Use "task [command] --help" for more information about a command.`)
  },
}

func init() {
  
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}




