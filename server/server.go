package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pldcanfly/heimdall/storage"
)

type Server struct {
	listenAddr string

	store  storage.Storage
	router *echo.Echo
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{listenAddr: listenAddr, store: store}
}

func (s *Server) initRouter() error {
	h := NewHandler(s.store)
	e := echo.New()
	e.Renderer = NewTemplate()

	e.Static("/static", "web/static")
	e.GET("/", h.GetRoot)
	e.GET("/test", h.GetTest)
	e.GET("/api/components/responsegraph/:id", h.GetResponsegraph)

	s.router = e

	return nil
}

func (s *Server) Serve() {

	err := s.initRouter()
	if err != nil {
		fmt.Println(err)
		panic("couldn't init router")
	}

	s.router.Logger.Fatal(s.router.Start(":1323"))
	if err := http.ListenAndServe(s.listenAddr, s.router); err != nil {
		fmt.Printf("%v", err)
		panic("could not start listener")
	}
}

// func (s *Server) makeHandleFunc(f Handler) (c echo.Context) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		err := f(s, w, r)
// 		if err != nil {
// 			log.Printf("handler: %v\n", err)
// 			return
// 		}
// 	}

// }

func print404(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(404)
	w.Write([]byte("404 not found"))
}
