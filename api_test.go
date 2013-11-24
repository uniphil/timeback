package main

import (
	"testing"
	"bytes"
	"net/http"
)

func TestBlah(t *testing.T) {
	l, _ := LoadAccount("hello")
	if l.Email != "hello" {
		t.Error("email didn't match")
	}
}

func TestAuthIndexViews(t *testing.T) {
	auth_endpoint := GetAuthEndpoint()
	req, _ := http.NewRequest("GET", "/auth/", bytes.NewBufferString("."))
	data := auth_endpoint.Index(req)
	if data != "auth hello" {
		t.Error("noooooooo this is such a bad test")
	}
}


