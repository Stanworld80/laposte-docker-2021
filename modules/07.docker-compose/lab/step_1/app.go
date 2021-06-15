
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Message received: %s\n", r.URL.Path[1:])
}

func main() {
    lastcheck := time.Now()
    http.HandleFunc("/healthz", func (w http.ResponseWriter, r *http.Request) {
        duration := time.Now().Sub(lastcheck)
        lastcheck = time.Now()
        if duration.Seconds() > 10 {
            w.WriteHeader(500)
            w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
        } else {
            w.WriteHeader(200)
            w.Write([]byte("ok"))
        }
    })
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}