package reddit

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"super-procrastinator/models"
)

type result struct {
	Articles []tmpArticle `xml:"entry"`
}

type link struct {
	Link string `xml:"href,attr"`
}

type tmpArticle struct {
	Author  string `xml:"author>name"`
	Title   string `xml:"title"`
	Updated string `xml:"updated"`
	Link    link   `xml:"link"`
}

//Stories is...
func Stories(ch chan []models.Article) {
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

	ch <- joinArticles(r.Articles)
}

func joinArticles(arts []tmpArticle) (articles []models.Article) {
	for _, x := range arts {
		articles = append(articles, models.Article{
			URL:    x.Link.Link,
			Title:  x.Title,
			Author: x.Author[3:],
			Date:   123,
			Source: "Reddit",
		})
	}

	return articles
}
