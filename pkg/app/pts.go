package app

import (
  "log"
  "net/http"
	"strings"
  "github.com/PuerkitoBio/goquery"
	"github.com/Zanets/tclient/pkg/ui"
)

const url_news = "https://news.pts.org.tw/dailynews.php"

type APP_pts struct {
	ui.APP_UI
}

func (this APP_pts) GetName() string {
	return "PTS News"
}

func (this APP_pts) GetContent() string {
	var content strings.Builder
  // Request the HTML page.
  res, err := http.Get(url_news)
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

	news_list := doc.Find(".mid-news").First()
	if news_list == nil {
		log.Fatalf("cannot find news_list")
	}

	titles := news_list.Find(".m-left-side").First()
	if titles == nil {
		log.Fatalf("cannot find titles")
	}

	title := titles.Find("span").First().Text() 

	content.WriteString(title)
	content.WriteString("\n")

	contents := news_list.Find(".m-right-side").First()
	if contents == nil {
		log.Fatalf("cannot find contents")
	}

	contents.Find(".news-list").Each(func (i int, s *goquery.Selection) {
		news_title := s.Find(".text-title").Find("span").First().Text()
		news_content := s.Find(".text-content").First().Text()
		content.WriteString("[")
		content.WriteString(news_title)
		content.WriteString("]\n")
		
		content.WriteString(news_content)
		content.WriteString("\n")
	})

	return content.String()
}