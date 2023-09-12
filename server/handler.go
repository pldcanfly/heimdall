package server

import (
	"math"
	"math/rand"
	"net/http"
	"time"
)

type ResponseGraph struct {
	Timepoints []Timepoint
	Max        int
}

type Timepoint struct {
	Online              bool
	ResponseTime        time.Duration
	ResponseTimePercent int
}

func handleGetRoot(w http.ResponseWriter, r *http.Request) error {
	executePageTemplate(w, "index", nil)
	return nil
}

func handleGetResponsegraph(w http.ResponseWriter, r *http.Request) error {
	tps := make([]Timepoint, 10)
	max := time.Duration(0)
	for i := range tps {
		tps[i].Online = true
		tps[i].ResponseTime = time.Millisecond * time.Duration(rand.Int31n(500))
		if max < tps[i].ResponseTime {
			max = tps[i].ResponseTime
		}
	}

	for i := range tps {

		tps[i].ResponseTimePercent = int(math.Round(float64(time.Duration.Milliseconds(tps[i].ResponseTime)) / float64(time.Duration.Milliseconds(max)) * 100))
	}

	executeComponentTemplate(w, "responsegraph", ResponseGraph{
		Timepoints: tps,
		Max:        int(time.Duration.Milliseconds(max)),
	})
	return nil
}

func handleGetTest(w http.ResponseWriter, r *http.Request) error {
	executeComponentTemplate(w, "button", nil)
	return nil
}
