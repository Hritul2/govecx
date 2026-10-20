// Command basic shows the smallest useful govecx program.
package main

import (
	"fmt"
	"log"

	"github.com/Hritul2/govecx"
)

func main() {
	idx, err := govecx.New(
		govecx.WithDimension(3),
		govecx.WithMetric(govecx.Cosine),
	)
	if err != nil {
		log.Fatal(err)
	}

	docs := map[string][]float32{
		"cat":   {0.9, 0.1, 0.0},
		"dog":   {0.8, 0.2, 0.1},
		"plane": {0.0, 0.1, 0.9},
	}
	for id, v := range docs {
		if err := idx.Add(id, v); err != nil {
			log.Fatal(err)
		}
	}

	results, err := idx.Search([]float32{1, 0, 0}, 2)
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range results {
		fmt.Printf("%s\t%.4f\n", r.ID, r.Score)
	}
}
