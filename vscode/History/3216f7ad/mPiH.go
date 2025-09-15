package db

import "github.com/boltdb/bolt"
var taskBucket = []byte("tasks")
var db *bolt.DB

type Tasks struct{
	Key int
	Value string
}
	func 