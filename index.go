package main

import (
	"html/template"
	"net/http"
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
	indexTempl.Execute(w, req.FormValue("s"))
}
