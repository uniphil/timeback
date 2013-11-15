package main

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"
)

func main() {
	fmt.Println("hello")
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/auth/", apiHandler(auth))
	http.HandleFunc("/account/", apiHandler(account))
	http.HandleFunc("/tasks/", apiHandler(tasks))
	http.ListenAndServe(":8000", nil)
	fmt.Println("goodbye")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello hello")
}

func apiHandler(endpoint func(*http.Request) interface{}) func(http.ResponseWriter, *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		data := endpoint(r)
		jdata, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jdata)
	}
	return handler
}

func auth(r *http.Request) interface{} {
	t, _ := LoadTask("asdf")
	return t
}

func account(r *http.Request) interface{} {
	t, _ := LoadTask("asdf")
	return t
}

func tasks(r *http.Request) interface{} {
	t, _ := LoadTask("asdf")
	return t
}


type Task struct {
	Project string
	Description string
	Duration time.Duration
	Date time.Time
	Creation time.Time
}

func (t *Task) save() error {
	return nil
}

func (t *Task) update() error {
	return nil
}

func (t *Task) remove() error {
	return nil
}

func LoadTask(task_id string) (*Task, error) {
	dur, _ := time.ParseDuration("10s")
	return &Task{Project: "test",
				 Description: "hello world",
				 Duration: dur,
				 Date: time.Date(2013, time.November, 15, 10, 0, 0, 0, time.UTC),
				 Creation: time.Date(2013, time.November, 15, 11, 0, 0, 0, time.UTC),
				 }, nil
}


