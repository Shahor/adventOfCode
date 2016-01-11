package Input

import "os"
import "io/ioutil"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Get(path string) string {
    pwd, err := os.Getwd()
    check(err)

    dat, err := ioutil.ReadFile(pwd + "/" + path)
    check(err)

    return string(dat)
}
