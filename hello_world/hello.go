package main

import (
	"fmt"
    "time"
)

func greenting() string {
    return "hi mom. the time is: " + time.Now().String()
}

func main() {
	fmt.Println(greenting())
}
