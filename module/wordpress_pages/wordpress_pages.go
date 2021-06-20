package wordpress_pages

type Bloginhalt struct {
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

func Start() {

}
