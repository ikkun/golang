package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://outlook.office.com/webhook/6683c0f0-1a59-47e1-8f84-687ea738217d@6f9d6451-ab20-4353-b585-7f7b9836eda4/IncomingWebhook/6268443d56644d9f8aa17d2c988be410/7fcb6ab6-ed1d-4cd6-9ae8-3731cca9c1b4"
	var jsonStr = []byte(`{"text": "![TestImage](https://teamsnodesample.azurewebsites.net/static/img/image5.png) mBanking Alert"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
