package main

import (
	"net/http"
	"strings"
)

func getId(r *http.Request) string {
	p := strings.Split(r.URL.Path, "/")
	if len(p) == 1 {
		return p[0]
	} else if len(p) > 1 {
		return p[1]
	} else {
		return ""
	}
}
