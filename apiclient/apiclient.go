package apiclient

import (
	"log"
	"net/http"
	"newsapietl/models"
	"time"
)

type ApiAuthDetails struct {
	ApiKey string
	ApiUrl string
}

type NewsApiHTTPClient struct {
	apiAuthDetails ApiAuthDetails
	client         http.Client
}

func MakeNewsApiHTTPClient(apiAuthDetails ApiAuthDetails) NewsApiHTTPClient {
	client := http.Client{Timeout: time.Second * 5}
	if string(apiAuthDetails.ApiUrl[len(apiAuthDetails.ApiUrl)-1]) != "/" {
		apiAuthDetails.ApiUrl += "/"
	}
	return NewsApiHTTPClient{apiAuthDetails, client}
}

func (n *NewsApiHTTPClient) MakeRequest(path string) (*http.Response, error) {
	if string(path[0]) == "/" {
		path = path[1:]
	}
	url := n.apiAuthDetails.ApiUrl + path
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Failed to create requqest: %v", err)
	}
	req.Header.Set("X-Api-Key", n.apiAuthDetails.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err1 := n.client.Do(req)
	return resp, err1
}

func (n *NewsApiHTTPClient) GetSources(language string) []models.NewsSource {
	return []models.NewsSource{}
}

func (n *NewsApiHTTPClient) GetTopHeadlines(sources []models.NewsSource) []models.SourceTopHeadline {
	return []models.SourceTopHeadline{}
}
