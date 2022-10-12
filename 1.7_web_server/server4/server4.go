package main

import (
	"fmt"
	lissajous "livro_golang/1.4lissajous/lissajous_export"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handleLissajous)
	http.HandleFunc("/count", counter)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handleLissajous(w http.ResponseWriter, r *http.Request) {
	cycles, _ := strconv.ParseFloat(r.URL.Query().Get("cycles"), 64)
	lissajous.Lissajous(w, cycles)

}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
	mu.Unlock()
}
