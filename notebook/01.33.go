package main

/*
 * Server over https.
 *
 * Generate the keys `go run /usr/local/go/src/pkg/crypto/tls/generate_cert.go -ca=true -host="localhost"`
 */

import (
    . "fmt"
    "net/http"
)

const MESSAGE = "hello world"
const SECURE_ADDRESS = ":1025"

func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain")
        Fprintf(w, MESSAGE)
    })
    Printf("Server starting on port %v \n", SECURE_ADDRESS)
    if e := http.ListenAndServeTLS(SECURE_ADDRESS, "secret/cert.pem", "secret/key.pem", nil); e != nil {
        Println(e)
    }
}
