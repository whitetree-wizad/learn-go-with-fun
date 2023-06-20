package main

import (
	"fmt"
	"net/url"
)

func main() {
	fakeurl := "https://a b.com"
	_, err := url.Parse(fakeurl)

	fmt.Printf("%v\n", err)

	value, ok := err.(*url.Error)
	if ok {
		fmt.Printf("%v\n%v\n%v", value.Err, value.URL, value.Op)
	}
}
