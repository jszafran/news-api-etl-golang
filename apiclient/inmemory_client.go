package apiclient

import "newsapietl/models"

type InMemoryNewsAPIClient struct {
	Sources      []models.NewsSource
	TopHeadlines []models.SourceTopHeadline
}

func (c *InMemoryNewsAPIClient) GetSources(language string) ([]models.NewsSource, error) {
	return c.Sources, nil
}

func (c *InMemoryNewsAPIClient) GetTopHeadlines(sources []models.NewsSource) ([]models.SourceTopHeadline, error) {
	return c.TopHeadlines, nil
}
