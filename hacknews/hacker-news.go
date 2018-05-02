package hacknews

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"super-procrastinator/models"
	"sync"
)

const (
	urlBase = "https://hacker-news.firebaseio.com/v0/"
)

//Stories is the main function to catch stories
func Stories() []models.Article {
	topID := topStories(15)

	var stories []models.Article
	var wg sync.WaitGroup

	for _, x := range topID {
		wg.Add(1)
		go func(i int) {
			stories = append(stories, getStorie(i))
			wg.Done()
		}(x)
	}
	wg.Wait()

	return stories
}

//getStorie is responsible for take the content of a storie
func getStorie(id int) models.Article {
	r, err := http.Get(urlBase + "item/" + strconv.Itoa(id) + ".json?print=pretty")

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var dat models.Article

	err = json.Unmarshal(body, &dat)

	if err != nil {
		log.Fatal(err)
	}

	return dat
}

//topStories return the top N of stories
func topStories(numStories int) []int {
	resp, err := http.Get(urlBase + "topstories.json?print=pretty")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var topArticles []int

	err = json.Unmarshal(body, &topArticles)

	if err != nil {
		log.Fatal(err)
	}

	return topArticles[:numStories]
}
