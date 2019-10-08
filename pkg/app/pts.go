package app

import (
  "log"
  "net/http"
	"strings"
  "github.com/PuerkitoBio/goquery"
	"github.com/rivo/tview"
	"github.com/Zanets/tclient/pkg/ui"
	"fmt"
)

const url_news = "https://news.pts.org.tw/dailynews.php"
const news_list_item = ".mid-news"
const news_title_item = ".m-left-side"
const news_content_item = ".m-right-side"
type APP_pts struct {
	ui.APP_UI
	content_ui *tview.Grid
}

func NewAPP_PTS() *APP_pts {
	ins := APP_pts{}
	ins.content_ui = tview.NewGrid()
	return &ins
}

func (this APP_pts) GetName() string {
	return "PTS News"
}

func (this APP_pts) Focus() {
	
}

func (this APP_pts) GetContent() (*tview.Grid, error) {
	var content strings.Builder
  // Request the HTML page.
  res, err := http.Get(url_news)
  if err != nil {
		return nil, err;
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
		return nil, fmt.Errorf("Get %s fail with status code: %d, status: %s\n", url_news, res.StatusCode, res.Status) 
  }

  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
		return nil, err
  }

	news_list := doc.Find(news_list_item).First()
	if news_list == nil {
		return nil, fmt.Errorf("Find news list fail. Using %s.", news_list_item) 
	}

	titles := news_list.Find(news_title_item).First()
	if titles == nil {
		log.Fatalf("cannot find titles")
	}

	title := titles.Find("span").First().Text() 

	content.WriteString(title)
	content.WriteString("\n")

	contents := news_list.Find(news_content_item).First()
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

	main_ui := tview.NewTextView()
	main_ui.SetText(content.String())

	this.content_ui.AddItem(main_ui, 0, 1, 1, 1, 0, 0, false)

	return this.content_ui, nil
}