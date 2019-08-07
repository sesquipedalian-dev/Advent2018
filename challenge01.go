package main

import (
	"bufio"
	"io"
	"net/http"
	"strconv"
)

func challenge01(w http.ResponseWriter, req *http.Request) {
	// init
	sum := 0
	scanner := bufio.NewScanner(req.Body)

	// read input
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			panic("number formatting exception")
		}
		sum += num
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, strconv.Itoa(sum))
}
