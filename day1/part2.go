package main

import "fmt"
import "adventOfCode/helpers"

func main() {
    count := 0

    input := Input.Get("day1/input.txt")
    for i, c := range input {
        switch string(c) {
            case "(":
                count++
            case ")":
                count--
        }

        if count == -1 {
            fmt.Println(i + 1)
            break
        }
    }
}
