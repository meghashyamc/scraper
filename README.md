# scraper
Scrapes some product specific details from amazon.com's product pages and stores them/retrieves them from a database

Offers two endpoints:

**:9000/scrape** - This endpoint requires a request body JSON like so:
{"url": "https://www.amazon.com/PlayStation-4-Pro-1TB-Console/dp/B01LOP8EZC"}

**:9000/getscraped** - This endpoint also requires a request body JSON exactly like /scrape.

/scrape extracts some product info from the amazon.com product page specified like name of the product, price, image URL etc. 

**Note:** Only amazon.com's product pages (like the URL mentioned above) are supported.

### Pre-requisites

To run scraper, the following must also be running:

* Redis on port 6379
* The binary from github.com/meghashyamc/dbconnect on port 9050

**To run**, migrate to the cmd directory and type go run main.go in a terminal.
