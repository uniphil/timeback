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
	http.HandleFunc(AUTH_URI, apiHandler(AUTH_URI, getAuthEndpoint()))
	//http.HandleFunc(ACCOUNT_URI, apiHandler(ACCOUNT_URI, acctIndex, acctPost, acctGet, acctPut, acctDelete))
	//http.HandleFunc(TASKS_URI, apiHandler(TASKS_URI, tasksIndex, tasksPost, tasksGet, tasksPut, tasksDelete))
	http.ListenAndServe(":8000", nil)
	fmt.Println("goodbye")
}


func apiHandler(root string, endpoint *Endpoint) func(http.ResponseWriter, *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var data interface{}
		var error int
		var endpoint_root bool
		error = 0
		endpoint_root = (len(r.URL.Path) <= len(root))

		if endpoint_root {
			if r.Method == "GET" {
				data = endpoint.Index(r)
			} else if r.Method == "POST" {
				data = endpoint.Post(r)
			} else {
				error = 400
				data = "\"bad request\""
			}
		} else {
			id := r.URL.Path[len(root):]
			if r.Method == "GET" {
				data = endpoint.Get(r, id)
			} else if r.Method == "PUT" {
				data = endpoint.Put(r, id)
			} else if r.Method == "DELETE" {
				data = endpoint.Delete(r, id)
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

type Endpoint struct {
	Index func(*http.Request) interface{}
	Post func(*http.Request) interface{}
	Get func(*http.Request, string) interface{}
	Put func(*http.Request, string) interface{}
	Delete func(*http.Request, string) interface{}
}


func getAuthEndpoint() *Endpoint {
	index := func(r *http.Request) interface{} {
		return "auth hello"
	}
	post := func(r *http.Request) interface{} {
		return "auth post"
	}
	get := func(r *http.Request, id string) interface{} {
		return "auth get"
	}
	put := func(r *http.Request, id string) interface{} {
		return "auth put"
	}
	delete := func(r *http.Request, id string) interface{} {
		return "auth delete"
	}
	return &Endpoint{Index:index, Post:post, Get:get, Put:put, Delete:delete}
}


func acctIndex(r *http.Request) interface{} {
	return "account index"
}
func acctPost(r *http.Request) interface{} {
	return "account post"
}
func acctGet(r *http.Request, id string) interface{} {
	return "acct get"
}
func acctPut(r *http.Request, id string) interface{} {
	return "acct put"
}
func acctDelete(r *http.Request, id string) interface{} {
	return "acct delete"
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
