package apiclient

import (
	"encoding/json"
	"io/ioutil"
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

func MakeNewsApiHTTPClient(apiAuthDetails ApiAuthDetails) *NewsApiHTTPClient {
	client := http.Client{Timeout: time.Second * 5}
	if string(apiAuthDetails.ApiUrl[len(apiAuthDetails.ApiUrl)-1]) != "/" {
		apiAuthDetails.ApiUrl += "/"
	}
	return &NewsApiHTTPClient{apiAuthDetails, client}
}

func (n *NewsApiHTTPClient) MakeGetRequest(path string) (*http.Response, error) {
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

func (n *NewsApiHTTPClient) GetSources(language string) ([]models.NewsSource, error) {
	path := "sources"
	if language != "" {
		path += "?language=" + language + "&pageSize=100"
	}

	resp, err := n.MakeGetRequest(path)
	if err != nil {
		return nil, err
	}

	var sourcesFromApi models.SourceApiResponse
	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err1 := json.Unmarshal(b, &sourcesFromApi)

	if err1 != nil {
		return nil, err1
	}

	sources := make([]models.NewsSource, 0)

	for _, s := range sourcesFromApi.Sources {
		sources = append(sources, models.NewsSource{Id: s.Id, Name: s.Name})
	}
	return sources, nil
}

func (n *NewsApiHTTPClient) GetTopHeadlines(sources []models.NewsSource) ([]models.SourceTopHeadline, error) {
	path := "top-headlines/?sources="
	for _, s := range sources {
		path += s.Id + ","
	}
	// strip trailing coma
	path = path[:len(path)-1]
	path += "&pageSize=100"

	resp, err := n.MakeGetRequest(path)
	if err != nil {
		return nil, err
	}

	var topHeadlinesFromApi models.TopHeadlinesApiResponse

	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err1 := json.Unmarshal(b, &topHeadlinesFromApi)

	if err1 != nil {
		return nil, err1
	}

	ths := make([]models.SourceTopHeadline, 0)

	for _, th := range topHeadlinesFromApi.Articles {
		ths = append(ths, models.SourceTopHeadline{Title: th.Title, SourceId: th.Source.Id})
	}
	return ths, nil
}
