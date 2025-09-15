package db

import (
	"time"

	"github.com/boltdb/bolt"
)
var taskBucket = []byte("tasks")
var db *bolt.DB

type Tasks struct{
	Key int
	Value string
}
	func Init(dbPath string)error{
		var err error
		db,err = bolt.Open(dbPath,0600, &bolt.Options{Timeout: 1*time.Second })
		if err!=nil {
			return err
			
		}
		return db.Update(func(tx *bolt.Tx) error {
			_,err:=tx.CreateBucketIfNotExists( taskBucket)
			return err
		})
	}
	func CreateTask(task string)(int,error){
		err:= db.Update(func(tx *bolt.Tx) error {
			
		})
		return 0,nil;
	}