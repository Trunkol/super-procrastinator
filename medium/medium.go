package medium

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type feedReturn struct {
	Payload struct {
		Collection struct {
			Slug string `json:"slug"`
		}
		Posts []struct {
			Title string `json:"title"`
			Date  int    `json:"updatedAt"`
			URL   string `json:"uniqueSlug"`
		}
	}
}

type storie struct {
	title  string
	url    string
	author string
	date   int
}

//Stories is a func
func Stories() (stories []storie) {
	feeds := []string{"message", "the-launchism"}

	for _, content := range feeds {
		dat := feedReturn{}

		err := json.Unmarshal(getFeed(content), &dat)
		if err != nil {
			log.Fatal(err)
		}

		for _, tmp := range dat.Payload.Posts {
			stories = append(stories, storie{
				url:    "https://medium.com/" + dat.Payload.Collection.Slug + "/" + tmp.URL,
				title:  tmp.Title,
				author: "",
				date:   tmp.Date,
			})
		}
	}

	return stories
}

func getFeed(feed string) []byte {
	req, err := http.Get("https://medium.com/" + feed + "/latest?format=json")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	return []byte(string(body)[16:])
}
