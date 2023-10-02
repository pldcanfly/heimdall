package watcher

import (
	"fmt"
	"time"

	"github.com/pldcanfly/heimdall/storage"
)

type Watcher interface {
	Ping() error
	Store(report WatchReport) error
	Watch()
}

type WatchReport struct {
	Online       bool
	ResponseTime time.Duration
}

type WatchMaster struct {
	store    storage.Storage
	Watchers []Watcher
}

func NewWatchMaster(store storage.Storage) (*WatchMaster, error) {
	return &WatchMaster{store: store}, nil
}

func (w *WatchMaster) Watch() {
	w.initWatchers()

	for _, watcher := range w.Watchers {
		go watcher.Watch()
	}
}

func (w *WatchMaster) AddWatcher(watcher Watcher) {
	w.Watchers = append(w.Watchers, watcher)
}

func (w *WatchMaster) initWatchers() {
	watchers, err := w.store.GetWatchers()
	if err != nil {
		panic(err)
	}

	fmt.Println(watchers)

	for i := range watchers {
		w.AddWatcher(NewHTTPWatcher(watchers[i].ID, watchers[i].URL, w.store))

	}

}
