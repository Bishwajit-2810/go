package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func WebRequest() {
	fmt.Println("connecting web server")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("there is error in getting response", err)
		return
	}
	fmt.Printf("%T\n", res)
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error while extracting data", err)
		return
	}
	fmt.Println(string(data))

}
