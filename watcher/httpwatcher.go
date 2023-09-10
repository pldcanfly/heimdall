package watcher

import (
	"fmt"
	"net/http"
	"time"
)

type HTTPWatcher struct {
	Url string
}

func (w *HTTPWatcher) Ping() error {
	start := time.Now()
	res, err := http.Get(w.Url)
	if err != nil {
		return fmt.Errorf("httpwatcher: Ping failed %v", err)
	}
	defer res.Body.Close()

	fmt.Printf("%v responded with: %v %v\n", w.Url, res.StatusCode, time.Since(start))
	return nil
}
