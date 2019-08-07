package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

func main() {
	flag.Parse()

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/1", http.HandlerFunc(challenge01))

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
