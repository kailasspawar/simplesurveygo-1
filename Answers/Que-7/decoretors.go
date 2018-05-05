package main

import (
    "fmt"
    "log"
    "net/http"
)

func simpleHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "ok")
}

func timetaken(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        startTime := time.Now()
        h.ServeHTTP(w, r)
        duration := time.Now().Sub(startTime)
       fmt.Fprintln(w,duration)
    })
}

func main() {
    http.HandleFunc("/decorated", timetaken(simpleHandler))
    http.HandleFunc("/notdecorated", simpleHandler)

    http.ListenAndServe("127.0.0.1:8080", nil)
}
