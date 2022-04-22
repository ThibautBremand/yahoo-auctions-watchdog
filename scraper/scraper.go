package scraper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
	"yahoo-auctions-watchdog/cache"
	"yahoo-auctions-watchdog/web"
)

const dateLayout = "2006-01-02 15:04:05"

type Scraper struct {
	URLs           []string
	Changerate     float64
	DownloadPhotos bool
}

func NewScraper(URLs []string, changerate float64, downloadPhotos bool) *Scraper {
	return &Scraper{
		URLs:           URLs,
		Changerate:     changerate,
		DownloadPhotos: downloadPhotos,
	}
}

type Listing struct {
	URL                 string  `json:"url"`
	Title               string  `json:"title"`
	PriceY              string  `json:"price_y"`
	PriceCurrency       string  `json:"price_currency"`
	BuyNowPriceY        string  `json:"buy_now_price_y"`
	BuyNowPriceCurrency string  `json:"buy_now_price_currency"`
	ID                  string  `json:"id"`
	Endtime             string  `json:"endtime"`
	SellerID            string  `json:"seller_id"`
	ImagePath           *string `json:"image_path,omitempty"`
}

// Scrape starts the scraping for the URLs that are configured.
// It returns a list of Listing, to be sent to Telegram. It also returns a map[string]Listing which will be used when
// updating the cache.
func (s *Scraper) Scrape(cache map[string]cache.CachedIDs) (
	[]Listing,
	map[string][]string,
	error,
) {
	log.Println("Scraping new listings")
	listings, lastItemIDs, err := s.scrapeListings(cache)
	if err != nil {
		return nil, nil, fmt.Errorf("could not start scraping: %v", err)
	}

	log.Printf("Got %d new listings!\n", len(listings))
	return listings, lastItemIDs, nil
}

func (s *Scraper) scrapeListings(
	scraped map[string]cache.CachedIDs,
) (
	[]Listing,
	map[string][]string,
	error,
) {
	var pulledListings []Listing
	lastItemIDs := make(map[string][]string)

	// Used to dedupe the final list of scraped products
	knownListings := make(map[string]int)

	for _, searchURL := range s.URLs {
		log.Printf("Searching with url %s\n", searchURL)
		doc, err := web.Get(searchURL)
		if err != nil {
			log.Printf("could not make request to search URL page %s: %s\n", searchURL, err)
			continue
		}

		if doc == nil {
			log.Printf("received an empty result for URL %s, skipping...\n", searchURL)
			continue
		}

		productsGrid := doc.Find("div.Products--grid")
		if productsGrid == nil {
			log.Printf("could not find srp-river-results div for URL %s, skipping...\n", searchURL)
			continue
		}

		productsList := productsGrid.Find("li.Product")
		if productsList == nil {
			log.Printf("received zero items for URL %s, skipping...\n", searchURL)
			continue
		}

		lastItemIDs[searchURL] = make([]string, 0)

		productsList.EachWithBreak(func(i int, sel *goquery.Selection) bool {
			listing, b := s.parseItem(sel, scraped, searchURL)
			if listing != nil {
				if _, ok := knownListings[listing.URL]; !ok {
					// Add the listing to the results of scraped listings if the listing hasn't been already scraped
					// for another search URL.
					pulledListings = append(pulledListings, *listing)

					knownListings[listing.URL] = 1
				}

				if len(lastItemIDs[searchURL]) < cache.LastIDsSize {
					lastItemIDs[searchURL] = append(lastItemIDs[searchURL], listing.ID)
				}
			}

			return b
		})

		// We space each queries just in case, to prevent getting throttled
		time.Sleep(2 * time.Second)
	}

	return pulledListings, lastItemIDs, nil
}

func (s *Scraper) parseItem(
	sel *goquery.Selection,
	scraped map[string]cache.CachedIDs,
	searchUrl string,
) (*Listing, bool) {
	_, isKnownURL := scraped[searchUrl]

	URL, err := parseURL(sel)
	if err != nil {
		log.Printf("error while parsing URL: %s, skipping...\n", err)
		return nil, false
	}

	ID, err := parseProductID(sel)
	if err != nil {
		log.Printf("error while parsing product ID from URL %s: %s, skipping...\n", URL, err)
		return nil, false
	}

	if isAlreadyScrapedListing(ID, scraped, searchUrl) {
		log.Println("Stop - Reached a listing that has already been scraped!")
		return nil, false
	}

	title, err := parseTitle(sel)
	if err != nil {
		log.Printf("error while parsing title: %s, skipping...\n", err)
		return nil, false
	}

	price, err := parsePrice(sel)
	if err != nil {
		log.Printf("error while parsing price: %s, skipping...\n", err)
		return nil, false
	}

	// Since the buy now price and sellerID are optional, we do not handle the errors here.
	buyNowPrice, _ := parseBuyNowPrice(sel)

	sellerID, _ := parseSellerID(sel)

	endtime, err := parseEndtime(sel)

	priceCurrency, err := changeCurrency(price, s.Changerate)
	if err != nil {
		log.Printf("could not convert currency for %s: %s, setting it as empty...", price, err)
	}

	buyNowPriceCurrency, err := changeCurrency(buyNowPrice, s.Changerate)
	if err != nil {
		log.Printf("could not convert currency for %s: %s, setting it as empty...", price, err)
	}

	image, err := parseImageURL(sel, ID, s.DownloadPhotos)
	if err != nil {
		log.Printf("could not find image for product %s, setting it as empty...", URL)
	}

	listing := Listing{
		ID:                  ID,
		URL:                 URL,
		Title:               title,
		PriceY:              price,
		PriceCurrency:       priceCurrency,
		BuyNowPriceY:        buyNowPrice,
		BuyNowPriceCurrency: buyNowPriceCurrency,
		Endtime:             endtime,
		SellerID:            sellerID,
		ImagePath:           image,
	}

	log.Printf("Successfully scraped 1 listing details (ID: %s)\n", listing.ID)

	if !isKnownURL {
		// This was the first time scraping this searchUrl. As we only want to check for new listings,
		// we won't scrape all the next listings and we will just wait for new ones. This is why we
		// will break out of the loop.
		return &listing, false
	}

	return &listing, true
}

