package handler

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		w.Write([]byte("bad"))
	}
}

func render(w http.ResponseWriter, data map[string]interface{}, patterns ...string) error {
	if data["title"] == nil {
		data["title"] = "Gluten Free Horoscopes"
	}
	if data["contentOnly"] == true {
		return renderContent(w, data, patterns...)
	} else {
		w.Header().Set("Content-Type", "text/html")
		err := parse(patterns...).ExecuteTemplate(w, "base", data)
		if err != nil {
			http.Error(w, err.Error(), 0)
			log.Print("Template executing error: ", err)
		}
		return err
	}
}

func renderContent(w http.ResponseWriter, data map[string]interface{}, patterns ...string) error {
	w.Header().Set("Content-Type", "text/html")
	err := parse(patterns...).ExecuteTemplate(w, "main", data)
	if err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return err
}

func parse(patterns ...string) *template.Template {
	patterns = append(patterns, "layout.html")
	for i := 0; i < len(patterns); i++ {
		patterns[i] = "templates/" + patterns[i]
	}
	return template.Must(template.New("base").Funcs(funcs).ParseFiles(patterns...))
}

var funcs = template.FuncMap{
	"divide": func(a, b int) float32 {
		if a == 0 || b == 0 {
			return 0
		}
		return float32(a) / float32(b)
	},
	"progress": func(a, b int) float32 {
		if a == 0 || b == 0 {
			return 0
		}
		return float32(a) / float32(b) * 100
	},
	"add": func(a, b int) int {
		return a + b
	},
	"sub": func(a, b int) int {
		return a - b
	},
	"currentYear": func() int {
		return time.Now().UTC().Year()
	},
	"svg": func(path string) template.HTML {
		return inlineSVG(path)
	},
	"cssversion": func() string {
		file, err := os.Stat("assets/style.css")
		if err != nil {
			fmt.Println(err)
		}
		modifiedtime := file.ModTime().Nanosecond()
		return fmt.Sprint(modifiedtime)
	},
	"unescape": func(s string) template.HTML {
		return template.HTML(s)
	},
	"rating": func(n int) int {
		return rand.Intn(n-1) + 1
	},
}

func inlineSVG(path string) template.HTML {
	b, err := ioutil.ReadFile("assets/" + path)
	if err != nil {
		log.Println("includeHTML - error reading file: ", err)
		return ""
	}

	return template.HTML(string(b))
}
