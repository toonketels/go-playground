package main

import (
    . "fmt"
    "net/http"
)

const MESSAGE = "hello world"
const ADDRESS = ":1024"

func main() {
    http.HandleFunc("/hello", Hello)
    Printf("Server starting on port %v \n", ADDRESS)
    if e := http.ListenAndServe(ADDRESS, nil); e != nil {
        Println(e)
    }
}

func Hello(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    // File print formatted
    // ResponseWriter is a file handle
    Fprintf(w, MESSAGE)
}
