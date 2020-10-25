package main

import (
	"fmt"

	"github.com/raspiantoro/warta"
)

func main() {

	var helloTopic warta.TopicName = "hello"

	helloFunc := func(fName string, lName string) {
		fmt.Printf("Hello %s %s\n", fName, lName)
	}

	warta := warta.New()

	l, err := warta.ListenCreate(helloTopic, helloFunc)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	warta.BroadcastCreate(helloTopic, "John", "Doe")

}
