package main

import (
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!</h1>"))
	log.Print("someone is here /")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/hi", indexHandler)
	mux.HandleFunc("/bye", byeHandler)
	http.ListenAndServe(":"+port, mux)
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Bye World!</h1>"))
	log.Print("someone is here /bye")
}
