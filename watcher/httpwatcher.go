package watcher

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pldcanfly/heimdall/storage"
)

type HTTPWatcher struct {
	ID    int
	URL   string
	store storage.Storage
}

func NewHTTPWatcher(id int, url string, store storage.Storage) *HTTPWatcher {
	return &HTTPWatcher{ID: id, URL: url, store: store}
}

func (w *HTTPWatcher) Ping() error {
	start := time.Now()
	res, err := http.Get(w.URL)
	if err != nil {
		return fmt.Errorf("httpwatcher: Ping failed %v", err)
	}
	defer res.Body.Close()

	w.Store(WatchReport{
		Online:       res.StatusCode < 400,
		ResponseTime: time.Since(start),
	})
	return nil
}

func (w *HTTPWatcher) Watch() {
	for {
		w.Ping()
		time.Sleep(30 * time.Second)
	}
}

func (w *HTTPWatcher) Store(report WatchReport) error {
	return w.store.InsertResponse(w.ID, report.Online, report.ResponseTime)
}
