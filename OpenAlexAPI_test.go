package GOpenAlexAPI

import (
	"fmt"
	"testing"
)

func TestWork(t *testing.T) {
	client := NewClient()

	// fmt.Println(client)

	w := &Query{
		ID: "https://doi.org/10.3168/jds.s0022-0302(47)92414-4",
	}

	// json_test := GetResponse("https://api.openalex.org/works/https://doi.org/10.3168/jds.s0022-0302(47)92414-4")

	b := client.GetWork(w)

	fmt.Println(b.CitedByPercentileYear)

}
