package main

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"
)

const AUTH_URI = "/auth/"
const ACCOUNT_URI = "/account/"
const TASKS_URI = "/tasks/"


func main() {
	fmt.Println("hello")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "hello") })
	http.HandleFunc(AUTH_URI, apiHandler(AUTH_URI, auth_index, auth_post, auth_get, auth_put, auth_delete))
	//http.HandleFunc(ACCOUNT_URI, apiHandler(ACCOUNT_URI, acct_index, acct_post, acct_get, acct_put, acct_delete))
	//http.HandleFunc(TASKS_URI, apiHandler(TASKS_URI, tasks_index, tasks_post, tasks_get, tasks_put, tasks_delete))
	http.ListenAndServe(":8000", nil)
	fmt.Println("goodbye")
}


func apiHandler(root string, index, post func(*http.Request) interface{}, get, put, delete func(*http.Request, string) interface{}) func(http.ResponseWriter, *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var data interface{}
		var error int
		var endpoint_root bool
		error = 0
		endpoint_root = (len(r.URL.Path) <= len(root))

		if endpoint_root {
			if r.Method == "GET" {
				data = index(r)
			} else if r.Method == "POST" {
				data = post(r)
			} else {
				error = 400
				data = "\"bad request\""
			}
		} else {
			id := r.URL.Path[len(root):]
			if r.Method == "GET" {
				data = get(r, id)
			} else if r.Method == "PUT" {
				data = put(r, id)
			} else if r.Method == "DELETE" {
				data = delete(r, id)
			} else {
				error = 400
				data = "\"bad request (not root though)\""
			}
		}
		w.Header().Set("Content-Type", "application/json")
		if error == 0 {
			jdata, _ := json.Marshal(data)
			w.Write(jdata)
		} else {
			http.Error(w, data.(string), error)
		}
	}
	return handler
}


func auth_index(r *http.Request) interface{} {
	return "auth hello"
}
func auth_post(r *http.Request) interface{} {
	return "auth post"
}
func auth_get(r *http.Request, id string) interface{} {
	return "auth get"
}
func auth_put(r *http.Request, id string) interface{} {
	return "auth put"
}
func auth_delete(r *http.Request, id string) interface{} {
	return "auth delete"
}


type Account struct {
	Email string
}

func (a *Account) save() error {
	return nil
}

func (a *Account) update() error {
	return nil
}

func (a *Account) remove() error {
	return nil
}

func LoadAccount(email string) (*Account, error) {
	return &Account{Email: email}, nil
}

type Task struct {
	ID string
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
	return &Task{ID: task_id,
				 Project: "test",
				 Description: "hello world",
				 Duration: dur,
				 Date: time.Date(2013, time.November, 15, 10, 0, 0, 0, time.UTC),
				 Creation: time.Date(2013, time.November, 15, 11, 0, 0, 0, time.UTC),
				 }, nil
}

func LoadAccountTasks(account_id string) ([]*Task, error) {
	dur, _ := time.ParseDuration("10s")
	return []*Task{&Task{ID: account_id,
			 Project: "test",
			 Description: "hello world",
			 Duration: dur,
			 Date: time.Date(2013, time.November, 15, 10, 0, 0, 0, time.UTC),
			 Creation: time.Date(2013, time.November, 15, 11, 0, 0, 0, time.UTC),
			 }}, nil
}
