package main

import (
	"github.com/pldcanfly/heimdall/server"
)

func main() {
	// w := watcher.HTTPWatcher{Url: "https://www.matukas.de"}
	// err := w.Ping()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Printf("%v", w)

	server := server.NewServer("127.0.0.1:8080")
	server.Serve()

}
