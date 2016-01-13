package main

import (
    "fmt"
    "adventOfCode/helpers"

    "strings"
    "strconv"
    "sort"
)


/*
The elves are also running low on ribbon. Ribbon is all the same width, so they only have to worry about the length they need to order, which they would again like to be exact.

The ribbon required to wrap a present is the shortest distance around its sides, or the smallest perimeter of any one face. Each present also requires a bow made out of ribbon as well; the feet of ribbon required for the perfect bow is equal to the cubic feet of volume of the present. Don't ask how they tie the bow, though; they'll never tell.

For example:

A present with dimensions 2x3x4 requires 2+2+3+3 = 10 feet of ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
A present with dimensions 1x1x10 requires 1+1+1+1 = 4 feet of ribbon to wrap the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.
How many total feet of ribbon should they order?
 */

func main() {
    input := Input.Get("day2/input.txt")
    gifts := strings.Split(input, "\n")
    total := 0

    for i := 0; i < len(gifts); i++ {
        gift := strings.TrimSpace(gifts[i])
        if len(gift) > 0 {
            parsed := parse(gift)

            total += computeRibbon(parsed)
        }
    }

    fmt.Println(total)
}

func parse(str string) []int {
    dimensions := strings.Split(str, "x")
    l, _ := strconv.Atoi(dimensions[0])
    w, _ := strconv.Atoi(dimensions[1])
    h, _ := strconv.Atoi(dimensions[2])

    sizes := []int{ l, w, h }

    sort.Ints(sizes)
    return sizes
}

func computePresentWrap(sizes []int) int {
    return 2 * sizes[0] + 2 * sizes[1]
}

func computeBow(sizes []int) int {
    return sizes[0] * sizes[1] * sizes[2]
}

func computeRibbon(sizes []int) int {
    return computePresentWrap(sizes) + computeBow(sizes)
}
