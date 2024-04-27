package main

import (
	"encoding/json"
	"fmt"
)

var strJson = `{"first_name": "John", "last_name": "Doe", "age": 25, "hobi": ["membaca", "menulis"], "arr": [1, 2, 3, 4, 5], "alamat": {"jalan": "Jalan 1", "kota": "Jakarta", "provinsi": "DKI Jakarta"}}`

/*
	{
	  "first_name": "John",
	  "last_name": "Doe",
	  "age": 25,
	  "hobi": [
	    "membaca",
	    "menulis"
	  ],
	  "arr": [
	    1,
	    2,
	    3,
	    4,
	    5
	  ],
	  "alamat": {
	    "jalan": "Jalan 1",
	    "kota": "Jakarta",
	    "provinsi": "DKI Jakarta"
	  }
	}
*/

/*
  - When we mapping from JSON to struct, the field name must be the same as the JSON key, it can use tag to change the field name
  - Mapping all json data to struct is not necessary, we can choose which data we want to map
*/

type User struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	Hobbies   []string `json:"hobi"`
	Arr       []int    `json:"arr"`
	Address   Address  `json:"alamat"`
}

type Address struct {
	Street   string `json:"jalan"`
	City     string `json:"city"`
	Province string `json:"provinsi"`
}

func main() {
	bagas := User{}

	// decode strJson to struct User
	// use json.Unmarshal([]byte, &struct)
	err := json.Unmarshal([]byte(strJson), &bagas)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("First Name:", bagas.FirstName)
	fmt.Println("Last Name:", bagas.LastName)
	fmt.Println("Age:", bagas.Age)
	fmt.Println("Hobbies:", bagas.Hobbies)
	fmt.Println("Arr:", bagas.Arr)
	fmt.Println("Address:", bagas.Address)
	fmt.Println()

	salsa := User{
		FirstName: "Salsa",
		LastName:  "Ainun",
		Age:       22,
		Hobbies:   []string{"membaca", "menulis"},
		Address: Address{
			Street:   "Jalan 2",
			City:     "Jakarta",
			Province: "DKI Jakarta",
		},
	}

	// encode struct User to JSON
	// use json.Marshal(struct)
	jsonSalsa, err := json.Marshal(salsa)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println(string(jsonSalsa))
}
