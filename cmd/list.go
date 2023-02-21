package cmd

import (
	"fmt"
    "log"
    "encoding/gob"
    "bytes"
    "time"

	"github.com/spf13/cobra"
    "github.com/boltdb/bolt"
)

var listCmd = &cobra.Command{
  Use: "list",
  Short: "list all current tasks",
  Run: func(cmd *cobra.Command, args []string) {
      db, err := bolt.Open("task.db", 0600, nil)
      if err != nil {
          log.Fatal(err)
      }
      defer db.Close()

      // print all Task in task bucket
      err = db.View(func(tx *bolt.Tx) error {
          bucket := tx.Bucket(task_bucket)
          if bucket == nil {
              return fmt.Errorf("Bucket %q not found!", task_bucket)
          }

          c := bucket.Cursor()
          for k, v := c.First(); k != nil; k, v = c.Next() {
              var task Task
              buf := bytes.NewBuffer(v)
              decoder := gob.NewDecoder(buf)
              decoder.Decode(&task)

              fmt.Printf("%s. %s\n",k,task.Description)
              fmt.Printf("time since added %s\n\n", time.Since(task.StartTime).String())              
          }
          return nil
      })

  if err != nil {
      log.Fatal(err)
  }},
}

//var completedFlag bool
//var incompleteFlag bool

func init() {
	rootCmd.AddCommand(listCmd)
	//listCmd.Flags().BoolVarP(&completedFlag, "completed", "c", false, "display only completed tasks")
    //listCmd.Flags().BoolVarP(&incompleteFlag, "incomplete", "i", false, "display only incomplete tasks")
}