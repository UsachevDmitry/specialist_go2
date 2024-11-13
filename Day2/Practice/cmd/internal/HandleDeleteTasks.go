package internal

import (
	"net/http"
	"encoding/json"
)

func HandleDeleteTasks() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		DeleteTasks()
		GlobalSugar.Infoln("All tasks Deleted.")
		StatusCode = WriteHeaderAndSaveStatus(http.StatusOK, w)
		message := Message{Message: "All tasks has successfully deleted from database."}
		json.NewEncoder(w).Encode(message)
	}
	return http.HandlerFunc(fn)
}