package main

/*
 * Server over https and http.
 *
 * Generate the keys `go run /usr/local/go/src/pkg/crypto/tls/generate_cert.go -ca=true -host="localhost"`
 */

import (
    . "fmt"
    . "net/http"
)

const MESSAGE = "hello world"
const ADDRESS = ":1024"
const SECURE_ADDRESS = ":1025"

func main() {
    // Run https server in main thread, blockking the main function from terminating.
    // Run the http server in a goroutine
    HandleFunc("/hello", func(w ResponseWriter, r *Request) {
        w.Header().Set("Content-Type", "text/plain")
        Fprintf(w, MESSAGE)
    })
    go func() {
        Printf("Server starting on port %v \n", ADDRESS)
        ListenAndServe(ADDRESS, nil)
    }()
    Printf("Server starting on port %v \n", SECURE_ADDRESS)
    ListenAndServeTLS(SECURE_ADDRESS, "secret/cert.pem", "secret/key.pem", nil)
}
