package medium

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type contentResponse struct {
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

//Stories is a func
func Stories() {
	feeds := []string{"message", "the-launchism"}

	for _, content := range feeds {
		dat := contentResponse{}

		err := json.Unmarshal(getFeed(content), &dat)
		if err != nil {
			log.Fatal(err)
		}

	}
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
