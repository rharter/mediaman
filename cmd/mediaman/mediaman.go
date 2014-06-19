package main

import (
	"fmt"
	"github.com/rharter/go-tvdb"
)

func main() {

	seriesList, err := tvdb.GetSeriesWithLang("Arrow", "en")
	if (err != nil) {
		fmt.Printf("Error finding series: %v", err)
	}
	fmt.Printf("Found %d series\n", len(seriesList.Series))
	for _, show := range seriesList.Series {
		fmt.Printf("%+v\n", show)
	}
}