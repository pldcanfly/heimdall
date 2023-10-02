package main

import (
	"github.com/pldcanfly/heimdall/server"
	"github.com/pldcanfly/heimdall/storage"
	"github.com/pldcanfly/heimdall/watcher"
)

func main() {
	ch := make(chan struct{})
	store, err := storage.NewPostgresStore()
	if err != nil {
		panic(err)
	}
	defer store.Close()
	server := server.NewServer("127.0.0.1:8080", store)
	go server.Serve()
	watchmaster, err := watcher.NewWatchMaster(store)
	if err != nil {
		panic(err)
	}
	go watchmaster.Watch()
	<-ch
}
