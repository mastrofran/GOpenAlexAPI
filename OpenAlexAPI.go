package GOpenAlexAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const base_endpoint = "https://api.openalex.org/"

func NewClient() *Client {
	return &Client{BaseURL: base_endpoint}
}

func (c *Client) GetWork(work *Query) *Work {
	endpoint := base_endpoint + "works/" + work.ID

	if work.Select != nil {
		endpoint += "?select="
		for idx, kw := range work.Select {
			endpoint += kw
			if idx != len(work.Select)-1 {
				endpoint += ","
			}
		}
	}

	var r Work

	// fmt.Println(GetResponse(endpoint))
	get_response := GetResponse(endpoint)
	if err := json.Unmarshal(get_response, &r); err != nil {
		log.Fatal(err)
	}

	// b, _ := json.Marshal(r)
	return &r

}

func GetResponse(endpoint string) []byte {
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var b []byte

	if resp.StatusCode == http.StatusOK {
		b, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Printf("error reading http request: %s\n", err)
			log.Fatal(err)
		}
		return b
	}
	return b
}
