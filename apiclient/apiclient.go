package apiclient

import (
	"net/http"
	"newsapietl/models"
	"time"
)

type ApiAuthDetails struct {
	apiKey string
	apiUrl string
}

type NewsApiHTTPClient struct {
	apiAuthDetails ApiAuthDetails
	client         http.Client
}

func MakeNewsApiHTTPClient(apiAuthDetails ApiAuthDetails) NewsApiHTTPClient {
	client := http.Client{Timeout: time.Second * 5}
	if string(apiAuthDetails.apiUrl[len(apiAuthDetails.apiUrl)-1]) != "/" {
		apiAuthDetails.apiUrl += "/"
	}
	return NewsApiHTTPClient{apiAuthDetails, client}
}

func (n *NewsApiHTTPClient) doRequest(path string) http.Response {
	return http.Response{}
}

func (n *NewsApiHTTPClient) GetSources(language string) []models.NewsSource {
	return []models.NewsSource{}
}

func (n *NewsApiHTTPClient) GetTopHeadlines(sources []models.NewsSource) []models.SourceTopHeadline {
	return []models.SourceTopHeadline{}
}
