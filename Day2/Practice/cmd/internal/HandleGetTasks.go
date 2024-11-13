package internal

import (
	"net/http"
	"encoding/json"
)


func HandleGetTasks() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(DB)
		GlobalSugar.Infoln("All tasks Found.")
		StatusCode = WriteHeaderAndSaveStatus(http.StatusOK, w)
	}
	return http.HandlerFunc(fn)
}