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

	fmt.Println("Sources:")
	for _, s := range sources {
		fmt.Println(s)
	}

	ths, err1 := client.GetTopHeadlines(sources)

	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("Top headlines:")
	for _, th := range ths {
		fmt.Println(th)
	}
}
