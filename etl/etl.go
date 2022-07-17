package etl

import (
	"log"
	"newsapietl/models"
	"time"
)

func extract(apiClient models.NewsAPIClient) []models.SourceTopHeadline {
	sources, err := apiClient.GetSources("en")
	if err != nil {
		log.Fatal(err)
	}

	headlines, err1 := apiClient.GetTopHeadlines(sources)
	if err1 != nil {
		log.Fatal(err1)
	}

	return headlines
}

func transform(headlines []models.SourceTopHeadline) []models.SourceAggregatedTopHeadlines {
	m := make(map[string][]models.TopHeadline)
	for _, h := range headlines {
		th := models.TopHeadline{Title: h.Title}
		m[h.SourceId] = append(m[h.SourceId], th)
	}
	aggTopHeadlines := make([]models.SourceAggregatedTopHeadlines, 0)

	for k, v := range m {
		aggTopHeadline := models.SourceAggregatedTopHeadlines{
			SourceId:     k,
			TopHeadlines: v,
		}
		aggTopHeadlines = append(aggTopHeadlines, aggTopHeadline)
	}
	return aggTopHeadlines
}

func RunETL(apiClient models.NewsAPIClient, loader models.DataLoader) error {
	runTimestamp := time.Now().UTC().Format(time.RFC3339)
	data := extract(apiClient)
	dataTransformed := transform(data)
	err := loader.LoadHeadlines(dataTransformed, runTimestamp)
	return err
}
