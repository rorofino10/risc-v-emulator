package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	t := time.NewTicker(time.Second * 3)
	for range t.C {
		fmt.Println("tick")
	}
}
