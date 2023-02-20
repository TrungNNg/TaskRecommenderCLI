package cmd

import (
    "log"
    "encoding/gob"
    "bytes"
    "time"

	"github.com/spf13/cobra"
    "github.com/boltdb/bolt"
)

var doCmd = &cobra.Command{
  Use:   "do <task id>",
  Short: "mark task as complete",
  Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
      db, err := bolt.Open("task.db", 0600, nil)
      if err != nil {
          log.Fatal(err)
      }
      defer db.Close()
      err = db.Update(func(tx *bolt.Tx) error {
          bucket := tx.Bucket(bucket_name)
          if err != nil {
            log.Fatal(err)
          }
          for _, arg := range args {
              var task Task
              data := bucket.Get([]byte(arg))
              buf := bytes.NewBuffer(data)
              decoder := gob.NewDecoder(buf)
              decoder.Decode(&task)

              task.Done = true
              task.Done_time = time.Now()

              buf = new(bytes.Buffer)
              encoder := gob.NewEncoder(buf)
              encoder.Encode(task)

              bucket.Put([]byte(task.Id), buf.Bytes())
          }
          return nil
      })

      if err != nil {
          log.Fatal(err)
      }
  },
}

func init() {
    rootCmd.AddCommand(doCmd)
}

