package watcher

import "time"

type Watcher interface {
	Ping()
}

type WatchReport struct {
	StatusCode   int
	ResponseTime time.Duration
}
