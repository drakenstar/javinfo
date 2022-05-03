package javinfo

import (
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type JAVLibraryBackend struct{}

func (b *JAVLibraryBackend) FindByCode(code Code) ([]*Title, error) {
	endpoint, _ := url.Parse("https://www.javlibrary.com/en/vl_searchbyid.php")
	query := url.Values{}
	query["keyword"] = []string{string(code)}
	endpoint.RawQuery = query.Encode()

	titles := []*Title{}

	r, err := http.Get(endpoint.String())
	if err != nil {
		return nil, err
	}

	t, err := parseTitleFromHTML(r)
	if err != nil {
		return nil, err
	}

	titles = append(titles, t)
	return titles, nil
}

func parseTitleFromHTML(r *http.Response) (*Title, error) {
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	t := &Title{
		ID:    r.Request.URL.Query().Get("v"),
		Code:  Code(doc.Find("#video_id .text").Text()),
		Title: doc.Find("#video_title").Text(), // TODO Trim code from the start of the title
	}

	if releaseDate, err := time.Parse("", doc.Find("#video_date .text").Text()); err == nil {
		t.ReleaseDate = releaseDate
	}

	doc.Find("#video_cast .cast").Each(func(i int, s *goquery.Selection) {
		pathString, _ := s.Find("a").Attr("href")
		path, _ := url.Parse(pathString)
		// TODO Error handling for the above cases
		t.Models = append(t.Models, &Model{
			ID:   path.Query().Get("s"),
			Name: s.Text(),
		})
	})

	doc.Find("#video_genres .genre").Each(func(i int, s *goquery.Selection) {
		pathString, _ := s.Find("a").Attr("href")
		path, _ := url.Parse(pathString)
		// TODO Error handling for the above cases
		t.Tags = append(t.Tags, &Tag{
			ID:   path.Query().Get("g"),
			Name: s.Text(),
		})
	})

	return t, nil
}
