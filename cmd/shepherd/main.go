package main

import (
	"net/http"

	"github.com/alwindoss/prism"
	"github.com/alwindoss/shepherd"
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
	renderer := prism.New(&cfg)
	http.HandleFunc("/", homeHandler(renderer))
	http.HandleFunc("/about", aboutHandler(renderer))
	http.HandleFunc("/contact", contactHandler(renderer))

	// Serve on port 8080
	http.ListenAndServe(":8080", nil)
}
