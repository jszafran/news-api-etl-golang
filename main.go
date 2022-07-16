package main

import (
	"fmt"
	"newsapietl/apiclient"
	"newsapietl/models"
	"os"
)

func main() {
	apiKey := os.Getenv("NEWSAPI_API_KEY")
	httpClient := apiclient.MakeNewsApiHTTPClient(apiclient.ApiAuthDetails{
		ApiKey: apiKey,
		ApiUrl: "https://newsapi.org/v2/",
	})

	memClient := &apiclient.InMemoryNewsAPIClient{
		Sources: []models.NewsSource{{Id: "some-id", Name: "Some Name"}, {Id: "other-id", Name: "Other Name"}},
		TopHeadlines: []models.SourceTopHeadline{
			{SourceId: "some-id", Title: "Title 1"},
			{SourceId: "other-id", Title: "Title 2"},
		},
	}

	printSources := func(c models.NewsAPIClient) {
		sources, _ := c.GetSources("en")
		for _, source := range sources {
			fmt.Println(source.Id)
		}
	}

	clients := []models.NewsAPIClient{memClient, &httpClient}

	for _, c := range clients {
		printSources(c)
	}
}
