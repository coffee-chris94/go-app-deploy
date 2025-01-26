package main

import (
        "fmt"
        "log"
        "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Mondoo Engineer!")
}

func main() {
        http.HandleFunc("/", handler)
        
        // Check if there's an error returned by ListenAndServe
        if err := http.ListenAndServe(":8080", nil); err != nil {
                log.Fatal(err) // Log the error and stop the program if it fails
        }
}

