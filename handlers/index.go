package handler

import (
	"net/http"
)

// IndexHandler serves the front page
func IndexHandler(w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})
	if r.Header.Get("HX-Request") == "true" {
		data["contentOnly"] = true
	}
	return render(w, data, "index.html")
}
