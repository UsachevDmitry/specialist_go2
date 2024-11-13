package main

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"	
	"net/http"
)

var Number1 int = 100
var Number2 int = 200

// Наша модель
type Calc struct {
	Number1 int `json:"Number1"`
	Number2 int `json:"Number2"`
	Sum     int `json:"sum"`
	Sub     int `json:"sub"`
	Mul     int `json:"mul"`
	Div     float64 `json:"div"`
}

var Calc1 = Calc{
	Number1: 0,
	Number2: 0,
	Sum:     0,
	Sub:     0,
	Mul:     0,
	Div: 	 0,
}

func main() {

	var port string = "8080"
	router := mux.NewRouter()
	router.HandleFunc("/info", GetInfo).Methods("GET")
	router.HandleFunc("/first", GetFirst).Methods("GET")
	router.HandleFunc("/second", GetSecond).Methods("GET")
	router.HandleFunc("/add", GetAdd).Methods("GET")
	router.HandleFunc("/sub", GetSub).Methods("GET")
	router.HandleFunc("/mul", GetMul).Methods("GET")
	router.HandleFunc("/div", GetDiv).Methods("GET")
	log.Println("Router configured successfully! Let's go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/info - Информация об API\n")
	fmt.Fprintf(w, "/first - Случайное число 1\n")
	fmt.Fprintf(w, "/second - Случайное число 2\n")
	fmt.Fprintf(w, "/add - Cумма двух случайных чисел\n")
	fmt.Fprintf(w, "/sub - Разность двух случайных чисел\n")
	fmt.Fprintf(w, "/mul - Произведение двух случайных чисел\n")
	fmt.Fprintf(w, "/div - Деление двух случайных чисел\n")
}

func GetFirst(w http.ResponseWriter, r *http.Request) {
	Calc1.Number1 = rand.Intn(100)
	json.NewEncoder(w).Encode(Calc1.Number1)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func GetSecond(w http.ResponseWriter, r *http.Request) {
	Calc1.Number2 = rand.Intn(100)
	json.NewEncoder(w).Encode(Calc1.Number2)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func GetAdd(w http.ResponseWriter, r *http.Request) {
	Calc1.Sum = Calc1.Number1 + Calc1.Number2
	json.NewEncoder(w).Encode(Calc1.Sum)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func GetSub(w http.ResponseWriter, r *http.Request) {
	Calc1.Sub = Calc1.Number1 - Calc1.Number2
	json.NewEncoder(w).Encode(Calc1)
	json.NewEncoder(w).Encode(Calc1.Sub)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func GetMul(w http.ResponseWriter, r *http.Request) {
	Calc1.Mul = Calc1.Number1 * Calc1.Number2
	json.NewEncoder(w).Encode(Calc1)
	json.NewEncoder(w).Encode(Calc1.Mul)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func GetDiv(w http.ResponseWriter, r *http.Request) {
	Calc1.Div = float64(Calc1.Number1) / float64(Calc1.Number2)
	json.NewEncoder(w).Encode(Calc1)
	json.NewEncoder(w).Encode(Calc1.Div)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
