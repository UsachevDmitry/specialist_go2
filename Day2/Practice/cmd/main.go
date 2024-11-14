package main

import (
	"Go2/Day1/Practice/Day2/Practice/cmd/internal"
	"flag"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

const defaultAddr = "localhost:8080"

var addr = flag.String("a", defaultAddr, "Адрес HTTP-сервера")

func main() {
	internal.Sqlite() 
	flag.Parse()
	internal.Logger()

	addrEnv := os.Getenv("ADDRESS")
	if addrEnv != "" {
		*addr = addrEnv
	}
	router := mux.NewRouter()
	
	router.HandleFunc("/tasks/", internal.WithLoggingHandle(internal.HandlePostTasks())).Methods(http.MethodPost)
	router.HandleFunc("/tasks/", internal.WithLoggingHandle(internal.HandleGetTasks())).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{taskid}", internal.WithLoggingHandle(internal.HandleGetTaskByID())).Methods(http.MethodGet)
	router.HandleFunc("/tasks/", internal.WithLoggingHandle(internal.HandleDeleteTasks())).Methods(http.MethodDelete)
	router.HandleFunc("/tasks/{taskid}", internal.WithLoggingHandle(internal.HandleDeleteTaskByID())).Methods(http.MethodDelete)
	
	internal.GlobalSugar.Infow(
		"Starting server",
		"addr", *addr,
	)
	if err := http.ListenAndServe(*addr, router); err != nil {
		internal.GlobalSugar.Fatalw(err.Error(), "event", "start server")
	}
}
