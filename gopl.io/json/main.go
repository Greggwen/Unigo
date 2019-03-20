package main

import (
	"fmt"
	"encoding/json"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color, omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{
			Title: "Casablance",
			Year: 1942,
			Color: false,
			Actors: []string {"aaaa", "bbbb", "ccc",},
		},
		{
			Title: "Cool Hand Luke",
			Year: 1967,
			Color: false,
			Actors: []string {"aaaa", "bbbb", "ccc",},
		},
		{
			Title: "Bullitt",
			Year: 1968,
			Color: false,
			Actors: []string {"aaaa", "bbbb", "ccc",},
		},
	}

	//data, _ := json.Marshal(movies)

	data, _ := json.MarshalIndent(movies, "", "    ")
	fmt.Printf("%s\n", data)
}
