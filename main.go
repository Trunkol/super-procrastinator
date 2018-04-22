package main

import "fmt"

func main() {
	topStories := topStories()

	for _, i := range topStories {
		fmt.Println(getStorie(i))
	}
}
