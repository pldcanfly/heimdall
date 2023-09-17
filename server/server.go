package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/pldcanfly/heimdall/storage"
)

type Server struct {
	listenAddr string
	Templates  *template.Template
	store      storage.Storage
	router     http.Handler
}

type Handler = func(*Server, http.ResponseWriter, *http.Request) error

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{listenAddr: listenAddr, store: store}
}

func (s *Server) initTemplates() error {
	t, err := parseTemplates()
	if err != nil {
		return err
	}
	s.Templates = t
	return nil
}

func (s *Server) initRouter() error {
	r := mux.NewRouter()
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "web", "static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(filesDir))))

	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	r.HandleFunc("/", s.makeHandleFunc(handleGetRoot))
	r.HandleFunc("/test", s.makeHandleFunc(handleGetTest))
	r.HandleFunc("/api/components/responsegraph/{id}", s.makeHandleFunc(handleGetResponsegraph))

	s.router = r

	return nil
}

func (s *Server) Serve() {

	err := s.initTemplates()
	if err != nil {
		fmt.Println(err)
		panic("couldn't init templates")
	}

	err = s.initRouter()
	if err != nil {
		fmt.Println(err)
		panic("couldn't init router")
	}

	if err := http.ListenAndServe(s.listenAddr, s.router); err != nil {
		fmt.Printf("%v", err)
		panic("could not start listener")
	}
}

func (s *Server) makeHandleFunc(f Handler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(s, w, r)
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

func parseTemplates() (*template.Template, error) {
	t, err := template.ParseGlob("web/templates/base/*")
	if err != nil {

		return nil, fmt.Errorf("base templates: %v", err)
	}

	t, err = t.ParseGlob("web/components/*")
	if err != nil {

		return nil, fmt.Errorf("component templates: %v", err)
	}

	t, err = t.ParseGlob("web/templates/pages/*")
	if err != nil {

		return nil, fmt.Errorf("page templates: %v", err)
	}

	return t, nil

}

func (s *Server) executeTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)

	err := s.Templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		fmt.Printf("execute template: %v\n", err)
		print404(w)
		return
	}
}

// func FileServer(r chi.Router, path string, root http.FileSystem) {
// 	if strings.ContainsAny(path, "{}*") {
// 		panic("FileServer does not permit any URL parameters.")
// 	}

// 	if path != "/" && path[len(path)-1] != '/' {
// 		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
// 		path += "/"
// 	}
// 	path += "*"

// 	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
// 		rctx := chi.RouteContext(r.Context())
// 		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
// 		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
// 		fs.ServeHTTP(w, r)
// 	})
// }
