package main

import (
	"fmt"
	"net/url"
)

func main() {
	myURL := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	fmt.Printf("Type of URl: %T\n", myURL)

	parsedURL, err := url.Parse(myURL)

	if err != nil {
		fmt.Println("can't parse URL", err)
		return
	}

	fmt.Printf("Type of parsed URl: %T\n", parsedURL)

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Port:", parsedURL.Port())
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Raw Query:", parsedURL.RawQuery)
	// fmt.Println("Raw Query:", url.Values(parsedURL.Query()))

	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=3xp10it3r"

	newUrl := parsedURL.String()
	fmt.Println(newUrl)
	

}
