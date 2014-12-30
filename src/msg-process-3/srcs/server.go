package main

import (
	"net/http"
	"./routes"
	"os"
	"strings"
	"regexp"
	"github.com/belogik/goes"
	"github.com/go-martini/martini"
)

var (
	ES_HOST = "172.17.0.3"
	ES_PORT = "9200"
	m *martini.Martini
)

var rxExt = regexp.MustCompile(`(\.(?:xml|json))\/?$`)


type ipData struct {
	Ip string
	Used int
}

type keywordData struct {
	Keyword string
	Ips []interface{}
}

type keywordContainer struct {
	Keyword string
	Nb_used float64
}

type IpContainer struct {
	Keywords []keywordContainer
	Ip string
}

func init() {
	m = martini.New()

	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	m.Use(martini.Static("public"))
	m.Use(MapEncoder)

	// Setup routes
	r := martini.NewRouter()

	r.Post("/", processDataHandler)
	r.Get("/api/topKeywords", topKeywordsHandler)
	r.Get("/api/topIps", topIpsHandler)
	r.Get("/api/keywords/:keyword", keywordHandler)

	// Inject database
	m.Map(getConnection())

	// Add the router action
	m.Action(r.Handle)
}

func MapEncoder(c martini.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Get the format extension
	matches := rxExt.FindStringSubmatch(r.URL.Path)
	ft := ".json"
	if len(matches) > 1 {
		// Rewrite the URL without the format extension
		l := len(r.URL.Path) - len(matches[1])
		if strings.HasSuffix(r.URL.Path, "/") {
			l--
		}
		r.URL.Path = r.URL.Path[:l]
		ft = matches[1]
	}
	// Inject the requested encoder
	switch ft {
	case ".xml":
		c.MapTo(routes.XmlEncoder{}, (*routes.Encoder)(nil))
		w.Header().Set("Content-Type", "application/xml")
	default:
		c.MapTo(routes.JsonEncoder{}, (*routes.Encoder)(nil))
		w.Header().Set("Content-Type", "application/json")
	}
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

func main() {
	m.RunOnAddr(":8080")
}
