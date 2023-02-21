package cmd

import (
    "strings"
    "log"
    "strconv"
    "encoding/gob"
    "bytes"
    "time"
    "fmt"

	"github.com/spf13/cobra"
	"github.com/boltdb/bolt"
)

type Task struct {
    Id string // Task's Id in BoltDB
    Description string // Task's detail
    StartTime time.Time // Time when the task was added
}

var addCmd = &cobra.Command{
  Use:   "add <task descriptsion>",
  Short: "add task with given description. Ex: task add Do homework.",
  Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
	//fmt.Println("add task")	
	db, err := bolt.Open("task.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

      // ask user for task details then add to DB
      err = db.Update(func(tx *bolt.Tx) error {
          bucket := tx.Bucket(task_bucket)

          // get next id in task_bucket
          id, _ := bucket.NextSequence()
          ID := []byte(strconv.Itoa(int(id)))

          task := Task{Id:string(ID), Description:strings.Join(args, " "),StartTime:time.Now()}

          buf := new(bytes.Buffer)
          encoder := gob.NewEncoder(buf)
          encoder.Encode(task)

          err = bucket.Put( ID , buf.Bytes() )
          if err != nil {
              return err
          }
          fmt.Println("Successfuly added", task)
          return nil
      })

      if err != nil {
          log.Fatal(err)
      }
  },
}

func init() {
	rootCmd.AddCommand(addCmd)
}