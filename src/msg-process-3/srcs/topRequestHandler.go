package main

import (
	"log"
	"net/http"
	"./routes"
	"github.com/belogik/goes"
	"github.com/go-martini/martini"
)

func topIpsHandler(params martini.Params, conn *goes.Connection, enc routes.Encoder) (int, string) {
	if result, err := searchTopIps(conn); err == nil {
		if ret, err := enc.Encode(result); err == nil {
			log.Println(result)
			log.Println(ret)
			return http.StatusOK, ret
		}
	}
	return http.StatusInternalServerError, ""
}

func topKeywordsHandler(params martini.Params, conn *goes.Connection, enc routes.Encoder) (int, string) {
	if result, err := searchTopKeyword(conn); err == nil {
		if ret, err := enc.Encode(result); err == nil {
			log.Println(result)
			log.Println(ret)
			return http.StatusOK, ret
		}
	}
	return http.StatusInternalServerError, ""
}

func keywordHandler(params martini.Params, conn *goes.Connection, enc routes.Encoder) (int, string) {
	_, _ = searchKeyword(params["keyword"], conn)
	return 200, ""
}
