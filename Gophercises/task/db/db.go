package db

import (
	"fmt"
	"github.com/boltdb/bolt"
	"path/filepath"
	"time"
)

var db *bolt.DB
var bucketName = "tasks"
var appPath = "/Users/sammarathkapse/go/src/task/db"
var bucketExists = 0

type Task struct {
	Key int
	Value string
}

func Init(dbname string) error {
	dbname += ".db"
	dbPath := filepath.Join(appPath, dbname)
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	if bucketExists == 0 {
		return createBucket()
	}
	return nil
}

func createBucket() error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		bucketExists = 1
		return nil
	})
}

func GetTaskList() ([]Task, error) {
	tasks, err := getAllTasks()
	return tasks, err
}

func AddTask(task string) error {
	_, err := addTask(task)
	return err
}

func DeleteTask(id int) error {
	err := deleteTask(id)
	return err
}
