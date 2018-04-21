package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	urlBase = "https://hacker-news.firebaseio.com/v0/"
)

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

	tmp := strings.Split(string(body), ",")
	var topFifteen []int

	for _, x := range tmp[:15] {
		x = strings.Trim(x, "[ ")
		x = strings.TrimSpace(x)

		if s, err := strconv.Atoi(x); err == nil {
			topFifteen = append(topFifteen, s)
		}
	}

	return topFifteen
}
