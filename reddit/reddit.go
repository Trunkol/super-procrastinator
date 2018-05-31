package reddit

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type result struct {
	Articles []tmpArticle `xml:"entry"`
}

type tmpArticle struct {
	Author  string `xml:"author>name"`
	Title   string `xml:"title"`
	Updated string `xml:"updated"`
}

func TopStories() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.reddit.com/r/inthenews/.rss", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var r result
	err = xml.Unmarshal(body, &r)
	for _, i := range r.Articles {
		fmt.Println(i.Title)
	}

}
