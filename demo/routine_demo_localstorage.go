package main

import (
	"fmt"
	"github.com/isyscore/go-goid"
	"time"
)

var nameVar = goid.NewLocalStorage()

func main() {
	nameVar.Set("hello world")
	fmt.Println("name: ", nameVar.Get())

	// other goroutine cannot read it
	go func() {
		fmt.Println("name1: ", nameVar.Get())
	}()

	// but, the new goroutine could inherit/copy all local data from the current goroutine like this:
	goid.Go(func() {
		fmt.Println("name2: ", nameVar.Get())
	})

	// or, you could copy all local data manually
	ic := goid.BackupContext()
	go func() {
		goid.InheritContext(ic)
		fmt.Println("name3: ", nameVar.Get())
	}()

	time.Sleep(time.Second)
}
