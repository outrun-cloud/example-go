package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello world!")

	helloWorldBytes := []byte("Hello world!\n")

	if err := ioutil.WriteFile("./helloworld.txt", helloWorldBytes, 0644); err != nil {
		panic(err)
	}
}
