package main

import (
    "fmt"
    "adventOfCode/helpers"
)

/*
The next year, to speed up the process, Santa creates a robot version of himself, Robo-Santa, to deliver presents with him.

Santa and Robo-Santa start at the same location (delivering two presents to the same starting house), then take turns moving based on instructions from the elf, who is eggnoggedly reading from the same script as the previous year.

This year, how many houses receive at least one present?

For example:

^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.
^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.
^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.

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
    santa := Coordinates{0, 0}
    robot := Coordinates{0, 0}
    visited[santa] = true

    index := 0
    for _, m := range input {
        direction := string(m)

        if direction != "" {
            var whosMoving Coordinates

            isSantaMoving := index % 2 == 0
            if isSantaMoving {
                whosMoving = santa
            } else {
                whosMoving = robot
            }

            newCoords := count(direction, whosMoving)

            if _, ok := visited[newCoords]; !ok {
                visited[newCoords] = true
            }

            if isSantaMoving {
                santa = newCoords
            } else {
                robot = newCoords
            }
        }
        index++
    }

    fmt.Println(len(visited))
}
