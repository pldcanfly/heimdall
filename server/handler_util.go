package server

import (
	"math"
	"time"

	"github.com/pldcanfly/heimdall/storage"
)

func buildReponseGraph(s storage.Storage, watcher int) (ResponseGraph, error) {

	tps := make([]Response, 0)
	max := 0
	responses, err := s.GetLastResponses(watcher, 10)
	if err != nil {
		return ResponseGraph{}, err
	}

	for i := range responses {
		tps = append(tps, Response{
			Online:       responses[i].Online,
			ResponseTime: int(time.Duration.Milliseconds(responses[i].ReponseTime)),
		})
		if max < tps[i].ResponseTime {
			max = tps[i].ResponseTime
		}

	}

	for i := range tps {
		tps[i].ResponseTimePercent = int(math.Round(float64(tps[i].ResponseTime) / float64(max) * 100))
	}

	return ResponseGraph{
		ID:        watcher,
		Responses: tps,
		Max:       int(max),
	}, nil

}
