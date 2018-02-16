package main

import (
	"fmt"
	"time"

	"github.com/wuriyanto48/floop"
)

func someFunc() {
	res := mul(5, 5)
	fmt.Println(res)
}

func mul(a int, b int) int {
	return a * b
}

func main() {
	f := floop.New(1)
	f.Start(someFunc, time.Second*5)

}
