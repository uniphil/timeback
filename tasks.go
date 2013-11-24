package main

import(
	//"fmt"
	"time"
	"database/sql"
	_ "github.com/lib/pq"
)

type Task struct {
	ID string
	Project string
	Description string
	Duration time.Duration
	Date time.Time
	Creation time.Time
}
func (t *Task) Save() error {
	db, err := sql.Open("postgres", "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")
	//age := 23
	//rows, err := db.Query("SELECT name FROM users WHERE age=?", age)
	return nil
}
func (t *Task) Update() error {
	return nil
}
func (t *Task) Remove() error {
	return nil
}
func LoadTask(task_id string) (*Task, error) {
	dur, _ := time.ParseDuration("10s")
	return &Task{ID: task_id,
				 Project: "test",
				 Description: "hello world",
				 Duration: dur,
				 Date: time.Date(2013, time.November, 15, 10, 0, 0, 0, time.UTC),
				 Creation: time.Date(2013, time.November, 15, 11, 0, 0, 0, time.UTC),
				 }, nil
}

func LoadAccountTasks(account_id string) ([]*Task, error) {
	// SELECT * FROM tasks WHERE account=$1 LIMIT 50 ORDER_BT date ASC
	task, _ := LoadTask("lalala")
	return []*Task{task}, nil
}
