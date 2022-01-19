package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var(
	port = flag.String("port",":8080","port the server would run on")
)
func main(){
  flag.Parse()
  l := log.New(os.Stdout,"[PONG] ",log.LstdFlags)
  r := http.NewServeMux()

  r.HandleFunc("/ping",func(rw http.ResponseWriter, r *http.Request) {
	  l.Printf("request recieved from %s using %s",r.RemoteAddr,r.Method)
	  rw.Write([]byte("pong"))
  })
  l.Printf("booted up pong server on %s",*port)
  err := http.ListenAndServe(*port,r)
  if err != nil {
	  l.Fatal(err)
  }

}
