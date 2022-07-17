package etl

import (
	"log"
	"newsapietl/models"
)

type NewsAPIETL struct {
	ApiClient  models.NewsAPIClient
	DataLoader models.DataLoader
}

func (e *NewsAPIETL) extract() []models.SourceTopHeadline {
	srcs, err := e.ApiClient.GetSources("en")
	if err != nil {
		log.Fatal(err)
	}

	headlines, err1 := e.ApiClient.GetTopHeadlines(srcs)
	if err1 != nil {
		log.Fatal(err1)
	}

	return headlines
}

func (e *NewsAPIETL) transform(headlines []models.SourceTopHeadline) []models.SourceAggregatedTopHeadlines {
	m := make(map[string][]models.SourceTopHeadline)
	for _, h := range headlines {
		th := models.SourceTopHeadline{
			Title:    h.Title,
			SourceId: h.SourceId,
		}
		m[h.SourceId] = append(m[h.SourceId], th)
	}
	aggTopHeadlines := make([]models.SourceAggregatedTopHeadlines, 0)

	for k, v := range m {
		aggTopHeadline := models.SourceAggregatedTopHeadlines{
			SourceId:     k,
			TopHeadlines: v,
		}
	}
}
