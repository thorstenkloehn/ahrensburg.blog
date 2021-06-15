package feed

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Atom    string   `xml:"atom,attr"`
	Channel struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Description string `xml:"description"`
		Language    string `xml:"language"`
		Copyright   string `xml:"copyright"`
		PubDate     string `xml:"pubDate"`
		Item        []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Description string `xml:"description"`
			Link        string `xml:"link"`
			Guid        struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			PubDate string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}

func (rss *Rss) Lesen(urlEingeben string) *Rss {

	ergennis, error := http.Get(urlEingeben)
	if error != nil {
		log.Fatal(error)
	}
	robot, error1 := io.ReadAll(ergennis.Body)
	if error1 != nil {
		log.Fatal(error)
	}
	err1 := xml.Unmarshal(robot, &rss)
	if err1 != nil {
		log.Fatal(error)

	}

	return rss
}
