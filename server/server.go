package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	listenAddr string
}

type Handler = func(w http.ResponseWriter, r *http.Request) error

func NewServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}

func (s *Server) Serve() {
	r := chi.NewRouter()
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "frontpage", "static"))
	FileServer(r, "/static", filesDir)

	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	r.Get("/", s.makeHandleFunc(handleGetRoot))
	r.Get("/test", s.makeHandleFunc(handleGetTest))
	r.Route("/components", func(r chi.Router) {
		r.Get("/responsegraph", s.makeHandleFunc(handleGetResponsegraph))
	})

	if err := http.ListenAndServe(s.listenAddr, r); err != nil {
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
	t, err := template.ParseGlob("frontpage/templates/base/*")
	if err != nil {
		fmt.Printf("base templates: %v\n", err)
		print404(w)
		return
	}

	t, err = t.ParseFiles("frontpage/templates/" + tmpl + ".gohtml")
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
	t, err := template.ParseFiles("frontpage/components/" + tmpl + ".gohtml")
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

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
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
