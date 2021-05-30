package main

import (
	"fmt"
	"log"
	"net/http"
)

//type Handler interface {
//  ServeHTTP(ResponseWriter, *Request)
//}
//
//func ListenAndServe(addr string, handler Handler) error {
//
//}

func PlayerServer (w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "20")
}

func main()  {
  handler := http.HandlerFunc(PlayerServer)
  if err := http.ListenAndServe(":5000", handler); err != nil {
    log.Fatalf("could not listen on port 5000 %v", err)
  }
}
