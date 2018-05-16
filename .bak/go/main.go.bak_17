package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	folder string
	prefix string
	port   int
)

func Init() {
	flag.StringVar(&folder, "folder", "./html", "The html file folder path name. Default value is \"./html\".")
	flag.StringVar(&folder, "f", "./html", "The html file folder path name. Default value is \"./html\".")

	flag.StringVar(&prefix, "prefix", "/html/", "The html url prefix. Default value is \"/html/\".")
	flag.StringVar(&prefix, "P", "/html/", "The html url prefix. Default value is \"/html/\".")

	flag.IntVar(&port, "port", 8080, "The http port number. Default value is \"8080\".")
	flag.IntVar(&port, "p", 8080, "The http port number. Default value is \"8080\".")

	flag.Parse()

	log.Printf("folder : %s, prefix : %s, port : %d", folder, prefix, port)
}

func main() {
	Init()

	http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(folder))))
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
