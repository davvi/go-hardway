package main

import "fmt"

func main() {
    fmt.Println("I will count my chickens")
    
    fmt.Println("Hens", 25 + 30 / 6)
    fmt.Println("Roosters", 100 - 25 % 4)
    fmt.Println("Now I will count the eggs:")
    fmt.Println(3 + 2 + 1 - 5 + 4 % 2 - 1 / 4 + 6)
    fmt.Println("Is it true that 3 + 2 < 5 - 7?")
    fmt.Println(3 + 2 < 5 - 7)
    fmt.Println("What is 3 + 2?", 3 + 2)
    fmt.Println("What is 5 - 7", 5 - 7)
    fmt.Println("Oh, that's why it's False.")
    fmt.Println("How about some more.")
    fmt.Println("Is it greater?", 5 > -2)
    fmt.Println("Is it less or equal?", 5 <= -2)
}
