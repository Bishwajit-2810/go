package crud

import (
	"fmt"
	"net/http"
)

func DeleteGo() {
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println("err in  delete", err)
		return
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("err in  sending req", err)
		return
	}
	defer res.Body.Close()
	fmt.Println(res.StatusCode)
}
