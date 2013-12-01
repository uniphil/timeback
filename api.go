/*

TODO

  * [ ] refactor db n stuff out of this file
  * [ ] redirect handler (return body in correct format

*/

package main

import (
	"os"
	"fmt"
    "time"
	"net/http"
	"database/sql"
)

const (
	AUTH_URI = "/auth/"
	ACCOUNT_URI = "/account/"
	TASKS_URI = "/tasks/"
)

func main() {
	fmt.Println("hello this is api")
	db := ConnectDB(os.Getenv("POSTGRES_URI"))
	AttachEndpoint(AUTH_URI, GetAuthEndpoint(db))
	AttachEndpoint(ACCOUNT_URI, GetAcctEndpoint(db))
	AttachEndpoint(TASKS_URI, GetTasksEndpoint(db))
	Listen(os.Getenv("HOST"))
}

func GetAuthEndpoint(db *sql.DB) *Endpoint {
	index := func(r *http.Request) (interface{}, int) {
		return "auth hello", 200
	}
	post := func(r *http.Request) (interface{}, int) {
		return "auth post", 200
	}
	get := func(r *http.Request, id string) (interface{}, int) {
		return "auth get", 200
	}
	put := func(r *http.Request, id string) (interface{}, int) {
		return "auth put", 200
	}
	delete := func(r *http.Request, id string) (interface{}, int) {
		return "auth delete", 200
	}
	return &Endpoint{Index:index, Post:post, Get:get, Put:put, Delete:delete}
}

func GetAcctEndpoint(db *sql.DB) *Endpoint {
	index := func(r *http.Request) (interface{}, int) {
		return "acct hello", 200
	}
	post := func(r *http.Request) (interface{}, int) {
		return "acct post", 200
	}
	get := func(r *http.Request, id string) (interface{}, int) {
		return "acct get", 200
	}
	put := func(r *http.Request, id string) (interface{}, int) {
		return "acct put", 200
	}
	delete := func(r *http.Request, id string) (interface{}, int) {
		return "acct delete", 200
	}
	return &Endpoint{Index:index, Post:post, Get:get, Put:put, Delete:delete}
}

func GetTasksEndpoint(db *sql.DB) *Endpoint {
	index := func(r *http.Request) (interface{}, int) {
		tasks, _ := LoadAccountTasks(db, "0")
		return tasks, 200
	}
	post := func(r *http.Request) (interface{}, int) {
		desc := r.Form["description"]
		durStr := r.Form["duration"]
		if len(desc) == 0 || len(durStr) == 0 {
			return "400 bad request", 400
		}
		dur, parseErr := time.ParseDuration(durStr[0])
		if parseErr != nil {
			return "bad duration", 400
		}
		t := &Task{Description: desc[0], Duration: dur}
		t.Save(db)
		return "Saved task", 200
	}
	get := func(r *http.Request, id string) (interface{}, int) {
		task, err := LoadTask(db, id)
		if err == sql.ErrNoRows {
			return "No task found for you there, sorry :(", 404
		}
		return task, 200
	}
	put := func(r *http.Request, id string) (interface{}, int) {
		return "tasks put (not implemented yet)", 200
	}
	delete := func(r *http.Request, id string) (interface{}, int) {
		DeleteTask(db, id)
		return "Deleted task.", 200
	}
	return &Endpoint{Index:index, Post:post, Get:get, Put:put, Delete:delete}
}
