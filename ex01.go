package main
//The above tells Go to that this package is the main executable and not a "shared library"
//The line below will import the fmt package, which is short for "Format"
import "fmt"

//The main function is the entry point into the program. This is where all the things are executed.
//All these lines are similar. They call Println, a method of fmt. It will print the string in the brackets to stout"
//Note: The double quotes in the last line around the word "said" had to be escaped.
func main() {
    fmt.Println("Hello World")
    fmt.Println("Hello Again")
    fmt.Println("I like typing this.")
    fmt.Println("This is fun.")
    fmt.Println("Yay! Printing.")
    fmt.Println("I'd must rather you 'not'.")
    fmt.Println("I \"said\" do not touch this.")
}
