package main

import (
	"net/http"
	"encoding/json"
)

type EndpointRootFunc func(*http.Request) (interface{}, int)
type EndpointIDFunc func(*http.Request, string) (interface{}, int)
type Endpoint struct {
	Index EndpointRootFunc
	Post EndpointRootFunc
	Get EndpointIDFunc
	Put EndpointIDFunc
	Delete EndpointIDFunc
}

func APIHandler(root string, endpoint *Endpoint) func(http.ResponseWriter, *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var data interface{}
		var code int
		var endpoint_root bool
		code = 200
		endpoint_root = (len(r.URL.Path) <= len(root))

		if endpoint_root {
			if r.Method == "GET" {
				data, code = endpoint.Index(r)
			} else if r.Method == "POST" {
				r.ParseForm()
				data, code = endpoint.Post(r)
			} else {
				code = 400
				data = "\"bad request\""
			}
		} else {
			id := r.URL.Path[len(root):]
			if r.Method == "GET" {
				data, code = endpoint.Get(r, id)
			} else if r.Method == "PUT" {
				r.ParseForm()
				data, code = endpoint.Put(r, id)
			} else if r.Method == "DELETE" {
				data, code = endpoint.Delete(r, id)
			} else {
				code = 400
				data = "\"bad request (not root though)\""
			}
		}
		w.Header().Set("Content-Type", "application/json")
		if code == 200 {
			jdata, _ := json.Marshal(data)
			w.Write(jdata)
		} else {
			http.Error(w, data.(string), code)
		}
	}
	return handler
}

func AttachEndpoint(root string, endpoint *Endpoint) {
	http.HandleFunc(root, APIHandler(root, endpoint))
}

func Listen(host string) {
	http.ListenAndServe(host, nil)
}
