package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	port  = flag.String("port", ":8080", "port the server would run on")
	pongs = 0
)

func main() {
	flag.Parse()
	l := log.New(os.Stdout, "[PONG] ", log.LstdFlags)
	r := http.NewServeMux()

	r.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		pongs += 1
		l.Printf("request recieved from %s using %s", r.RemoteAddr, r.Method)
		rw.Write([]byte("pong"))
	})
	r.HandleFunc("/pongs", func(rw http.ResponseWriter, r *http.Request) {
		pongs += 1
		rw.Write([]byte(strconv.Itoa(pongs)))
	})
	r.HandleFunc("/version", func(rw http.ResponseWriter, r *http.Request) {
		version := os.Getenv("PONG_VERSION")
		if version == "" {
			version = "v1"
		}
		rw.Write([]byte(version))
	})
	l.Printf("booted up pong server on %s", *port)
	err := http.ListenAndServe(*port, r)
	if err != nil {
		l.Fatal(err)
	}

}
