package coordinator

import (
	"bytes"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
	"yahoo-auctions-watchdog/cache"
	"yahoo-auctions-watchdog/config"
	"yahoo-auctions-watchdog/scraper"
	"yahoo-auctions-watchdog/web"
)

type Coordinator struct {
	Scraper        *scraper.Scraper
	SleepPeriod    time.Duration
	Tpl            *template.Template
	TelegramClient *web.TelegramClient
}

func NewCoordinator(
	searches []config.Search,
	sleepPeriod time.Duration,
	tpl *template.Template,
	changerate float64,
	telegramClient *web.TelegramClient,
	downloadPhotos bool,
) *Coordinator {
	URLs := make([]string, len(searches))
	for i, search := range searches {
		URLs[i] = search.URL
	}

	s := scraper.NewScraper(URLs, changerate, downloadPhotos)
	return &Coordinator{
		Scraper:        s,
		SleepPeriod:    sleepPeriod,
		Tpl:            tpl,
		TelegramClient: telegramClient,
	}
}

func (c *Coordinator) Start(scraped map[string]cache.CachedIDs) {
	makeAndClearTempDir()

	for {
		listings, lastItemIDs, err := c.Scraper.Scrape(scraped)
		if err != nil {
			log.Println("error while scraping new listings, skipping", err)
			time.Sleep(c.SleepPeriod)
			continue
		}

		scraped = buildCache(lastItemIDs, scraped)

		err = cache.UpdateCache(scraped)
		if err != nil {
			log.Println("error while updating scraped URLs, skipping", err)
		}

		c.sendToTelegram(listings, c.Tpl)

		makeAndClearTempDir()
		time.Sleep(c.SleepPeriod)
	}
}

// buildCache returns a map[string]cache.CachedIDs, ready to be persisted into the cache, from the given
// map[string]scraper.Listing which comes from the last scraping, and the map[string]cache.CachedIDs which is the
// previous cache.
// It uses data from both maps to build the new cache.
func buildCache(lastItemIDs map[string][]string, scraped map[string]cache.CachedIDs) map[string]cache.CachedIDs {
	res := make(map[string]cache.CachedIDs)

	for searchURL, IDs := range lastItemIDs {
		if len(IDs) == 0 {
			continue
		}

		cachedListing, ok := scraped[searchURL]
		formattedIDs := strings.Join(IDs, cache.LastIDsSeparator)

		if !ok {
			// If it is the first time that we scrape this URL
			// Then we directly persist the whole list of IDs in the cache for this URL.
			toPersist := cache.CachedIDs{
				URL:     searchURL,
				LastIDs: formattedIDs,
			}

			res[searchURL] = toPersist
			continue
		}

		if len(IDs) >= cache.LastIDsSize {
			// If we have scraped more than 'LastIDsSize'
			// Then we directly persist the truncated list of IDs in the cache for this URL.

			formattedIDs = strings.Join(IDs[:cache.LastIDsSize], cache.LastIDsSeparator)
		} else {
			// We have scraped less than 'LastIDsSize' product IDs, so we need to retrieve some product IDs from the
			// cache before refreshing the cache.

			cachedIDs := cachedListing.LastIDs
			split := strings.Split(cachedIDs, cache.LastIDsSeparator)

			nbToKeep := cache.LastIDsSize - len(IDs)
			if nbToKeep > len(split) {
				formattedIDs = fmt.Sprintf("%s%s%s", formattedIDs, cache.LastIDsSeparator, cachedIDs)
			} else {
				toKeep := strings.Join(split[:nbToKeep], cache.LastIDsSeparator)
				formattedIDs = fmt.Sprintf("%s%s%s", formattedIDs, cache.LastIDsSeparator, toKeep)
			}
		}

		toPersist := cache.CachedIDs{
			URL:     searchURL,
			LastIDs: formattedIDs,
		}

		res[searchURL] = toPersist
	}

	// Do not remove from the cache the URLs that did not have new scraped items
	for k, v := range scraped {
		if _, ok := res[k]; ok {
			continue
		}
		res[k] = v
	}

	return res
}

func (c *Coordinator) sendToTelegram(listings []scraper.Listing, tpl *template.Template) {
	for _, listing := range listings {
		buf := &bytes.Buffer{}
		err := tpl.Execute(buf, listing)
		var msg string
		if err != nil {
			log.Println("could not execute template", err)
			msg = listing.URL
		} else {
			msg = buf.String()
		}

		// Double quotes are not correctly parsed by Telegram
		msg = strings.ReplaceAll(msg, `"`, "")

		// If an image has been downloaded, send the Telegram message as a photo + caption
		if listing.ImagePath != nil {
			photo := &tele.Photo{
				File:    tele.FromDisk(*listing.ImagePath),
				Caption: msg,
			}

			_, err := c.TelegramClient.Client.Send(c.TelegramClient.ChatID, photo)
			if err != nil {
				log.Println("could not send Telegram message", err)
			}

			continue
		}

		// No image has been downloaded: send the Telegram message as raw text
		_, err = c.TelegramClient.Client.Send(c.TelegramClient.ChatID, msg)
		if err != nil {
			log.Println("could not send Telegram message", err)
		}
	}
}

// makeAndClearTempDir creates or empties the directory that will contain downloaded photos.
func makeAndClearTempDir() {
	err := os.RemoveAll("./temp")
	if err != nil {
		log.Fatalf("could not clear temporary directory: %s", err)
	}

	err = os.MkdirAll(filepath.Join(".", "temp"), os.ModePerm)
	if err != nil {
		log.Fatalf("could not create temporary directory")
	}
}
