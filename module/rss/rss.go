package rss

import (
	"encoding/xml"
	"fmt"
	"git.ahrensburg.digital/Thorsten/ahrensburg.blog/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Content string   `xml:"content,attr"`
	Wfw     string   `xml:"wfw,attr"`
	Dc      string   `xml:"dc,attr"`
	Atom    string   `xml:"atom,attr"`
	Sy      string   `xml:"sy,attr"`
	Slash   string   `xml:"slash,attr"`
	Channel struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Description     string `xml:"description"`
		LastBuildDate   string `xml:"lastBuildDate"`
		Language        string `xml:"language"`
		UpdatePeriod    string `xml:"updatePeriod"`
		UpdateFrequency string `xml:"updateFrequency"`
		Generator       string `xml:"generator"`
		Item            []struct {
			Text     string   `xml:",chardata"`
			Title    string   `xml:"title"`
			Link     string   `xml:"link"`
			Creator  string   `xml:"creator"`
			PubDate  string   `xml:"pubDate"`
			Category []string `xml:"category"`
			Guid     struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			Description string `xml:"description"`
			Encoded     string `xml:"encoded"`
		} `xml:"item"`
	} `xml:"channel"`
}

type RssSite struct {
	Inhalt []Rss
	Titel  string
	Urls   string
}
type Rssausgabe struct {
	RssSite     []RssSite
	SeitenTitel string
}

var vorlagen, _ = template.ParseGlob("views/*")

func Start() {
	db, err := gorm.Open(sqlite.Open("Datenbank.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	ausgaben := []models.Rsstabele{}
	db.Find(&ausgaben)

	rssausgabe := []Rssausgabe{}
	for _, ausgabe := range ausgaben {

		rssausgabe1 := Rssausgabe{RssSite: []RssSite{{Urls: ausgabe.Url, Titel: ausgabe.Titel}}}
		rssausgabe = append(rssausgabe, rssausgabe1)
	}

	for _, ausgabe := range rssausgabe {
		ausgabe.RssSite[0].Lesen()
	}
	f, err := os.Create("output/rss.html")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	vorlagen.ExecuteTemplate(f, "rss.html", &rssausgabe)
	f.Close()
	fmt.Println("Seite bearbeitet")

}

func (blog *RssSite) Lesen() *RssSite {

	url := blog.Urls

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body1, _ := ioutil.ReadAll(res.Body)
	xml.Unmarshal((body1), &blog.Inhalt)
	return blog

}
