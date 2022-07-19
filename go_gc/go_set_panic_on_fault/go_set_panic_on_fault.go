package main

import (
	"log"
	"time"
)

func main() {
	var a1 []int
	log.Println(a1[100])
	//debug.SetPanicOnFault()
	var a map[int]string
	b := a[22]
	log.Println(b, len(a))
	log.Printf("%p", a)

	log.Println(a[1], a[2])

	time.Sleep(time.Second)
}
