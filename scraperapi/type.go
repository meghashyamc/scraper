package scraperapi

import "time"

const (
	errorStatus      = "error"
	successStatus    = "success"
	urlNotRecognized = "unexpected response from URL"
	requestTimeout   = 30 * time.Second
	currency         = "$"
)
const Port = "9000"

const scrapeddocs = "scrapeddocs"
const dbURL = "http://localhost:"
const dbAPIPort = "9050"
const (
	dbaddEndpoint = "/addtodb"
	dbgetEndpoint = "/getfromdb"
)

type URLReq struct {
	URL string `json:"url"`
}
type URLDetails struct {
	Product   ProductDetails `json:"product"`
	TimeStamp time.Time      `json:"timestamp,omitempty"`
}

type ProductDetails struct {
	Name         string `json:"name"`
	ImageURL     string `json:"imageURL"`
	Description  string `json:"description"`
	Price        string `json:"price"`
	TotalReviews int    `json:"totalReviews"`
}

type apiResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type Dbpayload struct {
	Key   string
	Field string
	Value string
}
