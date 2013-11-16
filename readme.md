# timeback api

API Access to timeback.

`go build api.go`


## Data

All responses are `Content-Type: Application/JSON`. XML encoding and other variations are planned.


## Endpoints

URI                     | Description
---                     | ---
[`/auth/`](#auth)       | Manage OAuth stuff
[`/account/`](#account) | Get and manage details for the authenticated account
[`/tasks/`](#tasks)     | The tasks woo


### Auth

timeback uses OAuth for authentication.





### Account

`GET /account/`: retrieve account details for the authenticated user

`PUT /account/`: update account details for the authenticated user


### Tasks

`GET /tasks/`: return the first page of tasks for the authenticated user

`POST /tasks/`: save a new task

`GET /tasks/task_id`: full details about a specified task

`PUT /tasks/task_id`: update a task

`DELETE /tasks/task_id`: remove a task


