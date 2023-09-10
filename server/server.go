package server

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Server struct {
	listenAddr string
}

type Handler = func(w http.ResponseWriter, r *http.Request) error

func NewServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}

func (s *Server) Serve() {
	http.HandleFunc("/", s.makeHandleFunc(handleGetRoot))
	http.HandleFunc("/test", s.makeHandleFunc(handleGetTest))

	if err := http.ListenAndServe(s.listenAddr, nil); err != nil {
		fmt.Printf("%v", err)
		panic("could not start listener")
	}
}

func (s *Server) makeHandleFunc(f Handler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			log.Printf("handler: %v\n", err)
			return
		}
	}

}

func print404(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(404)
	w.Write([]byte("404 not found"))
}

func executePageTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseGlob("templates/base/*")
	if err != nil {
		fmt.Printf("base templates: %v\n", err)
		print404(w)
		return
	}

	t, err = t.ParseFiles("templates/" + tmpl + ".gohtml")
	if err != nil {
		fmt.Printf("page template: %v\n", err)
		print404(w)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)

	err = t.ExecuteTemplate(w, "base.gohtml", data)
	if err != nil {
		fmt.Printf("execute template: %v\n", err)
		print404(w)
		return
	}

}

func executeComponentTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("templates/components/" + tmpl + ".gohtml")
	if err != nil {
		fmt.Printf("component template: %v\n", err)
		print404(w)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)

	err = t.Execute(w, data)
	if err != nil {
		fmt.Printf("execute component template: %v\n", err)
		print404(w)
		return
	}
}
