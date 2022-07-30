package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func main() {
	fmt.Println("Hello world!")

	helloWorldBytes := []byte("Hello world!\n")

	if err := ioutil.WriteFile("./helloworld.txt", helloWorldBytes, 0644); err != nil {
		panic(err)
	}

	fmt.Println("http://localhost:8090")
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
