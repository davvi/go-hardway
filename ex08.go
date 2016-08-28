package main

import "fmt"

func main() {
    formatter := "%#v %#v %#v %#v \n"

    fmt.Printf(formatter, 1, 2, 3, 4)
    fmt.Printf(formatter, "one", "two", "three", "four")
    fmt.Printf(formatter, true, false, true, false)
    fmt.Printf(formatter, "I had this thing.", "That you could type up right.", "But it didn't sing.", "So I said good night.")
}
