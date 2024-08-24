package crud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func PostGo() {
	todos1 := Todo{
		UserId:    1414,
		Title:     "bishwajit",
		Completed: true,
	}

	// convert todo struct to json
	jsonData, err := json.Marshal(todos1)
	if err != nil {
		fmt.Println("err in  marshaling", err)
		return
	}
	jsonString := string(jsonData)
	jsonReader := strings.NewReader(jsonString)

	res, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", jsonReader)
	if err != nil {
		fmt.Println("error sending request", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
	fmt.Println(todos1)
	fmt.Println(todos1.Title)
	fmt.Println(todos1.Completed)
	fmt.Println(res.StatusCode)

}
