package medium

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Stories is a func
func Stories() {
	feeds := []string{"message", "the-story"}

	for _, content := range feeds {
		body := getFeed(content)

		var dat map[string]interface{}

		err := json.Unmarshal(body, &dat)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(dat)
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
