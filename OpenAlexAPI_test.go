package OpenAlexAPI

import (
	"fmt"
	"testing"
)

func TestWork(t *testing.T) {
	client := NewClient()

	fmt.Println(client)

	w := &Work{
		ID:     "https://doi.org/10.3168/jds.s0022-0302(47)92414-4",
		Select: []string{"id", "doi"},
	}

	client.GetWork(w)
}
