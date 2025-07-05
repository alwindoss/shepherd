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

// Handlers for different pages
// func homeHandler(renderer prism.Renderer) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		renderer.Render(w, "index.page.html", "base", map[string]interface{}{
// 			"Title": "Home Page",
// 		})
// 	}
// }

// func aboutHandler(renderer prism.Renderer) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		renderer.Render(w, "about.page.html", "base", map[string]interface{}{
// 			"Title": "About Us",
// 		})
// 	}
// }

// func contactHandler(renderer prism.Renderer) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		renderer.Render(w, "contact.page.html", "base", map[string]interface{}{
// 			"Title": "Contact Us",
// 		})
// 	}
// }

func main() {
	r := chi.NewMux()
	// vueFS, _ := fs.Sub(shepherd.WebFS, "web/")
	webFS, err := fs.Sub(shepherd.WebFS, "web")
	if err != nil {
		log.Fatal(err)
	}
	r.Handle("/*", http.FileServer(http.FS(webFS)))
	// r.Handle("/public/*", http.StripPrefix("/public/", pubFS))
	// r.HandleFunc("/", homeHandler(renderer))
	// r.HandleFunc("/about", aboutHandler(renderer))
	// r.HandleFunc("/contact", contactHandler(renderer))

	// Serve on port 8080
	addr := net.JoinHostPort("", "8080")
	fmt.Println("Listening on", addr)
	http.ListenAndServe(addr, r)
}
