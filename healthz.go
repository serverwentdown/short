package main

import (
	"fmt"
	"net/http"
	"runtime"
)

var version string

func healthz(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "ok\n")
	fmt.Fprintf(w, "ver: %v\n", version)
	fmt.Fprintf(w, "go ver: %v\n", runtime.Version())
}
