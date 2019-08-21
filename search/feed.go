package search

import (
    "encoding/json"
    "os"
)

const dataFile = "data/data.json"

// Feed contains informaton we need to process a feed 
type Feed struct {
    Name string `json:"site"`
    URI string `json:"link"`
    Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
    file, err := os.Open(dataFile)
    if err != nil {
        return nil, err
    }

    // Schedule the file to be closed o
    // once the function returns
    defer file.Close()

    // Decode the file into a slice of pointers to Feed values
    var feeds []*Feed
    err = json.NewDecoder(file).Decode(&feeds)

    // We dont need to check for errors, the caller can do this.
    return feeds, err
}