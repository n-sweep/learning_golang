package main

import (
	"fmt"
    "time"
)

// go run hi_mom.go

func greenting() string {
    return "hi mom, the time is: " + time.Now().String()
}

func main() {
	fmt.Println(greenting())
}
