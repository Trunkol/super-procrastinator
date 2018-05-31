package consumer

import (
	"super-procrastinator/hacknews"
	"super-procrastinator/medium"
	"super-procrastinator/models"
	"super-procrastinator/reddit"
)

//Stories is responsible for join stories for diferents packages and return to main
func Stories() []models.Article {
	var stories []models.Article

	articles := make(chan []models.Article)

	go reddit.Stories(articles)
	go hacknews.Stories(articles)
	go medium.Stories(articles)

	for _, x := range <- articles {
		stories = append(stories, x)
	}

	return stories
}
