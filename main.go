package main

import (
	"fmt"
	"log"
	"net/http"

	config "WhistleNewsBackend/config"
	server "WhistleNewsBackend/server"
)

func main() {

	server.Init()
	router := server.NewRouter()
	fmt.Printf("listening to http://127.0.0.1:%s%s\n", config.ServerPort, " ...")
	log.Fatal(http.ListenAndServe("127.0.0.1:"+config.ServerPort, router))

}
