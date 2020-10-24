package main

import (
	"fmt"
	"log"

	"github.com/raspiantoro/warta"
)

func main() {
	var err error
	e := warta.New()

	var en warta.TopicName = "greet"

	fmt.Println("listen to greet topic with callback")
	cb := callBack
	l, err := e.ListenCreate(en, cb)
	if err != nil {
		log.Println(err)
		return
	}

	// l.Callback()

	// err = l.Close()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	fmt.Println("listen to greet topic with callback2")
	cb2 := callBack2
	// _, err = e.Listen(en, cb2)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	go e.Listen(en, cb2)
	go e.Listen(en, cb2)
	go e.Listen(en, cb2)
	go e.Listen(en, cb2)
	go e.Listen(en, cb2)
	go e.Listen(en, cb2)
	go e.Listen(en, cb2)
	go e.Listen(en, cb2)
	go e.Listen(en, cb2)

	fmt.Println("close listener for callback")
	go l.Close()
	go l.Close()
	go l.Close()
	go l.Close()
	go l.Close()
	go l.Close()

	// fmt.Println("listen to greet topic with callback3")
	// cb3 := callBack3
	// err = e.On(en, cb3)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	_, err = e.CreateTopic("calc")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("listen to calc topic with callback4")
	cb4 := callBack4
	l2, err := e.Listen("calc", cb4)
	if err != nil {
		log.Println(err)
		return
	}

	go l2.Close()
	go l2.Close()
	go l2.Close()
	go l2.Close()
	go l2.Close()

	fmt.Println("listen to calc topic with callback5")
	cb5 := callBack5
	_, err = e.Listen("calc", cb5)
	if err != nil {
		log.Println(err)
		return
	}

	go e.Listen("calc", cb5)
	go e.Listen("calc", cb5)
	go e.Listen("calc", cb5)
	go e.Listen("calc", cb5)
	// fmt.Println("listen to calc topic with callback4")
	// cb6 := callBack5
	// err = e.On("calc", cb6)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// fmt.Println("listen to calc topic with callback4")
	// cb7 := callBack5
	// err = e.On("calc", cb7)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// fmt.Println("listen to calc topic with callback4")
	// cb8 := callBack5
	// err = e.On("calc", cb8)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	//time.Sleep(1 * time.Second)

	// t, err := e.Topic("calc")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// err = t.RemoveListener(cb4)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	err = e.BroadcastCreate(en, "Mario", "Raspiantoro")
	if err != nil {
		log.Println(err)
		return
	}

	err = e.Broadcast("calc", 3, 2)
	if err != nil {
		log.Println(err)
		return
	}

}

func callBack(firstName string, lastName string) {
	fmt.Printf("Hello %s %s\n", firstName, lastName)
}

func callBack2(firstName string, lastName string) {
	fmt.Printf("Hello 2 %s %s\n", firstName, lastName)
}

func callBack3(firstName string, lastName string) {
	fmt.Printf("Hello 3 %s %s\n", firstName, lastName)
}

func callBack4(a int, b int) {
	fmt.Println(a * b)
}

func callBack5(a int, b int) {
	fmt.Println(a - b)
}

func callBack6(a int, b int) int {
	return a + b
}
