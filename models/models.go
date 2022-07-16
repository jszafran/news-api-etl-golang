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

type SourceApiResponse struct {
	Status  string `json:"status"`
	Sources []struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Url         string `json:"url"`
		Category    string `json:"category"`
		Language    string `json:"language"`
		Country     string `json:"country"`
	} `json:"sources"`
}

type TopHeadlinesApiResponse struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"source"`
		Author      string `json:"author"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"articles"`
}

type NewsAPIClient interface {
	GetSources(language string) ([]NewsSource, error)
	GetTopHeadlines(sources []NewsSource) ([]SourceTopHeadline, error)
}

type DataLoader interface {
	LoadHeadlines(sourceTopHeadlines []SourceAggregatedTopHeadlines, timestamp string)
}
