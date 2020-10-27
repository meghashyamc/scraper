package scraperapi

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

func scrapeProductDetails(urlToScrape string) (int, *URLDetails, error) {

	urlDetails := URLDetails{}

	client := &http.Client{
		Timeout: requestTimeout,
	}
	resp, err := client.Get(urlToScrape)
	if err != nil {
		return http.StatusNotFound, nil, errors.New(urlNotRecognized + ": " + err.Error())

	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return resp.StatusCode, nil, errors.New(urlNotRecognized)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	var productName, price string
	if productName = findProductName(doc); productName == "" {
		return http.StatusBadRequest, nil, errors.New(urlNotRecognized + ": couldn't identify product name")
	}
	if price = findPrice(doc); !strings.HasPrefix(price, currency) {
		return http.StatusBadRequest, nil, errors.New(urlNotRecognized + ": couldn't identify product price")
	}
	imageURL := findImageURL(doc)
	totalReviews := findReviewCount(doc)
	description := findDescription(doc)

	urlDetails.Product = ProductDetails{Name: productName, ImageURL: imageURL, Description: description, Price: price, TotalReviews: totalReviews}
	return resp.StatusCode, &urlDetails, nil
}

func findProductName(doc *goquery.Document) string {
	var name string
	doc.Find(".a-size-large.product-title-word-break").Each(func(i int, element *goquery.Selection) {
		if i > 0 {
			return
		}
		name = strings.TrimRight(strings.TrimLeft(element.Text(), "\n"), "\n")
	})
	return name

}

func findDescription(doc *goquery.Document) string {
	var description string

	doc.Find(".a-unordered-list.a-vertical.a-spacing-mini").Children().Each(func(i int, element *goquery.Selection) {

		if _, hasid := element.Attr("id"); !hasid {
			class, ok := element.Children().Attr("class")
			if ok && class == "a-list-item" {
				description += strings.TrimRight(strings.TrimLeft(element.Children().Text(), "\n"), "\n") + "\n"
			}
		}

	})

	return description
}

func findImageURL(doc *goquery.Document) string {
	var imageURL string
	doc.Find("#landingImage").Each(func(i int, element *goquery.Selection) {

		imageURLsStr, ok := element.Attr("data-a-dynamic-image")
		if ok {
			imageURLs := strings.SplitN(imageURLsStr, ",", 2)
			imageURL = strings.TrimLeft(imageURLs[0], "{\"")
			imageURL = trimRightToGetURL(imageURL)

		}
	})
	return imageURL
}

func findPrice(doc *goquery.Document) string {
	var price string
	doc.Find(".a-size-base.a-color-price").Each(func(i int, element *goquery.Selection) {

		if i > 0 && price != "" {
			return
		}
		price = element.Text()

	})

	if !strings.HasPrefix(price, currency) {
		price = ""
		doc.Find(".a-size-base.a-color-base").Each(func(i int, element *goquery.Selection) {

			if i > 0 && price != "" {
				return
			}
			price = element.Text()

		})
	}
	return price
}

func findReviewCount(doc *goquery.Document) int {
	var reviewCountStr string

	doc.Find("#acrCustomerReviewText").Each(func(i int, element *goquery.Selection) {

		if i > 0 {
			return
		}
		reviewCountStr = element.Text()

	})

	reviewCount, err := strconv.Atoi(keepOnlyNumbers(reviewCountStr))
	if err != nil {
		return -1
	}
	return reviewCount
}

func keepOnlyNumbers(s string) string {
	strings.TrimRightFunc(s, func(r rune) bool {
		return !unicode.IsNumber(r)
	})

	var onlyNumStr string
	for i := 0; i < len(s); i++ {
		if unicode.IsNumber(rune(s[i])) {
			onlyNumStr += string(s[i])
		}
	}
	return onlyNumStr
}

func trimRightToGetURL(s string) string {

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '"' {
			return s[:i]
		}
	}
	return ""
}
