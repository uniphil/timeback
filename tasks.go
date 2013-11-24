package main

import (
	"log"
	//"fmt"
	"time"
	"strconv"
	"database/sql"
)

type Task struct {
	/*CREATE TABLE tasks (
	id serial PRIMARY KEY,
	description text NOT NULL,
	duration integer NOT NULL,
	user_id integer NOT NULL);*/
	ID string
	//Project string
	Description string
	Duration time.Duration
	//Date time.Time
	//Creation time.Time
}
func (t *Task) Save(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO tasks (description, duration, user_id) VALUES ($1, $2, 1)",
							t.Description, DurationToDB(t.Duration))
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
func (t *Task) Update() error {
	return nil
}
func (t *Task) Remove() error {
	return nil
}
func LoadTask(db *sql.DB, task_id string) (*Task, error) {
	var id, dur int
	var desc string
	row := db.QueryRow("SELECT id, description, duration FROM tasks WHERE id=$1", task_id)
	err := row.Scan(&id, &desc, &dur)
	if err != nil {
		log.Fatal(err)
	}
	task := &Task{ID: strconv.Itoa(id), Description: desc, Duration: DurationFromDB(dur)}

	return task, nil
}

func LoadAccountTasks(db *sql.DB, account_id string) ([]*Task, error) {

	// SELECT * FROM tasks WHERE account=$1 LIMIT 50 ORDER_BT date ASC
	rows, err := db.Query("SELECT id, description, duration FROM tasks")
	if err != nil {
		log.Fatal(err)
	}

	var tasks []*Task
	for rows.Next() {
		var id, dur int
		var desc string

		if err := rows.Scan(&id, &desc, &dur); err != nil {
			log.Fatal(err)
		}

	    //dur, _ := time.ParseDuration("10s")
		task := &Task{ID: strconv.Itoa(id), Description: desc, Duration: DurationFromDB(dur)}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func DurationFromDB(s int) time.Duration {
	return time.Duration(int64(s) * int64(time.Second))
}

func DurationToDB(d time.Duration) int {
	return int(d / time.Second)
}
