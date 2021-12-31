// Gluten Free Horoscopes
package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	handler "github.com/nathanhollows/Lint/go/pkg/mod/github.com/nathanhollows/gluten-free-horoscopes/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/", handler.Handler(handler.IndexHandler))
	r.Handle("/starsign", handler.Handler(handler.StarsignHandler))
	r.Handle("/{starsign:(aries|taurus|gemini|sagittarius|pisces|aquarius|libra|virgo|scorpio|cancer|leo|capricorn)}", handler.Handler(handler.HoroscopeHandler))

	// Assets fileserver
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "assets"))
	fileServer(r, "/assets", filesDir)

	http.ListenAndServe(":8000", r)
}

// fileServer conveniently sets up a http.fileServer handler to serve
// static files from a http.FileSystem.
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
