package taskstore

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

type TaskStore struct {
	sync.Mutex

	tasks  map[int]Task
	nextId int
}

func New() *TaskStore {
	ts := &TaskStore{}
	ts.tasks = make(map[int]Task)
	ts.nextId = 0
	return ts
}

func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int {
	ts.Lock()
	defer ts.Unlock()

	task := Task{
		Id:   ts.nextId,
		Text: text,
		Due:  due,
	}
	task.Tags = make([]string, len(tags))
	copy(task.Tags, tags)

	ts.tasks[ts.nextId] = task
	ts.nextId++
	return task.Id

}

func (ts *TaskStore) GetTask(id int) (Task, error) {
	ts.Lock()
	defer ts.Unlock()

	t, ok := ts.tasks[id]
	if ok {
		return t, nil
	} else {
		return Task{}, fmt.Errorf("task with id=%d not found", id)
	}
}

func (ts *TaskStore) DeleteTask(id int) error {
	return
}

// DeleteAllTasks удаляет из хранилища все задачи.
func (ts *TaskStore) DeleteAllTasks() error {
	return
}

func (ts *TaskStore) GetAllTasks() []Task {
	return
}

func (ts *TaskStore) GetTaskByTag(tag string) []Task {
	return
}

func (ts *TaskStore) GetTasksByDueDate(year int, month time.Month, day int) []Task {
	return
}

// corner taskHandler

// func (ts *TaskServer)
