package internal

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func HandleGetTaskByID() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["taskid"])
		if err != nil {
			StatusCode = WriteHeaderAndSaveStatus(http.StatusBadRequest, w)
			message := Message{Message: "don't use ID parametr as uncasted to int."}
			json.NewEncoder(w).Encode(message)
			return
		}
		task, ok := FindTaskByID(id)
		GlobalSugar.Infoln("Get task with id:", id)
		if !ok {
			StatusCode = WriteHeaderAndSaveStatus(http.StatusNotFound, w)
			message := Message{Message: "task with that ID does not exist in database."}
			json.NewEncoder(w).Encode(message)
		} else {
			StatusCode = WriteHeaderAndSaveStatus(http.StatusOK, w)
			GlobalSugar.Infoln("task with id =", id, "Found.")
			json.NewEncoder(w).Encode(task)
		}
	}
	return http.HandlerFunc(fn)
}