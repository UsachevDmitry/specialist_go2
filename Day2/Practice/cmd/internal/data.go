package internal

import (
	//"time"
	"slices"
)

var DB []Task

type Task struct {
	ID int `json:"id"`
	Text string `json:"text"`
	Tags string `json:"tags"`
	Due string `json:"due"`
}

func init() {
	task := Task{
		ID: 1,
		Text: "Task 1",
		Tags: "Tag1",
		Due: "2006-01-02 15:00:00",
	}
	DB = append(DB, task)
}

func FindTaskByID(id int) (Task, bool) {
	var task Task
	var found bool
	for _, t := range DB {
		if t.ID == id{
			task = t
			found = true
			break
		}
	}
	return task, found
}

func DeleteTaskById(id int) bool {
	for idx, t := range DB {
		if t.ID == id {
			DB = slices.Delete(DB, idx, idx+1)
			return true
		}
	}
	return false
}

func DeleteTasks() {
			DB = nil
}
