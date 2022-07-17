package loaders

import (
	"encoding/csv"
	"newsapietl/models"
	"os"
	"path/filepath"
	"strings"
)

type LocalDiskCSVLoader struct {
	Path string
}

func saveHeadlinesToCSV(headlines []models.TopHeadline, csvPath string) error {
	p, err := os.Create(csvPath)
	defer p.Close()
	if err != nil {
		return err
	}
	writer := csv.NewWriter(p)
	defer writer.Flush()

	records := make([][]string, len(headlines))

	for _, h := range headlines {
		records = append(records, []string{h.Title})
	}
	err1 := writer.WriteAll(records)
	if err1 != nil {
		return err1
	}
	return nil
}

func (l *LocalDiskCSVLoader) LoadHeadlines(
	sourceTopHeadlines []models.SourceAggregatedTopHeadlines,
	timestamp string) error {
	if _, err := os.Stat(l.Path); os.IsNotExist(err) {
		_ = os.Mkdir(l.Path, os.ModePerm)
	}
	timestamp = strings.Replace(timestamp[:19], ":", "", -1)

	for _, h := range sourceTopHeadlines {
		sourceDir := filepath.Join(l.Path, h.SourceId)
		if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
			_ = os.Mkdir(sourceDir, os.ModePerm)
		}
		csvPath := filepath.Join(sourceDir, timestamp+"_headlines.csv")
		err1 := saveHeadlinesToCSV(h.TopHeadlines, csvPath)
		if err1 != nil {
			return err1
		}
	}
	return nil
}
