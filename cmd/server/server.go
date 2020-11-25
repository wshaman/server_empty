package main

import (
	"github.com/wshaman/server_empty/api/server"
	"log"
	"net/http"
)

func main()  {
	if err := serve(); err != nil {
		log.Fatal(err)
	}
}

func serve() error{
	r, err := server.NewServer("/api/v1")
	if err != nil {
		return err
	}
	h := http.Server{
		Addr:              "",
		Handler:           r,
	}
	if err = h.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
