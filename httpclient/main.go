package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
Demo of create a http client.

- GET request
 url : https://v2.jokeapi.dev/joke/Any?type=twopart
*/

type Joke struct {
	Category string `json:"category"`
	Type     string `json:"type"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
	Error    bool   `json:"error"`
}

type Anime struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

var jokeURL = "https://v2.jokeapi.dev/joke/Any?type=twopart"
var animeURL = "https://animechan.xyz/api/random"

func GetJoke() (*Joke, error) {
	// make a GET request to the joke API
	resp, err := http.Get(jokeURL)
	if err != nil {
		fmt.Println("Error while making a GET request to the joke API :", err)
		return nil, err
	}
	defer resp.Body.Close()

	// decode the response from the joke API to a Joke struct
	var joke Joke
	err = json.NewDecoder(resp.Body).Decode(&joke)
	if err != nil {
		fmt.Println("Error while decoding the response from the joke API :", err)
		return nil, err
	}

	return &joke, nil
}

func GetAnime() (*Anime, error) {
	resp, err := http.Get(animeURL)
	if err != nil {
		fmt.Println("Error while making a GET request to the anime API :", err)
		return nil, err
	}

	var anime Anime
	err = json.NewDecoder(resp.Body).Decode(&anime)
	if err != nil {
		fmt.Println("Error while decoding the response from the anime API :", err)
		return nil, err
	}

	return &anime, nil
}

func main() {

	// get a joke from the joke API
	joke, err := GetJoke()
	if err != nil {
		fmt.Println("Error while getting a joke from the joke API :", err)
		return
	}

	// print the joke
	fmt.Println("Category :", joke.Category)
	fmt.Println("Type :", joke.Type)
	fmt.Println("Setup :", joke.Setup)
	fmt.Println("Delivery :", joke.Delivery)

	// get an anime quote from the anime API
	anime, err := GetAnime()
	if err != nil {
		fmt.Println("Error while getting an anime quote from the anime API :", err)
		return
	}

	// print the anime quote
	fmt.Printf("%s said \"%s\" in %s \n", anime.Character, anime.Quote, anime.Anime)

	/*
			- Format of the output: <anime> said "<setup>", <delivery>
			- add flags and safe in joke struct
		    - check if joke is safe, then print the joke, check if joke is not safe, then print "This joke is not safe for work, because it contains <flags yang true>"

				Example output:
				Fuyou Kaede said joke "I'm not saying my son is ugly...", "But on Halloween he went to tell the neighbors to turn down their TV and they gave him some candy."
	*/

}
