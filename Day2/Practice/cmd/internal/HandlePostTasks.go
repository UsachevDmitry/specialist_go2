package internal

import (
	"net/http"
	"encoding/json"
)
var StatusCode int

func HandlePostTasks() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var task Task
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&task)
		if err != nil {
			StatusCode = WriteHeaderAndSaveStatus(http.StatusBadRequest, w)
			message := Message{Message: "provided json file is invalid."}
			json.NewEncoder(w).Encode(message)
			return
		}

		var NewTaskID int = len(DB) + 1
		task.ID = NewTaskID
		DB = append(DB, task)
		GlobalSugar.Infoln("task with id =", NewTaskID, "Created.")
		StatusCode = WriteHeaderAndSaveStatus(http.StatusCreated, w)
		json.NewEncoder(w).Encode(task)
	}
	return http.HandlerFunc(fn)
}

