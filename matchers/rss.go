package matchers

import (
    "encoding/xml"
    "errors"
    "fmt"
    "log"
    "net/http"
    "regexp"

    "github.com/derekmt/goinaction/search"
)

type (
    // Item defines the fields associated with the item tag 
    // in the rss document
    item struct {
        XMLName xml.Name `xml:"item"`
        PubDate string `xml:"pubDate"`
        Title string `xml:"title"`
        Description string `xml:"description"`
        Link string `xml:"link"`
        GUID string `xml:"guid"`
        GeoRssPoint string `xml:"georss:point"`
    }

    // image defines the fields associated with the image tag 
    // in the RSS doc
    image struct {
        XMLName xml.Name `xml:"image"`
        URL string `xml:"url"`
        Title string `xml:"title"`
        Link string `xml:"link"`
    }

    // channel defines the fields associated with the channel tag 
    // in the rss doc 
    channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

    // rssDocument defines the fields associated with the RSS doc
    rssDocument struct {
        XMLName string `xml:"rss"`
        Channel channel `xml:"channel"`
    }
)

// rssMatcher implements the Matcher interface
type rssMatcher struct{}

// init registers the matcher with the program
func init() {
    var matcher rssMatcher
    search.Register("rss", matcher)
}