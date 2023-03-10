package cmd

import (
    "log"

	"github.com/spf13/cobra"
    "github.com/boltdb/bolt"
)

var doneCmd = &cobra.Command{
  Use:   "done <task id>",
  Short: "mark task done and remove using task id",
  Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
      db, err := bolt.Open("/Users/trungng/desktop/TaskRecommenderCLI/task.db", 0600, nil)
      if err != nil {
          log.Fatal(err)
      }
      defer db.Close()
      err = db.Update(func(tx *bolt.Tx) error {
          bucket := tx.Bucket(task_bucket)
          if err != nil {
            log.Fatal(err)
          }
          for _, arg := range args {
              bucket.Delete([]byte(arg))
          }
          return nil
      })

      if err != nil {
          log.Fatal(err)
      }
  },
}

func init() {
    rootCmd.AddCommand(doneCmd)
}
