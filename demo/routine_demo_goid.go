package main

import (
	"fmt"
	"github.com/isyscore/go-goid"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Second)
	}()
	id := goid.Goid()
	ids := goid.AllGoids()
	fmt.Printf("curr goid: %d\n", id)
	fmt.Printf("all goids: %v\n", ids)
}
