package cmd

import (
	"strings"
    "log"
    "strconv"
    "encoding/gob"
    "bytes"
    "time"

	"github.com/spf13/cobra"
	"github.com/boltdb/bolt"
)

type Task struct {
    Id string
    Name string
    Done bool
    Done_time time.Time
}

var addCmd = &cobra.Command{
  Use:   "add <task name>",
  Short: "add task",
  Run: func(cmd *cobra.Command, args []string) {
	//fmt.Println("add task")	
	db, err := bolt.Open("task.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


      // store some data
      err = db.Update(func(tx *bolt.Tx) error {
          bucket := tx.Bucket(bucket_name)

          id, _ := bucket.NextSequence()
          ID := []byte(strconv.Itoa(int(id)))

          task_name := strings.Join(args, "")
          task := Task{Id:string(ID), Done_time:time.Now(), Name:task_name, Done:false}

          buf := new(bytes.Buffer)
          encoder := gob.NewEncoder(buf)
          encoder.Encode(task)

          err = bucket.Put( ID , buf.Bytes() )
          if err != nil {
              return err
          }
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