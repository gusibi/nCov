package api

type rssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Author      string `xml:"author"`
}

type rss struct {
	Version     string `xml:"version,attr"`
	Description string `xml:"channel>description"`
	Link        string `xml:"channel>link"`
	Title       string `xml:"channel>title"`

	Item []rssItem `xml:"channel>item"`
}
