package main

import (
	"log"
	"newsapietl/apiclient"
	"newsapietl/etl"
	"newsapietl/loaders"
	"os"
)

func main() {
	apiKey := os.Getenv("NEWSAPI_API_KEY")
	httpClient, err := apiclient.MakeNewsApiHTTPClient(apiclient.ApiAuthDetails{
		ApiKey: apiKey,
		ApiUrl: "https://newsapi.org/v2/",
	})

	if err != nil {
		log.Fatal(err)
	}

	loader := loaders.LocalDiskCSVLoader{Path: "results"}
	err = etl.RunETL(httpClient, &loader)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("ETL finished!")

}
