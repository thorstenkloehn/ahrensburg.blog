package main

import (
	"encoding/xml"
	"fmt"
	"github.com/thorstenkloehn/ahrensburg.digital/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
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

type Rsstabele struct {
	ID  uint `gorm:"primaryKey"`
	Url string
}

type Rsswebsiten struct {
	Url        []string
	Rsswebsite []Rss
}

var vorlagen, _ = template.ParseGlob("views/*")

func main() {
	db, err := gorm.Open(sqlite.Open("Datenbank.db"), &gorm.Config{})
	if err != nil {

		log.Fatal(err)
	}

	ausgaben := []models.Rsstabele{}
	db.Find(&ausgaben)

	rss := []Rsswebsiten{}
	rssausgabe := Rss{}

	for _, rssurlausgabe := range ausgaben {
		resp, _ := http.Get(rssurlausgabe.Url)
		document, _ := io.ReadAll(resp.Body)
		xml.Unmarshal([]byte(document), &rssausgabe)
		start := Rsswebsiten{Url: []string{rssurlausgabe.Url}, Rsswebsite: []Rss{rssausgabe}}
		rss = append(rss, start)
	}
	f, err := os.Create("output/rss.html")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(rss[0].Rsswebsite[0].Channel.Item[0].Title)
	vorlagen.ExecuteTemplate(f, "rss.html", &rss)
	f.Close()
	fmt.Println("Rss Seite ist herstellt")

}
