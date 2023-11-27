package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "time"
)

func randint(a, b int) int {
    return rand.Intn(b - a) + a
}

func main() {

    rand.Seed(time.Now().UnixNano())

    responses := [3]string{
        "yes",
        "no",
        "maybe",
    }

    v := randint(0, len(responses))

    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Submit your query, peasant.")
    input, _ := reader.ReadString('\n')

    fmt.Println(responses[v])

}
