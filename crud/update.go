package crud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func UpdateGo() {
	todos1 := Todo{
		UserId:    1414,
		Title:     "bishwajit",
		Completed: true,
	}
	jsonData, err := json.Marshal(todos1)
	if err != nil {
		fmt.Println("err in  marshaling", err)
		return
	}
	jsonString := string(jsonData)
	jsonReader := strings.NewReader(jsonString)

	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", jsonReader)
	if err != nil {
		fmt.Println("err in  put", err)
		return
	}
	req.Header.Set("Content-type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("err in  sending req", err)
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
