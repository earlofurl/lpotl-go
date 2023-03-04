package scraper

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/earlofurl/lpotl-go/sqlc"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/rs/zerolog/log"
)

var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"

type Episode = sqlc.Episode

var scrapers []Scraper

type ScraperFunc func(*sync.WaitGroup, chan<- *ScrapedEpisode) error

type Scraper struct {
	ID     string
	Name   string
	Scrape ScraperFunc
}

type ScrapedEpisode struct {
	ID            int32     `json:"id"`
	Name          string    `json:"name"`
	NumberSeries  int32     `json:"number_series"`
	NumberOverall int32     `json:"number_overall"`
	ReleaseDate   time.Time `json:"release_date"`
	Description   string    `json:"description"`
	Body          string    `json:"body"`
	TranscriptUrl string    `json:"transcript_url"`
	PodcastID     int32     `json:"podcast_id"`
	SeriesID      int32     `json:"series_id"`
	LastUpdated   time.Time `json:"last_updated"`
	Headline      string    `json:"headline"`
}

func (s *ScrapedEpisode) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}

func (s *ScrapedEpisode) Log() error {
	j, err := json.MarshalIndent(s, "", "  ")
	log.Debug().Msgf("%v", string(j))
	return err
}

func GetScrapers() []Scraper {
	return scrapers
}

func RegisterScraper(id string, name string, f ScraperFunc) {
	s := Scraper{}
	s.ID = id
	s.Name = name
	s.Scrape = f
	scrapers = append(scrapers, s)
}

func createCollector(domains ...string) *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(domains...),
		colly.CacheDir(getScrapeCacheDir()),
		colly.UserAgent(UserAgent),
	)

	c = createCallbacks(c)
	return c
}

func cloneCollector(c *colly.Collector) *colly.Collector {
	x := c.Clone()
	x = createCallbacks(x)
	return x
}

func createCallbacks(c *colly.Collector) *colly.Collector {
	const maxRetries = 15

	c.OnRequest(func(r *colly.Request) {
		attempt := r.Ctx.GetAny("attempt")

		if attempt == nil {
			r.Ctx.Put("attempt", 1)
		}

		log.Print("visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		attempt := r.Ctx.GetAny("attempt").(int)

		if r.StatusCode == 429 {
			log.Err(err).Msgf("Error: %v", r.StatusCode)

			if attempt <= maxRetries {
				unCache(r.Request.URL.String(), c.CacheDir)
				log.Err(err).Msg("Waiting 2 seconds before next request...")
				r.Ctx.Put("attempt", attempt+1)
				time.Sleep(2 * time.Second)
				err := r.Request.Retry()
				if err != nil {
					return
				}
			}
		}
	})

	return c
}

func DeleteScrapeCache() error {
	return os.RemoveAll(getScrapeCacheDir())
}

func getScrapeCacheDir() string {
	//return config.ScrapeCacheDir
	// TODO: make this configurable
	return "./lpotl_cache"
}

func registerScraper(id string, name string, f ScraperFunc) {
	RegisterScraper(id, name, f)
}

func unCache(URL string, cacheDir string) {
	sum := sha1.Sum([]byte(URL))
	hash := hex.EncodeToString(sum[:])
	dir := path.Join(cacheDir, hash[:2])
	filename := path.Join(dir, hash)
	if err := os.Remove(filename); err != nil {
		log.Fatal().Err(err).Msgf("Error removing file %v", filename)
	}
}

func logScrapeStart(scraperName string, podcastName string) {
	log.Printf("Starting scraper %v for Podcast: %v", scraperName, podcastName)
}

func logScrapeFinished(scraperName string, podcastName string) {
	log.Printf("Finished scraper %v for Podcast: %v", scraperName, podcastName)
}

func getTextFromHTMLWithSelector(data string, sel string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating goquery document from reader")
	}
	return strings.TrimSpace(doc.Find(sel).Text())
}
