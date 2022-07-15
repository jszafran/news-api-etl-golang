package apiclient

import (
	"net/http"
	"newsapietl/models"
)

type NewsApiHTTPClient struct {
	apiKey string
	apiUrl string
	client http.Client
}

func doRequest(path string) {

}

func (n *NewsApiHTTPClient) GetSources(language string) []models.NewsSource {

}

func (n *NewsApiHTTPClient) GetTopHeadlines(sources []models.NewsSource) []models.SourceTopHeadline {

}

