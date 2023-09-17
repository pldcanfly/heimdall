package server

import (
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pldcanfly/heimdall/storage"
)

type ResponseGraph struct {
	ID        int
	Responses []Response
	Max       int
}

type Response struct {
	Online              bool
	ResponseTime        time.Duration
	ResponseTimePercent int
}

type RootData struct {
	Watchers []WatcherData
}

type WatcherData struct {
	Watcher       storage.WatcherRecord
	ResponseGraph ResponseGraph
}

func handleGetRoot(s *Server, w http.ResponseWriter, r *http.Request) error {
	res := make([]WatcherData, 0)
	watchers, err := s.store.GetWatchers()
	if err != nil {
		return err
	}

	for i := range watchers {
		rg, err := buildReponseGraph(s, watchers[i].ID, r)
		if err != nil {
			return err
		}

		res = append(res, WatcherData{
			Watcher:       watchers[i],
			ResponseGraph: rg,
		})

	}

	s.executeTemplate(w, "index.gohtml", RootData{
		Watchers: res,
	})
	return nil
}

func buildReponseGraph(s *Server, watcher int, r *http.Request) (ResponseGraph, error) {

	tps := make([]Response, 0)
	max := time.Duration(0)
	responses, err := s.store.GetLastResponses(watcher, 10)
	if err != nil {
		return ResponseGraph{}, err
	}

	for i := range responses {
		tps = append(tps, Response{
			Online:       responses[i].Online,
			ResponseTime: responses[i].ReponseTime,
		})
		if max < tps[i].ResponseTime {
			max = tps[i].ResponseTime
		}

	}

	for i := range tps {
		tps[i].ResponseTimePercent = int(math.Round(float64(time.Duration.Milliseconds(tps[i].ResponseTime)) / float64(time.Duration.Milliseconds(max)) * 100))
	}

	return ResponseGraph{
		ID:        watcher,
		Responses: tps,
		Max:       int(time.Duration.Milliseconds(max)),
	}, nil

}

func handleGetResponsegraph(s *Server, w http.ResponseWriter, r *http.Request) error {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return err
	}

	rg, err := buildReponseGraph(s, id, r)
	if err != nil {
		return err
	}

	s.executeTemplate(w, "responsegraph.gohtml", rg)

	return nil
}

func handleGetTest(s *Server, w http.ResponseWriter, r *http.Request) error {
	// executeComponentTemplate(w, "button", nil)
	return nil
}
