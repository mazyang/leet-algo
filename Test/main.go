package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan int
	timeout := time.After(5 * time.Second)
	select {
	case s := <-c:
		fmt.Println(s)
	case <-timeout:
		fmt.Println("You talk too much.")
		return
	}
}
