package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pldcanfly/heimdall/storage"
)

type Handler struct {
	store storage.Storage
}

func NewHandler(store storage.Storage) *Handler {
	h := &Handler{store: store}

	return h
}

type ResponseGraph struct {
	ID        int
	Responses []Response
	Max       int
}

type Response struct {
	Online              bool
	ResponseTime        int
	ResponseTimePercent int
}

type RootData struct {
	Watchers []WatcherData
}

type IDRequest struct {
	ID int `param:"id"`
}

type WatcherData struct {
	Watcher       storage.WatcherRecord
	ResponseGraph ResponseGraph
}

// GET /
func (h *Handler) GetRoot(c echo.Context) error {
	res := make([]WatcherData, 0)
	watchers, err := h.store.GetWatchers()
	if err != nil {
		return err
	}

	for i := range watchers {
		rg, err := buildReponseGraph(h.store, watchers[i].ID)
		if err != nil {
			return err
		}

		res = append(res, WatcherData{
			Watcher:       watchers[i],
			ResponseGraph: rg,
		})

	}

	return c.Render(http.StatusOK, "index.gohtml", RootData{
		Watchers: res,
	})
}

func (h *Handler) GetResponsegraph(c echo.Context) error {

	var request IDRequest
	err := c.Bind(&request)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	rg, err := buildReponseGraph(h.store, request.ID)
	if err != nil {
		return err
	}

	//h.executeTemplate(w, "responsegraph.gohtml", rg)

	return c.Render(http.StatusOK, "responsegraph.gohtml", rg)
}

func (h *Handler) GetTest(c echo.Context) error {

	// executeComponentTemplate(w, "button", nil)
	return c.JSON(http.StatusOK, "test")
}
