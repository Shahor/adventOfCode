package main

import "fmt"
import "adventOfCode/helpers"

func main() {
    count := 0

    input := Input.Get("day1/input.txt")
    for _, c := range input {
        switch string(c) {
            case "(":
                count++
            case ")":
                count--
        }
    }

    fmt.Println(count)
}
