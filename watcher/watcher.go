package watcher

import "time"

type Watcher interface {
	Ping()
}

type WatchReport struct {
	Online       bool
	ResponseTime time.Duration
}
