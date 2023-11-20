package main

import (
	"weeklyNews/source"

	"github.com/gorilla/mux"
)

func AddServer(router *mux.Router) {
	router.HandleFunc("/getLastVideoFromCCTV", source.GetLatestVideoFromCCTV)
}
