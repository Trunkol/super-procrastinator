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

	return convertForInt(body, 15)
}

/*
	convertForInt take a chunk of string like this [12, 34, 12, 45]
	and return the size of itens
*/
func convertForInt(bodyRequest []byte, size int) []int {

	tmp := strings.Split(string(bodyRequest), ",")
	var topFifteen []int

	for _, x := range tmp[:size] {
		x = strings.Trim(x, "[ ")
		x = strings.TrimSpace(x)

		if s, err := strconv.Atoi(x); err == nil {
			topFifteen = append(topFifteen, s)
		}
	}

	return topFifteen
}
