package scraperapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func persistInDB(URL string, urlDetails *URLDetails) (int, *time.Time, error) {

	if urlDetails == nil {
		return http.StatusBadRequest, nil, errors.New("can't persist nil URL details")
	}

	urlDetails.TimeStamp = time.Now()
	value, err := json.Marshal(*urlDetails)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	field := URL
	key := scrapeddocs

	payloadBytes, err := json.Marshal(Dbpayload{Key: key, Field: field, Value: string(value)})
	resp, err := http.Post(dbURL+dbAPIPort+dbaddEndpoint, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}
		return resp.StatusCode, nil, errors.New(string(respBytes))
	}
	return resp.StatusCode, &urlDetails.TimeStamp, nil
}

func getFromDB(key, field string) (int, []byte, error) {

	payloadBytes, err := json.Marshal(Dbpayload{Key: key, Field: field})
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	resp, err := http.Post(dbURL+dbAPIPort+dbgetEndpoint, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, nil, errors.New(string(respBytes))
	}
	return resp.StatusCode, respBytes, nil

}
