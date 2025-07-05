package main

import (
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"

	"github.com/alwindoss/shepherd"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewMux()
	webFS, err := fs.Sub(shepherd.WebFS, "web")
	if err != nil {
		log.Fatal(err)
	}
	r.Handle("/*", http.FileServer(http.FS(webFS)))

	setupAPIRoutes(r)

	// Serve on port 8080
	addr := net.JoinHostPort("", "8080")
	fmt.Println("Listening on", addr)
	http.ListenAndServe(addr, r)
}

func setupAPIRoutes(mux *chi.Mux) {
	// v1 users route
	mux.Route("/shepherd/api/v1/users", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("This is the GET /users route"))
		})
	})
	// v1 groups route
	mux.Route("/shepherd/api/v1/groups", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("This is the GET /groups route"))
		})
	})
}
