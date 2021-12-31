package handler

import (
	"net/http"
)

// StarsignHandler serves the "What's my starsign?" page
func StarsignHandler(w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})
	if r.Header.Get("HX-Request") == "true" {
		data["contentOnly"] = true
	}
	return render(w, data, "starsign.html")
}
