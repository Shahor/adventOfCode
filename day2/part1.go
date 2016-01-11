package main

import (
    "fmt"
    "adventOfCode/helpers"

    "strings"
    "strconv"
)


/*
The elves are running low on wrapping paper, and so they need to submit an order for more.
They have a list of the dimensions (length l, width w, and height h) of each present, and only want to order exactly as much as they need.

Fortunately, every present is a box (a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier:
find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves also need a little extra paper for each present: the area of the smallest side.

For example:

A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.
All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?
 */

func main() {
    input := Input.Get("day2/input.txt")
    gifts := strings.Split(input, "\n")
    total := 0

    for i := 0; i < len(gifts); i++ {
        gift := strings.TrimSpace(gifts[i])
        if len(gift) > 0 {
            parsed := parse(gift)

            total += computeWrappingPaper(parsed)
        }
    }

    fmt.Println(total)
}

func parse(str string) map[string]int {
    _map := make(map[string]int)

    dimensions := strings.Split(str, "x")
    l, _ := strconv.Atoi(dimensions[0])
    w, _ := strconv.Atoi(dimensions[1])
    h, _ := strconv.Atoi(dimensions[2])

    _map["l"] = l
    _map["w"] = w
    _map["h"] = h

    return _map
}

func getSmallestSide(sides map[string]int) int {
    ret := sides["face"]

    if (sides["bottom"] <= sides["side"] && sides["bottom"] <= sides["face"]) {
        ret = sides["bottom"]
    } else if (sides["side"] <= sides["bottom"] && sides["side"] <= sides["face"]) {
        ret = sides["side"]
    }

    return ret
}

func computeSides (_map map[string]int) map[string]int {
    sides := make(map[string]int)

    sides["bottom"] = _map["l"] * _map["w"]
    sides["side"] = _map["w"] * _map["h"]
    sides["face"] = _map["h"] * _map["l"]

    return sides
}

func computeWrappingPaper (_map map[string]int) int {
    sides := computeSides(_map)
    smallest := getSmallestSide(sides)

    return 2 * sides["bottom"] +
        2 * sides["side"] +
        2 * sides["face"] +
        smallest
}
