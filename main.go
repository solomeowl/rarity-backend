package main

import (
	"fmt"
	"log"
	"net/http"
	"rarity-backend/routers"
	"time"
)

func main() {
	router := routers.Init()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8844),
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
