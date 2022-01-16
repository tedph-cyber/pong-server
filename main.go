package main

import (
	"log"
	"net/http"
	"os"
)

var(
	args = os.Args
)
func main(){
  l := log.New(os.Stdout,"[PONG] ",log.LstdFlags)
  port := args[1]
  if port == ""{
	  port = ":8080"
  }
  r := http.NewServeMux()
  r.HandleFunc("/ping",func(rw http.ResponseWriter, r *http.Request) {
	  l.Printf("request recieved from %s using %s",r.RemoteAddr,r.Method)
	  rw.Write([]byte("pong"))
  })
  l.Printf("booted up pong server on %s",port)
  err := http.ListenAndServe(port,r)
  if err != nil {
	  l.Fatal(err)
  }

}
