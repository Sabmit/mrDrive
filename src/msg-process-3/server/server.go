package main

import (
	"log"
	"net"
	"net/http"
)

func fillDatabase(r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	r.ParseForm();

	log.Println("Ip => " + ip)

	for key, value := range r.Form {
		log.Println("Key:", key, "Value:", value)
	}
}

func processDataHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		fillDatabase(r)
	} else {
		log.Println(r.Method + " Method received, skipped")
	}
}

func adminHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", processDataHandler)
	http.HandleFunc("/admin", adminHandler)

	http.ListenAndServe(":8080", nil)
}
