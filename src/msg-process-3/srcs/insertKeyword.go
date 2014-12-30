package main

import (
//	"log"
	"net"
	"bytes"
	"net/http"
	"encoding/base64"
	"strings"
	"github.com/belogik/goes"
	"github.com/go-martini/martini"
)

/* Here we should write a native script for better performance and race condition */
func insertKeyword(ipStr string, keyword string, hash string, conn *goes.Connection) (keywordData, error) {
	docs := []goes.Document{
		{
			Id:          hash,
			Index:       "mrdrive",
			Type:        "keywords",
			BulkCommand: goes.BULK_COMMAND_INDEX,
			Fields: map[string]interface{}{
				"keyword": keyword,
				"ips": []interface{}{},
			},
		},
	}

	if result, err := searchKeyword(keyword, conn); err == nil {
		if len(result.Ips) > 0 {
			var updated = false

			for _, ip := range result.Ips {
				if ip.(map[string]interface{})["ip"] == ipStr {
					ip.(map[string]interface{})["used"] = ip.(map[string]interface{})["used"].(float64) + 1
					updated = true

				}
			}
			if !updated {
				result.Ips = append(result.Ips, map[string]interface{} {
					"ip": ipStr,
					"used": 1,
				})
			}
			// update
			docs[0].Fields.(map[string]interface{})["ips"] = result.Ips
		} else {
			ips := append([]interface{}{}, map[string]interface{}{
				"ip": ipStr,
				"used": 1,
			})
			docs[0].Fields.(map[string]interface{})["ips"] = ips
		}

		if _, err := conn.BulkSend(docs); err == nil {
			//log.Printf("RESPONSE : %#v", response)
			return keywordData{Keyword : keyword,
				Ips: docs[0].Fields.(map[string]interface{})["ips"].([]interface{}),
			}, nil
		} else {
			return keywordData{}, err
		}
	}
	return keywordData{}, nil
}

func fillDatabase(params martini.Params, c *goes.Connection, r *http.Request) ([]keywordData, error) {
	var rets []keywordData

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	r.ParseForm();

	for _, message := range r.Form {
		for _, keyword := range strings.Split(message[0], " ") {
			var b bytes.Buffer

			w := base64.NewEncoder(base64.URLEncoding, &b)
			if _, err := w.Write(bytes.NewBufferString(keyword).Bytes()); err != nil {
				return rets, err
			}
			w.Close()

			hash := string(b.Bytes())

			if ret, err := insertKeyword(ip, keyword, hash, c); err == nil {
				rets = append(rets, ret)
			} else {
				return rets, err
			}

		}
	}
	return rets, nil
}
