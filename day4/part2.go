package main

import (
    "crypto/md5"
	"fmt"
	"io"
    "encoding/hex"
    "strconv"
)

/*
 Now find one that starts with six zeroes.
 */

var secret = "iwrupvqb"

func main () {
    i := 0

    var found int

    for {
        hash := md5.New()
        io.WriteString(hash, secret + strconv.Itoa(i))

        result := hash.Sum(nil)
        head := hex.EncodeToString(result)[0:6]

        if head == "000000" {
            found = i
            break
        }
        i++
    }

    fmt.Println(found)
}
