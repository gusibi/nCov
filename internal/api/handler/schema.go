package handler

type rssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link,omitempty"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	//Author      string `xml:"author,omitempty"`
}

type rss struct {
	Version     string `xml:"version,attr"`
	Description string `xml:"channel>description"`
	Link        string `xml:"channel>link"`
	Title       string `xml:"channel>title"`

	Item []rssItem `xml:"channel>item"`
}
