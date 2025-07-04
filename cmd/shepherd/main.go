package main

import (
	"net/http"

	"github.com/alwindoss/prism"
	"github.com/alwindoss/shepherd"
	"github.com/go-chi/chi/v5"
)

// Handlers for different pages
func homeHandler(renderer prism.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, "index.page.html", "base", map[string]interface{}{
			"Title": "Home Page",
		})
	}
}

func aboutHandler(renderer prism.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, "about.page.html", "base", map[string]interface{}{
			"Title": "About Us",
		})
	}
}

func contactHandler(renderer prism.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, "contact.page.html", "base", map[string]interface{}{
			"Title": "Contact Us",
		})
	}
}

func main() {
	cfg := prism.Config{
		LayoutPath:   "web/layouts/*.html",
		PagesPath:    "web/pages/*.html",
		PartialsPath: "web/partials/*.html",
		FS:           shepherd.WebFS,
	}
	r := chi.NewMux()
	renderer := prism.New(&cfg)
	pubFS := http.FileServerFS(shepherd.PublicFS)
	r.Handle("/public/*", http.StripPrefix("/public/", pubFS))
	r.HandleFunc("/", homeHandler(renderer))
	r.HandleFunc("/about", aboutHandler(renderer))
	r.HandleFunc("/contact", contactHandler(renderer))

	// Serve on port 8080
	http.ListenAndServe(":8080", r)
}
