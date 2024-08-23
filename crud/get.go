package crud

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// everything depends on the name of the variable if completed is not define in json by default it will use the named variable

type Todo struct {
	UserId    int    `json:"userid"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetGo() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error while getting", err)
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("error getting response", res.StatusCode)
	}

	// 1 way of getting data

	// defer res.Body.Close()
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("error reading", err)
	// 	return
	// }
	// fmt.Println(string(data))

	// another way of getting data
	// better way

	var todos Todo
	err = json.NewDecoder(res.Body).Decode(&todos)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(todos)
	fmt.Println(todos.Title)
	fmt.Println(todos.Completed)

}
