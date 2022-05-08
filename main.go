package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//unmarshal xml

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}

type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

func run() error {
	var r RSS
	data := readGoogleTrends()

	//unmarshal data(rss xml)
	err := xml.Unmarshal(data, &r)
	if err != nil {
		return fmt.Errorf("error unmarshalling: %w", err)
	}

	fmt.Println("\n Below are all the Google Search Trends For Today!")
	fmt.Println("----------------------------------------------------")

	for i := range r.Channel.ItemList {
		rank := (i + 1)
		fmt.Println("#", rank)
		fmt.Println("Search Term: ", r.Channel.ItemList[i].Title)
		fmt.Println("Link to the Trend: ", r.Channel.ItemList[i].Link)
		for n := range r.Channel.ItemList[i].NewsItems {
			fmt.Println("Heading: ", r.Channel.ItemList[i].NewsItems[n].Headline)
			fmt.Println("Link to article:", r.Channel.ItemList[i].NewsItems[n].HeadlineLink)
		}
		fmt.Println("------------------------------------------------------------------")
	}
	fmt.Println(err)
	return err

}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

}

func getGoogleTrends() *http.Response {
	response, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US")
	if err != nil {
		log.Fatal(err)
	}
	return response
}

func readGoogleTrends() []byte {
	response := getGoogleTrends()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData

}
