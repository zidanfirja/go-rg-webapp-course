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
	Flags    Flags  `json:"flags"`
	Safe     bool   `json:"safe"`
}

type Flags struct {
	Nsfw      bool `json:"nsfw"`
	Religious bool `json:"religious"`
	Political bool `json:"political"`
	Racist    bool `json:"racist"`
	Sexist    bool `json:"sexist"`
	Explicit  bool `json:"explicit"`
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

	// get an anime quote from the anime API
	anime, err := GetAnime()
	if err != nil {
		fmt.Println("Error while getting an anime quote from the anime API :", err)
		return
	}

	/*
			- Format of the output: <anime> said "<setup>", <delivery>
			- add flags and safe in joke struct
		    - check if joke is safe, then print the joke, check if joke is not safe, then print "This joke is not safe for work, because it contains <flags yang true>"

				Example output:
				Fuyou Kaede said joke "I'm not saying my son is ugly...", "But on Halloween he went to tell the neighbors to turn down their TV and they gave him some candy."
	*/

	if !joke.Safe {
		// check flags true

		flags := []string{}
		if joke.Flags.Nsfw {
			flags = append(flags, "nsfw")
		}
		if joke.Flags.Religious {
			flags = append(flags, "religious")
		}
		if joke.Flags.Political {
			flags = append(flags, "political")
		}
		if joke.Flags.Racist {
			flags = append(flags, "racist")
		}
		if joke.Flags.Sexist {
			flags = append(flags, "sexist")
		}

		strFlags := ""
		for i, flag := range flags {
			if i == 0 {
				strFlags += flag
			} else {
				strFlags += ", " + flag
			}
		}

		fmt.Printf("This joke is not safe for work, because it contains %v \n", strFlags)
		fmt.Printf("setup: %v, delivery: %v \n", joke.Setup, joke.Delivery)
	} else {
		fmt.Printf("%v said joke \"%v\", \"%v\" \n", anime.Character, joke.Setup, joke.Delivery)
	}

}
