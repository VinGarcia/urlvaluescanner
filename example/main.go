package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/vingarcia/urlvaluescanner"
)

func main() {
	uv := url.Values{
		"name": []string{"some name"},
		"type": []string{"type1", "type2"},
		"age":  []string{"42"},
	}

	var dto struct {
		Name  string   `schema:"name"`
		Types []string `schema:"type,required"`
		Age   int      `schema:"age"`
	}
	err := urlvaluescanner.Unmarshal(uv, &dto)
	if err != nil {
		log.Fatalf("error decoding url values: %s", err)
	}

	fmt.Printf("%+v\n", dto)
}
