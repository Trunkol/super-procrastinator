package consumer

import (
	"super-procrastinator/hacknews"
	"super-procrastinator/medium"
	"super-procrastinator/models"
)

//Stories is responsible for join stories for diferents packages and return to main
func Stories() []models.Article {
	var stories []models.Article

	for _, v := range hacknews.Stories() {
		stories = append(stories, v)
	}

	for _, v := range medium.Stories() {
		stories = append(stories, v)
	}

	return stories
}