// isAlreadyScrapedListing returns true if the current listing product ID has already been scraped and is in the cache.
func isAlreadyScrapedListing(
	productID string,
	scraped map[string]cache.CachedIDs,
	searchUrl string,
) bool {
	cachedListing, isKnownURL := scraped[searchUrl]
	if !isKnownURL {
		return false
	}

	return strings.Contains(cachedListing.LastIDs, productID)
}

func parseURL(sel *goquery.Selection) (string, error) {
	titleLink := sel.Find("a.Product__titleLink")
	if titleLink == nil {
		return "", fmt.Errorf("could not find URL")
	}

	URL, exists := titleLink.Attr("href")
	if !exists {
		return "", fmt.Errorf("could not find URL")
	}

	return URL, nil
}

func parseProductID(sel *goquery.Selection) (string, error) {
	productBonus := sel.Find("div.Product__bonus ")
	if productBonus == nil {
		return "", fmt.Errorf("could not find product ID")
	}

	productID, exists := productBonus.Attr("data-auction-id")
	if !exists {
		return "", fmt.Errorf("could not find product ID")
	}

	return productID, nil
}

func parseTitle(sel *goquery.Selection) (string, error) {
	titleLink := sel.Find("a.Product__titleLink")
	if titleLink == nil {
		return "", fmt.Errorf("could not find URL")
	}

	title, exists := titleLink.Attr("data-auction-title")
	if !exists {
		return "", fmt.Errorf("could not find URL")
	}

	return title, nil
}

func parsePrice(sel *goquery.Selection) (string, error) {
	productBonus := sel.Find("div.Product__bonus ")
	if productBonus == nil {
		return "", fmt.Errorf("could not find price")
	}

	price, exists := productBonus.Attr("data-auction-price")
	if !exists {
		return "", fmt.Errorf("could not find price")
	}

	return price, nil
}

func parseBuyNowPrice(sel *goquery.Selection) (string, error) {
	productBonus := sel.Find("div.Product__bonus ")
	if productBonus == nil {
		return "", fmt.Errorf("could not find buy now price")
	}

	buyNowPrice, exists := productBonus.Attr("data-auction-buynowprice")
	if !exists {
		return "", fmt.Errorf("could not find buy now price")
	}

	return buyNowPrice, nil
}

func parseSellerID(sel *goquery.Selection) (string, error) {
	productBonus := sel.Find("div.Product__bonus ")
	if productBonus == nil {
		return "", fmt.Errorf("could not find seller ID")
	}

	sellerID, exists := productBonus.Attr("data-auction-sellerid")
	if !exists {
		return "", fmt.Errorf("could not find seller ID")
	}

	return sellerID, nil
}

func parseEndtime(sel *goquery.Selection) (string, error) {
	productBonus := sel.Find("div.Product__bonus ")
	if productBonus == nil {
		return "", fmt.Errorf("could not find timestamp")
	}

	timestamp, exists := productBonus.Attr("data-auction-endtime")
	if !exists {
		return "", fmt.Errorf("could not find timestamp")
	}

	res, err := formatTimestamp(timestamp)
	if err != nil {
		return "", fmt.Errorf("error while parsing timestamp: %s", err)
	}

	return res, nil
}

func parseImageURL(sel *goquery.Selection, listingID string, downloadPhotos bool) (*string, error) {
	if !downloadPhotos {
		return nil, nil
	}

	img := sel.Find("img.Product__imageData")
	if img == nil {
		return nil, fmt.Errorf("could not find image")
	}

	imgURL, exists := img.Attr("src")
	if !exists {
		return nil, fmt.Errorf("could not find image")
	}

	fileName, err := DownloadFile(imgURL, fmt.Sprintf("./temp/%s", listingID))
	if err != nil {
		return nil, fmt.Errorf("could not download image as base64: %s", err)
	}

	return &fileName, nil
}

func formatTimestamp(timestamp string) (string, error) {
	n, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return "", fmt.Errorf("error while formatting timestamp %s: %s", timestamp, err)
	}

	t := time.Unix(n, 0)

	return t.Format(dateLayout), nil
}

func changeCurrency(yenValue string, changerate float64) (string, error) {
	f, err := strconv.ParseFloat(yenValue, 64)
	if err != nil {
		return "", fmt.Errorf("error while changing currency for %s: %s", yenValue, err)
	}

	converted := f * changerate
	rounded := math.Round(converted)

	return strconv.FormatFloat(rounded, 'f', -1, 64), nil
}
