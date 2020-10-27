package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/meghashyamc/scraper/scraperapi"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/scrape", scraperapi.ScrapeAndStore).Methods("POST")
	router.HandleFunc("/getscraped", scraperapi.GetScrapedURLInfo).Methods("POST")

	fmt.Println("scaper API listening on port", scraperapi.Port)
	http.ListenAndServe(":"+scraperapi.Port, router)
}
