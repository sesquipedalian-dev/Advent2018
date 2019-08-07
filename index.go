package main

import (
	"bufio"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

const indexTemplateStr = `
<html>
<head>
<title>Advent of Code 2018</title>
</head>
<body>
<h> Advent of Code 2018, in <a href="https://golang.org">Go!</a>! </h>
</body>
</html>
`

var indexTempl = template.Must(template.New("index").Parse(indexTemplateStr))

func index(w http.ResponseWriter, req *http.Request) {
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
