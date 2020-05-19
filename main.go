package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
)

func logRequestHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		fmt.Println(r.RemoteAddr, r.Method, r.URL, r.Proto)
	})
}

func main() {
	mux := &http.ServeMux{}

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "web/index.html")
	})
	mux.HandleFunc("/replicache-client-view", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, `{
  "lastMutationID": 0,
  "clientView": {
    "/todo/1": {
      "title": "Take out the trash",
      "description": "Don't forget to pull it to the curb.",
      "done": false
    },
    "/todo/2": {
      "title": "Pick up milk",
      "description": "Whole4",
      "done": true
    }
  }
}`)
	})
	mux.HandleFunc("/wasm_exec.js",
		func(w http.ResponseWriter, req *http.Request) {
			http.ServeFile(w, req, path.Join("web", req.URL.Path))
		})
	mux.HandleFunc("/repjs",
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/wasm")
			http.ServeFile(w, req, "web/repjs")
		})
	mux.HandleFunc("/repjs.br",
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/wasm")
			w.Header().Set("Content-Encoding", "br")
			http.ServeFile(w, req, "web/repjs.br")
		})

	port := 8090
	if p := os.Getenv("PORT"); p != "" {
		if i, err := strconv.Atoi(p); err == nil {
			port = i
		}
	}
	fmt.Printf("Listening on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), logRequestHandler(mux))
}
