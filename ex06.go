package main

import "fmt"

func main() {
    x := fmt.Sprintf("There are %s types of people.", "10")
    binary := "binary"
    do_not := "don't"
    y := fmt.Sprintf("Those who know %s and those who %s.", binary, do_not)

    fmt.Println(x)
    fmt.Println(y)

    fmt.Printf("I said: %s\n", x)
    fmt.Printf("I also said: %s\n", y)

    hilarious := "false\n"
    joke_evaluation := "Isn't that joke so funny?! %s"

    fmt.Printf(joke_evaluation, hilarious)

    w := "This is the left side of..."
    e := "a string with a right side."

    fmt.Println(w + e)
}
