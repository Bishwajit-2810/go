package json_go

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func JsonGo() {
	fmt.Println("json")

	person := Person{Name: "Bishwajit", Age: 22, IsAdult: true}
	fmt.Println(person)

	// convert person to json
	jasonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error marshaling", err)
		return
	}
	fmt.Println(string(jasonData))

	// unmarshaling data
	var personData Person
	err = json.Unmarshal(jasonData, &personData)
	if err != nil {
		fmt.Println("error unmashaling")
		return
	}
	fmt.Println(personData)

}
