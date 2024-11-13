package main

import (
	"github.com/xhermitx/itty-bitty/internal/controllers"
	"github.com/xhermitx/itty-bitty/internal/db"
	"github.com/xhermitx/itty-bitty/internal/url"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		urlDB = db.New()
		svc   = url.NewService(urlDB) // Pass a DB
		c     = controllers.NewController(svc)
	)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		shortURL := request.URL.Path

		if shortURL == "/" {
			controllers.ShowIndex(writer, request)
		} else {
			c.Redirect(writer, request)
		}
	})
	http.HandleFunc("/shorten", c.Shortener)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("$PORT must be set")
	}
	log.Fatal(http.ListenAndServe(port, nil))
}
