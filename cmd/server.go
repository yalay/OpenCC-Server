package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/yalay/OpenCC-Server/controllers"
)

var tplPath string
var listenPort int

func init() {
	flag.StringVar(&tplPath, "t", "views", "frontend template files path")
	flag.IntVar(&listenPort, "p", 8015, "listen port")
	flag.Parse()
}

func main() {
	router := httprouter.New()
	router.GET("/", controllers.S2tGetHandler)
	router.POST("/s2t", controllers.S2tPostHandler)

	router.ServeFiles("static/*filepath", http.Dir(tplPath))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(listenPort), router))
}
