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
func (t *Task) Update(db *sql.DB) error {
	// NOT IMPLEMENTED
	return nil
}
func DeleteTask(db *sql.DB, task_id string) error {
	result, err := db.Exec("DELETE FROM tasks WHERE id=$1", task_id)
	if err != nil {
		log.Fatal(err)
	}
	_, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	// SHOULD RAISE ERROR IF NON-EXISTENT ROWS?
	return nil
}
func LoadTask(db *sql.DB, task_id string) (*Task, error) {
	var id, dur int
	var desc string
	row := db.QueryRow("SELECT id, description, duration FROM tasks WHERE id=$1", task_id)
	err := row.Scan(&id, &desc, &dur)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
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
