package models

type NewsSource struct {
	Id   string
	Name string
}

type TopHeadline struct {
	Title string
}

type SourceTopHeadline struct {
	Title    string
	SourceId string
}

type SourceAggregatedTopHeadlines struct {
	SourceId     string
	TopHeadlines []TopHeadline
}

type NewsAPIClient interface {
	GetSources(language string) []NewsSource
	GetTopHeadlines(sources []NewsSource) []SourceTopHeadline
}

type DataLoader interface {
	LoadHeadlines(sourceTopHeadlines []SourceAggregatedTopHeadlines, timestamp string)
}
