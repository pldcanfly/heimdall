package storage

import (
	"time"
)

type MemoryStorage struct {
	Watchers  []WatcherRecord
	Responses []ResponseRecord
}

func NewMemorystore() (*MemoryStorage, error) {
	store := &MemoryStorage{
		Watchers: make([]WatcherRecord, 0),
	}

	store.InsertTestData()

	return store, nil
}

func (s *MemoryStorage) InsertTestData() {
	s.Watchers = append(s.Watchers, WatcherRecord{
		ID:      1,
		URL:     "https://www.google.com",
		Watcher: "http",
	})
	s.Watchers = append(s.Watchers, WatcherRecord{
		ID:      2,
		URL:     "https://www.matukal.de",
		Watcher: "http",
	})
	s.Watchers = append(s.Watchers, WatcherRecord{
		ID:      3,
		URL:     "https://www.styria.com",
		Watcher: "http",
	})
}

func (s *MemoryStorage) GetWatchers() ([]WatcherRecord, error) {
	return s.Watchers, nil
}

func (s *MemoryStorage) InsertWatcher(url string, watcher string) (WatcherRecord, error) {
	record := WatcherRecord{
		ID:      len(s.Watchers) + 1,
		URL:     url,
		Watcher: watcher,
	}

	s.Watchers = append(s.Watchers, record)

	return record, nil
}

func (s *MemoryStorage) GetLastResponses(watcher int, n int) ([]ResponseRecord, error) {

	res := make([]ResponseRecord, 0)
	for i := range s.Responses {
		if s.Responses[i].Watcher == watcher {
			res = append(res, s.Responses[i])
		}
	}

	if len(res) <= n {
		return res, nil
	}

	return res[len(res)-n:], nil
}

func (s *MemoryStorage) InsertResponse(watcher int, online bool, responsetime time.Duration) error {
	record := ResponseRecord{
		Watcher:     watcher,
		Online:      online,
		ReponseTime: responsetime,
	}
	s.Responses = append(s.Responses, record)
	return nil
}
