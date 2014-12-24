package main

import (
	"log"
	"fmt"
	"net"
//	"encoding/json"
	"net/http"
	"os"
	"github.com/belogik/goes"
)

var (
	ES_HOST = "172.17.0.21"
	ES_PORT = "9200"
)

type keywordContainer struct {
	keyword string
	nb_used float64
}

type IpContainer struct {
	keywords []keywordContainer
	ip string
}

func getConnection() (conn *goes.Connection) {
	h := os.Getenv("ELASTICSEARCH_HOST")
	if h == "" {
		h = ES_HOST
	}

	p := os.Getenv("ELASTICSEARCH_PORT")
	if p == "" {
		p = ES_PORT
	}

	conn = goes.NewConnection(h, p)
	return
}

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
	conn := getConnection()

	fmt.Println(searchTopKeyword(conn))
	fmt.Println(searchTopIps(conn))

	http.HandleFunc("/", processDataHandler)
	http.HandleFunc("/admin", adminHandler)

	http.ListenAndServe(":8080", nil)
}
