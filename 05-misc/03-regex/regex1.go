package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	r, _ := regexp.Compile(`^(?P<protocol>\w+):\/\/(?P<domain>(\w+\.\w+))\/?(?P<route>\w+)?\??(?P<query>.*)?`)
	m := r.FindStringSubmatch("http://myapi.com/products?page=1&offset=2")

	result := make(map[string]string)

	if len(m) == 0 {
		fmt.Println("No valid URL provided.")
		os.Exit(0)
	}

	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = m[i]
		}
	}

	fmt.Println(result)
}
