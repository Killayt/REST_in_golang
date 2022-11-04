package main

import (
	"encoding/json"
	"log"
	"mime"
	"net/http"
	"time"

	"stdlib/internal/taskstore"
)

type taskServer struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *taskServer {
	store := taskstore.New()
	return &taskServer{store: store}
}

func (ts *taskServer) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling task create at %s : \n", r.URL.Path)

	type RequestTask struct {
		Text string    `json:"text"`
		Tags []string  `json:"tags"`
		Due  time.Time `json:"due"`
	}

	type ResponsedID struct {
		ID int `json:"id"`
	}

	contentType := r.Header.Get("Content-type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	if mediaType != "application/json" {
		http.Error(w, "except application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var rt RequestTask
	if err := decoder.Decode(&rt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func startServer() error {
	mux := http.NewServeMux()
	server := taskstore.NewTaskServer()

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
