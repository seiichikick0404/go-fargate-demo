package main

import "fmt"
import "net/http"
import "log"


func main() {
    hello := func(w http.ResponseWriter, _ *http.Request) {
        fmt.Fprint(w, "Hello World!!")
    }

    http.HandleFunc("/hello", hello)

    log.Fatal(http.ListenAndServe(":8080", nil))
}