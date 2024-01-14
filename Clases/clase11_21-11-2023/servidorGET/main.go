package main

import (
	// "fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	queryParameters := r.URL.Query()
	paramValue := queryParameters.Get("parameter")
	w.Write([]byte("<h1>" + paramValue + "</h1>"))
	// fmt.Fprintf(w, "Parameter is, %s", paramValue)
}

func main() {
	http.HandleFunc("/", handler) // se debe solicitar el recurso escribiendo localhost:8080/?parameter=valor
	http.ListenAndServe(":3006", nil)
}
