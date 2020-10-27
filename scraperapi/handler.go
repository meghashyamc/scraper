package scraperapi

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func ScrapeAndStore(w http.ResponseWriter, req *http.Request) {

	urlReq := URLReq{}
	if err := json.NewDecoder(req.Body).Decode(&urlReq); err != nil {
		writeResponse(&w, err.Error(), errorStatus, http.StatusInternalServerError)
		return
	}

	statusCode, urlDetails, err := scrapeProductDetails(urlReq.URL)
	if err != nil {
		writeResponse(&w, err.Error(), errorStatus, statusCode)
		return
	}
	log.Println("received details after scraping", urlDetails)

	statusCode, timestamp, err := persistInDB(urlReq.URL, urlDetails)
	if err != nil {
		writeResponse(&w, err.Error(), errorStatus, statusCode)
		return
	}
	if timestamp == nil {
		log.Fatal("nil timestamp received while persisting in DB")
	}
	log.Println("successfully persisted in db")

	urlDetails.TimeStamp = *timestamp
	writeResponse(&w, urlDetails, successStatus, statusCode)

}

func GetScrapedURLInfo(w http.ResponseWriter, req *http.Request) {
	urlReq := URLReq{}
	if err := json.NewDecoder(req.Body).Decode(&urlReq); err != nil {
		writeResponse(&w, err.Error(), errorStatus, http.StatusInternalServerError)
		return
	}
	statusCode, scrapedInfoBytes, err := getFromDB(scrapeddocs, urlReq.URL)
	if err != nil {
		writeResponse(&w, err.Error(), errorStatus, statusCode)
		return
	}

	urlDetails := URLDetails{}
	if err := json.NewDecoder(bytes.NewBuffer(scrapedInfoBytes)).Decode(&urlDetails); err != nil {
		writeResponse(&w, err.Error(), errorStatus, http.StatusInternalServerError)
		return
	}

	writeResponse(&w, urlDetails, successStatus, statusCode)

}
