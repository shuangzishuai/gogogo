package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	// fmt.Printf("%#v\n", movies)

	//转为json
	// b, err := json.Marshal(movies)
	b, err := json.MarshalIndent(movies, "", "    ") //该函数有两个额外的字符串参数用于表示每一行输出的前缀和每一个层级的缩进：
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", b)
	/*
		[
			{
				"Title": "Casablanca",
				"released": 1942,
				"Actors": [
					"Humphrey Bogart",
					"Ingrid Bergman"
				]
			},
			{
				"Title": "Cool Hand Luke",
				"released": 1967,
				"color": true,
				"Actors": [
					"Paul Newman"
				]
			},
			{
				"Title": "Bullitt",
				"released": 1968,
				"color": true,
				"Actors": [
					"Steve McQueen",
					"Jacqueline Bisset"
				]
			}
		]
	*/
}
