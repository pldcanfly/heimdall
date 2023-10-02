package storage

import "time"

type Storage interface {
	Close()
	GetWatchers() ([]WatcherRecord, error)
	InsertWatcher(url string, watcher string) (WatcherRecord, error)
	GetLastResponses(watcher int, len int) ([]ResponseRecord, error)
	InsertResponse(watcher int, online bool, responsetime time.Duration) error
}

type WatcherRecord struct {
	ID      int
	URL     string
	Watcher string
}

type ResponseRecord struct {
	Watcher     int
	Online      bool
	ReponseTime time.Duration
}
