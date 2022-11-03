package main

import (
	"log"
	"net/http"

	taskstore "stdlib/iternal/taskstore/"
)

type taskServer struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *taskServer {
	store := taskstore.New()
	return &taskServer{store: store}
}

func startServer() error {
	mux := http.NewServeMux()
	server := NewTaskServer()

	mux.HandleFunc("/task/", server.GetTask)
	mux.HandleFunc("/tag/", server.GetTaskByTag)
	mux.Handler("/due/", server.GetTasksByDueDate)

	log.Fatal(http.ListenAndServe(":8080", mux))

	return http.ErrAbortHandler
}

func main() {
	if err := startServer(); err != nil {
		log.Fatalf(err.Error())
	}
}
