package main

import (
	"fmt"
	"io/ioutil"
	"newsapietl/apiclient"
	"os"
)

func main() {
	apiKey := os.Getenv("NEWSAPI_API_KEY")
	client := apiclient.MakeNewsApiHTTPClient(apiclient.ApiAuthDetails{
		ApiKey: apiKey,
		ApiUrl: "https://newsapi.org/v2/",
	})
	resp, err := client.MakeRequest("sources")
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Println(string(body), err)
}
