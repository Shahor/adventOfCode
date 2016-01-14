package main

import (
    "fmt"
    "adventOfCode/helpers"
)

/*
Santa is delivering presents to an infinite two-dimensional grid of houses.

He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

For example:

> delivers presents to 2 houses: one at the starting location, and one to the east.
^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.

 */

type Coordinates struct {
    x int
    y int
}

func count(direction string, coords Coordinates) Coordinates {
    switch direction {
    case "^":
        coords.y++
    case "v":
        coords.y--
    case ">":
        coords.x++
    case "<":
        coords.x--
    }

    return coords
}

func main() {
    visited := make(map[Coordinates]bool)
    input := Input.Get("day3/input.txt")

    // input = "^>v<"
    coords := Coordinates{0, 0}
    visited[coords] = true

    for _, m := range input {
        direction := string(m)

        if direction != "" {
            newCoords := count(direction, coords)

            if _, ok := visited[newCoords]; !ok {
                visited[newCoords] = true
            }

            coords = newCoords
        }
    }

    fmt.Println(len(visited))
}
