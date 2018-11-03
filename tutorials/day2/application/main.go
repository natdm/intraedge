package main

import (
	"log"
	"net/http"

	"github.com/natdm/intraedge/tutorials/day2/application/api"
	storage "github.com/natdm/intraedge/tutorials/day2/application/store"
)

func main() {
	store, err := storage.New()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := store.Close(); err != nil {
			log.Fatalf("unable to close store: %s\n", err)
		} else {
			log.Println("storage closed")
		}
	}()

	storageAPI, err := api.New("/file", store)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/file/", api.MiddlwareLogger(storageAPI))
	http.ListenAndServe(":8080", mux)
}
