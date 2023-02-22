package cmd

import (
    "log"
    "encoding/gob"
    "bytes"
    "time"
    "fmt"


	"github.com/spf13/cobra"
    "github.com/boltdb/bolt"
)

var doCmd = &cobra.Command{
  Use:   "do <task id>",
  Short: "recommend a task",
  Run: func(cmd *cobra.Command, args []string) {
      db, err := bolt.Open("/Users/trungng/desktop/TaskRecommenderCLI/task.db", 0600, nil)
      if err != nil {
          log.Fatal(err)
      }
      defer db.Close()

      err = db.View(func(tx *bolt.Tx) error {
          bucket := tx.Bucket(task_bucket)
          if err != nil {
            log.Fatal(err)
          }

          c := bucket.Cursor()
          task_id := ""
          task_des := ""
          var task_time int64

          // use when user want random task
          m := map[string]string{}

          for k, v := c.First(); k != nil; k, v = c.Next() {
              var task Task
              buf := bytes.NewBuffer(v)
              decoder := gob.NewDecoder(buf)
              decoder.Decode(&task)

              // save task's id and des in a map
              m[task.Id] = task.Description

              // recommend task with highest time since startTime
              if task_time < time.Since(task.StartTime).Milliseconds() {
                task_id = task.Id
                task_des = task.Description
                task_time = time.Since(task.StartTime).Milliseconds()
              }
          }

          if randomFlag {
            for k, v := range m {
                fmt.Printf("---\nRecommend task: %s. %s\n---\n", k, v)
                // map in Go is random, so just pick the first key
                return nil
            }
          }

          fmt.Printf("---\nRecommend task: %s. %s\n---\n",task_id, task_des)
          return nil
      })
      if err != nil {
          log.Fatal(err)
      }
  },
}

var randomFlag bool

func init() {
    rootCmd.AddCommand(doCmd)
    doCmd.Flags().BoolVarP(&randomFlag, "random", "r", false, "pick a random task")
}

