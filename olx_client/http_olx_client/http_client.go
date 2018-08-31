package http_olx_client

import (
	"net/http"
	"strings"

	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/entities"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/olx_client/dateParser"
	"github.com/PuerkitoBio/goquery"
	log "github.com/Sirupsen/logrus"
)

type OlxHttpClient struct {
	httpClient *http.Client
}

func NewOLXHTTPClient(hc *http.Client)*OlxHttpClient{
	return &OlxHttpClient{
		httpClient:hc,
	}
}

type Result struct {
	URL  string
	Date string
}

func (c *OlxHttpClient) GetDocumentByUrl(url string) *goquery.Document {
	log.Info("[GetDocumentByUrl]", url)

	// Request the HTML page.
	res, err := c.httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return nil
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return doc
}

func (c *OlxHttpClient) GetAdvertisements(doc *goquery.Document) []entities.Advertisement {
	var advrtmnts []entities.Advertisement
	if doc == nil {
		return advrtmnts
	}
	// Load the HTML document
	doc.Find("#offers_table").Each(func(i int, s *goquery.Selection) {
		s.Find(".wrap").Each(func(i int, s *goquery.Selection) {
			advrtmnt := entities.Advertisement{}
			title := s.Find("h3")
			url, _ := title.Find("a").Attr("href")
			time := s.Find(".space").Eq(2).Find("p").Last().Text()

			advrtmnt.Title = strings.TrimSpace(title.Text())
			advrtmnt.URL = strings.TrimSpace(url)
			advrtmnt.Time = dateParser.ParseTime(time)

			advrtmnts = append(advrtmnts, advrtmnt)
		})
	})
	return advrtmnts
}

// deep its max page count which will be read
func (c *OlxHttpClient) GetHTMLPages(url string, deep int) []entities.Advertisement {
	log.Info("[GetHTMLPages]", url)

	var advrtmnts []entities.Advertisement

	// Request the HTML page.
	res, err := c.httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	advrtmnts = append(advrtmnts, c.GetAdvertisements(doc)...)

	doc.Find(".pager").Each(func(i int, s *goquery.Selection) {
		if i > deep {
			s.Find(".next").Each(func(i int, s *goquery.Selection) {
				url, _ := s.Find("a").Attr("href")
				advrtmnts = append(advrtmnts, c.GetAdvertisements(c.GetDocumentByUrl(url))...)
			})
		}
	})
	return advrtmnts
}
