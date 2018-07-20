package main

import (
    "fmt"
    "log"
    "net/http"
    "json_parser/packages/parsers"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s", json.Parse(r))
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
