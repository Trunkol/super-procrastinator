package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	urlBase = "https://hacker-news.firebaseio.com/v0/"
)

func getStorie(id int) interface{} {
	r, err := http.Get(urlBase + "item/" + strconv.Itoa(id) + ".json?print=pretty")

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var dat map[string]interface{}

	err = json.Unmarshal(body, &dat)

	if err != nil {
		log.Fatal(err)
	}

	return dat
	//https://hacker-news.firebaseio.com/v0/item/16892211.json?print=pretty
}

func topStories() []int {
	resp, err := http.Get(urlBase + "topstories.json?print=pretty")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var topFifteen []int

	err = json.Unmarshal(body, &topFifteen)

	if err != nil {
		log.Fatal(err)
	}

	return topFifteen[:15]
}
