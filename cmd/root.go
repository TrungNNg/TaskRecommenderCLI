package cmd

import (
    "fmt"
    "os"
    "log"

    "github.com/spf13/cobra"
    "github.com/boltdb/bolt"
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
  done        Mark a task on your TODO list as complete
  list        List all of your incomplete tasks
  do          Recommend a task
Use "task [command] --help" for more information about a command.`)
  },
}

var task_bucket = []byte("task")

func init() {
  // init DB
  db, err := bolt.Open("task.db", 0600, nil)
  if err != nil {
      log.Fatal(err)
  }
  defer db.Close()

  err = db.Update(func(tx *bolt.Tx) error {
      _, err = tx.CreateBucketIfNotExists([]byte(task_bucket))
      if err != nil {
          return fmt.Errorf("create bucket: %s", err)
      }
      return nil
  })
  
  if err != nil {
      log.Fatal(err)
  }
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}




