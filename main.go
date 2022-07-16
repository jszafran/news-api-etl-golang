package main

import (
	"fmt"
	"log"
	"newsapietl/apiclient"
	"os"
)

func main() {
	apiKey := os.Getenv("NEWSAPI_API_KEY")
	client := apiclient.MakeNewsApiHTTPClient(apiclient.ApiAuthDetails{
		ApiKey: apiKey,
		ApiUrl: "https://newsapi.org/v2/",
	})
	sources, err := client.GetSources("en")

	if err != nil {
		log.Fatal(err)
	}
	for _, s := range sources {
		fmt.Println(s)
	}
}
