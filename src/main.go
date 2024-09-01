package main

import "fmt"
import "net/http"
import "log"


func main() {
    hello := func(w http.ResponseWriter, _ *http.Request) {
        fmt.Fprint(w, "Hello World!!")
    }

    goodbye := func(w http.ResponseWriter, _ *http.Request) {
        fmt.Fprint(w, "GoodBye World!!")
    }

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/goodbye", goodbye)

    log.Fatal(http.ListenAndServe(":8080", nil))
}