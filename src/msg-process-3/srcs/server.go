package main

import (
	"log"
	"net"
	"net/http"
	"./routes"
	"os"
	"strings"
	"regexp"
	"github.com/belogik/goes"
	"github.com/go-martini/martini"
	"github.com/codegangsta/martini-contrib/render"
)

var (
	ES_HOST = "172.17.0.7"
	ES_PORT = "9200"
	m *martini.Martini
)

var rxExt = regexp.MustCompile(`(\.(?:xml|json))\/?$`)


type keywordData struct {
	Keyword string
	Ips interface{}
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

	m.Use(render.Renderer())
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

func fillDatabase(r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	r.ParseForm();

	log.Println("Ip => " + ip)

	for key, value := range r.Form {
		log.Println("Key:", key, "Value:", value)
	}
}

func processDataHandler() (int, string) {
	// if r.Method == "POST" {
	// 	fillDatabase(r)
	// } else {
	// 	log.Println(r.Method + " Method received, skipped")
	// }
	return 200, "OK"
}

func main() {
	m.RunOnAddr(":8080")
}
