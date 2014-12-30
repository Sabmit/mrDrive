package main

import (
	"log"
	"net/http"
	"./routes"
	"github.com/belogik/goes"
	"github.com/go-martini/martini"
)

func processDataHandler(params martini.Params, conn *goes.Connection, r *http.Request, enc routes.Encoder) (int, string) {
	if result, err := fillDatabase(params, conn, r); err == nil {
		if ret, err := enc.Encode(result); err == nil {
			log.Printf("\n\nresult RET = %#v", ret)
			log.Printf("\n\nresult RESULT = %#v", result)
			return http.StatusOK, ret
		}
	}
	return http.StatusInternalServerError, ""
}
