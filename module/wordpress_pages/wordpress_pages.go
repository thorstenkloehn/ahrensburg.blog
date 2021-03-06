package wordpress_pages

import (
	"encoding/json"
	"fmt"
	"github.com/thorstenkloehn/ahrensburg.blog/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

type BlogInhalt struct {
	Id      int    `json:"id"`
	Date    string `json:"date"`
	DateGmt string `json:"date_gmt"`
	Guid    struct {
		Rendered string `json:"rendered"`
	} `json:"guid"`
	Modified    string `json:"modified"`
	ModifiedGmt string `json:"modified_gmt"`
	Slug        string `json:"slug"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Link        string `json:"link"`
	Title       struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
	Content struct {
		Rendered  string `json:"rendered"`
		Protected bool   `json:"protected"`
	} `json:"content"`
	Excerpt struct {
		Rendered  string `json:"rendered"`
		Protected bool   `json:"protected"`
	} `json:"excerpt"`
	Author        int           `json:"author"`
	FeaturedMedia int           `json:"featured_media"`
	CommentStatus string        `json:"comment_status"`
	PingStatus    string        `json:"ping_status"`
	Sticky        bool          `json:"sticky"`
	Template      string        `json:"template"`
	Format        string        `json:"format"`
	Meta          []interface{} `json:"meta"`
	Categories    []int         `json:"categories"`
	Tags          []int         `json:"tags"`
	YoastHead     string        `json:"yoast_head"`
	Links         struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
		About []struct {
			Href string `json:"href"`
		} `json:"about"`
		Author []struct {
			Embeddable bool   `json:"embeddable"`
			Href       string `json:"href"`
		} `json:"author"`
		Replies []struct {
			Embeddable bool   `json:"embeddable"`
			Href       string `json:"href"`
		} `json:"replies"`
		VersionHistory []struct {
			Count int    `json:"count"`
			Href  string `json:"href"`
		} `json:"version-history"`
		PredecessorVersion []struct {
			Id   int    `json:"id"`
			Href string `json:"href"`
		} `json:"predecessor-version"`
		WpFeaturedmedia []struct {
			Embeddable bool   `json:"embeddable"`
			Href       string `json:"href"`
		} `json:"wp:featuredmedia"`
		WpAttachment []struct {
			Href string `json:"href"`
		} `json:"wp:attachment"`
		WpTerm []struct {
			Taxonomy   string `json:"taxonomy"`
			Embeddable bool   `json:"embeddable"`
			Href       string `json:"href"`
		} `json:"wp:term"`
		Curies []struct {
			Name      string `json:"name"`
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
	} `json:"_links"`
}

type BlogSite struct {
	Inhalt []BlogInhalt
	Titel  string
	Urls   string
}

type Blogausgabe struct {
	BlogSite    []BlogSite
	SeitenTitel string
}

func Start() {
	db, err := gorm.Open(sqlite.Open("Datenbank.db"), &gorm.Config{})
	if err != nil {

		log.Fatal(err)
	}

	ausgaben := []models.Wordpress{}
	db.Find(&ausgaben)

	var vorlagen, _ = template.ParseGlob("views/*")

	blogausgabe := []Blogausgabe{}
	for _, ausgabe := range ausgaben {
		blogausgabe1 := Blogausgabe{BlogSite: []BlogSite{{Urls: ausgabe.Url, Titel: ausgabe.Titel}}}
		blogausgabe = append(blogausgabe, blogausgabe1)
	}
	for _, ausgabe := range blogausgabe {
		ausgabe.BlogSite[0].Lesen()
	}
	f, err := os.Create("output/index.html")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	vorlagen.ExecuteTemplate(f, "wordpresspagesSeite.html", &blogausgabe)
	f.Close()
	fmt.Println("Seite bearbeitet")

}

func (blog *BlogSite) Lesen() *BlogSite {

	url := blog.Urls + "/wp-json/wp/v2/posts"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body1, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal((body1), &blog.Inhalt)
	return blog

}
