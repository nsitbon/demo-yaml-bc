package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

var data = `
hello:
  <<: *mydefaults

mydefaults: &mydefaults
  greeting: "Hello"
`

func main() {
	var t struct {}

	err := yaml.Unmarshal([]byte(data), &t)

	if err != nil {
		panic(err)
	}

	fmt.Println("OK")
}