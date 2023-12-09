package GOpenAlexAPI

import (
	"fmt"
	"testing"
)

func TestWork(t *testing.T) {
	client := NewClient()

	// fmt.Println(client)

	w := WorkQuery{
		PerPage: 50,
		Select:  []string{"id", "doi", "title", "referenced_works"},
		Filter:  "type:article",
		Search:  "carbon nanotube DNA sorting",
		Sort:    "relevance_score:desc",
	}

	b := client.GetWorks(&w)

	i := 0
	for _, value := range b.Results {
		i += 1
		fmt.Println(i, ". ", value.Title, value.ReferencedWorks)
	}

}
