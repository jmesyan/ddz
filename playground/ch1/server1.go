package ch1

import (
	"log"
	"net/http"
	"strconv"
)

// Serve starts a lissajous server on localhost:8000
func Serve() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	defaultCycle := 5
	val := r.FormValue("cycles")
	if cycles, err := strconv.Atoi(val); err == nil {
		log.Println("lissajous with cycles=", cycles)
		Lissajous(w, cycles)
	} else {
		Lissajous(w, defaultCycle)
	}
}
