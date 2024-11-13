package internal

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func HandleDeleteTaskByID() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["taskid"])
		if err != nil {
			GlobalSugar.Infoln("Error occurs while parsing id field:", err)
			StatusCode = WriteHeaderAndSaveStatus(http.StatusBadRequest, w)
			message := Message{Message: "don't use ID parametr as uncasted to int."}
			json.NewEncoder(w).Encode(message)
			return
		}
	
		_, ok := FindTaskByID(id)
		
		if !ok {
			GlobalSugar.Infoln("task with id =", id, "not found.")
			StatusCode = WriteHeaderAndSaveStatus(http.StatusNotFound, w)
			message := Message{Message: "task with that ID does not exist in database."}
			json.NewEncoder(w).Encode(message)
			return
		} 
		
		// Delete task
		DeleteTaskById(id)
		GlobalSugar.Infoln("task with id =", id, "Deleted.")
		StatusCode = WriteHeaderAndSaveStatus(http.StatusOK, w)
		message := Message{Message: "task has successfully deleted from database."}
		json.NewEncoder(w).Encode(message)
	}
	return http.HandlerFunc(fn)
}