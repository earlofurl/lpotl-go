package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/rs/zerolog/log"
	"strings"
	"sync"
	"time"
)

func Scrape(wg *sync.WaitGroup, out chan<- *ScrapedEpisode) error {
	defer wg.Done()
	scraperName := "lpotl"
	podcastName := "Last Podcast On The Left"
	logScrapeStart(scraperName, podcastName)

	//episodeCollector := createCollector("lastpodcastontheleft.com, www.lastpodcastontheleft.com")
	//siteCollector := createCollector("lastpodcastontheleft.com, www.lastpodcastontheleft.com")
	episodeCollector := createCollector()
	siteCollector := createCollector()

	episodeCollector.OnHTML(`html`, func(e *colly.HTMLElement) {
		ep := &ScrapedEpisode{}
		ep.TranscriptUrl = e.Request.URL.String()
		// TODO: Episode Number
		// TODO: Subseries Name

		// Title
		e.ForEach(`a.u-url`, func(id int, e *colly.HTMLElement) {
			ep.Name = strings.TrimSpace(e.Text)
		})

		// Description
		e.ForEach(`rel="bookmark"`, func(id int, e *colly.HTMLElement) {
			ep.Description = strings.TrimSpace(e.Text)
		})

		// Transcript Body
		e.ForEach(`div.sqs-block-content`, func(id int, e *colly.HTMLElement) {
			ep.Body = strings.TrimSpace(e.Text)
		})

		// Date
		e.ForEach(`a.entry-dateline-link`, func(id int, e *colly.HTMLElement) {
			var err error
			ep.ReleaseDate, err = time.Parse("January 2, 2006", strings.TrimSpace(e.Text))
			if err != nil {
				log.Err(err).Msg("Error parsing date")
			}
		})

		out <- ep
	})

	siteCollector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// assign the href attribute to the "link" variable for later use
		link := e.Attr("href")

		// We only want the links marked as "next", to go back chronologically
		if e.Attr("rel") != "next" {
			//log.Printf("Skipping link: %q -> %s", e.Text, link)
			return
		}

		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		// Visit link to next page
		err := siteCollector.Visit("https://lastpodcastontheleft.com" + link)
		if err != nil {
			log.Err(err).Msgf("Error visiting link: %q -> %s", e.Text, link)
		}
	})

	//siteCollector.OnHTML(`div.item-col.-video a`, func(e *colly.HTMLElement) {
	//	sceneURL := e.Request.AbsoluteURL(e.Attr("href"))
	//
	//	// If scene exist in database, there's no need to scraper
	//	if !funk.ContainsString(knownScenes, sceneURL) {
	//		sceneCollector.Visit(sceneURL)
	//	}
	//})

	err := siteCollector.Visit("https://lastpodcastontheleft.com/episodetranscripts/")
	if err != nil {
		return err
	}

	//if updateSite {
	//	updateSiteLastUpdate(scraperID)
	//}
	logScrapeFinished(scraperName, podcastName)
	return nil
}

func init() {
	registerScraper("lpotl", "Last Podcast On The Left", Scrape)
}
