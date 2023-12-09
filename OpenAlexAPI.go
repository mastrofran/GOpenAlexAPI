package GOpenAlexAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const base_endpoint = "https://api.openalex.org/"

func NewClient() *Client {
	return &Client{BaseURL: base_endpoint}
}

func (c *Client) GetWork(work *WorkQuery) *Work {
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

func (c *Client) GetWorks(works *WorkQuery) *Works {
	endpoint := base_endpoint + "works?"

	if works.PageNum != 0 {
		endpoint += "page=" + strconv.Itoa(works.PageNum) + "&"
	}

	if works.PerPage != 0 {
		endpoint += "per-page=" + strconv.Itoa(works.PerPage) + "&"
	}

	if works.Sample != 0 && works.Seed != 0 {
		endpoint += "sample=" + strconv.Itoa(works.Sample) + "&" + "seed=" + strconv.Itoa(works.Seed) + "&"
	}

	if works.Select != nil {
		endpoint += "select="
		for idx, kw := range works.Select {
			endpoint += kw
			if idx != len(works.Select)-1 {
				endpoint += ","
			}
		}
		endpoint += "&"
	}

	if works.Filter != "" {
		filter := strings.Replace(works.Filter, " ", "%20", -1)
		endpoint += "filter=" + filter + "&"
	}

	if works.Search != "" {
		search := strings.Replace(works.Search, " ", "%20", -1)
		endpoint += "search=" + search + "&"
	}

	if works.Sort != "" {
		endpoint += "sort=" + works.Sort + "&"
	}

	var r Works

	// fmt.Println(endpoint)

	get_response := GetResponse(endpoint)
	if err := json.Unmarshal(get_response, &r); err != nil {
		log.Fatal()
	}

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
