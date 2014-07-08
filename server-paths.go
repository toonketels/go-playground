package main

import (
    "net/http"
    "fmt"
)

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

func (s String) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, s)
}

func (s *Struct) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, s.Greeting, " ", s.Punct," ", s.Who)
}

func main() {
    // your http.Handle calls here
    http.Handle("/string", String("Replying by string!"))
    http.Handle("/struct", &Struct{"Hello", "Mr", "Smith"})
    http.ListenAndServe("localhost:4000", nil)
}
