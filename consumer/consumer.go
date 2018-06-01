package consumer

import (
	"super-procrastinator/hacknews"
	"super-procrastinator/medium"
	"super-procrastinator/models"
	"super-procrastinator/reddit"
)

//Stories is responsible for join stories for diferents packages and return to main
func Stories() (stories []models.Article) {

	for _, x := range reddit.Stories() {
		stories = append(stories, x)
	}

	for _, x := range hacknews.Stories() {
		stories = append(stories, x)
	}

	for _, x := range medium.Stories() {
		stories = append(stories, x)
	}

	return stories
}
