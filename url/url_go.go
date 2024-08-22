package url_go

import (
	"fmt"
	"net/url"
)

func UrlGo() {
	myUrl := "https://example.com:8080/path/to/resource/key1=value1&key2=value2"
	fmt.Printf("%T\n", myUrl)

	parsedUrl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("parsing failed")
		return
	}
	fmt.Printf("%T\n", parsedUrl)
	fmt.Println(parsedUrl.Scheme)
	fmt.Println(parsedUrl.Host)
	fmt.Println(parsedUrl.Path)
	fmt.Println(parsedUrl.RawQuery)

	parsedUrl.Path = "/newPath"
	parsedUrl.RawQuery = "username=bishwajit"

	// constructing new url
	newUrl := parsedUrl.String()
	fmt.Println(newUrl)

}
