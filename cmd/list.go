package cmd

import (
	"fmt"
    "log"
    "encoding/gob"
    "bytes"

	"github.com/spf13/cobra"
    "github.com/boltdb/bolt"
)

var listCmd = &cobra.Command{
  Use: "list",
  Short: "list all current tasks",
  Long: "list all current tasks",
  Run: func(cmd *cobra.Command, args []string) {
      db, err := bolt.Open("task.db", 0600, nil)
      if err != nil {
          log.Fatal(err)
      }
      defer db.Close()

      // retrieve the data
      err = db.View(func(tx *bolt.Tx) error {
          bucket := tx.Bucket(bucket_name)
          if bucket == nil {
              return fmt.Errorf("Bucket %q not found!", bucket_name)
          }

          c := bucket.Cursor()
          for k, v := c.First(); k != nil; k, v = c.Next() {

              var task Task
              buf := bytes.NewBuffer(v)
              decoder := gob.NewDecoder(buf)
              decoder.Decode(&task)

              if completedFlag && !task.Done {
                  continue
              }

              if incompleteFlag && task.Done {
                  continue
              }

              fmt.Printf("%s. %s - ",k, task.Name)
              if task.Done {
                fmt.Println("DONE")
              } else {
                fmt.Println()
              }
          }
          return nil
      })

  if err != nil {
      log.Fatal(err)
  }},
}

var completedFlag bool
var incompleteFlag bool

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&completedFlag, "completed", "c", false, "display only completed tasks")
    listCmd.Flags().BoolVarP(&incompleteFlag, "incomplete", "i", false, "display only incomplete tasks")
}