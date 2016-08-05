package main

import "fmt"

func main() {
    var my_name string = "Andrew"
    var my_age int = 34 // not a lie
    var my_height int = 189 // in cm
    var my_weight int = 98 // kg ... yeah I need to lay off the donuts
    var my_eyes string = "Blue"
    var my_teeth string = "White"
    var my_hair string = "Brown"

    fmt.Printf("Let's talk about %[1]s.\n", my_name)
    fmt.Printf("He's %[1]d centremeters tall.\n", my_height)
    fmt.Printf("He's %[1]d kilos heavy.\n", my_weight)
    fmt.Println("Yep, need to lay of them donuts")
    fmt.Printf("He's got %[1]s eyes and %d hair.\n", my_eyes, my_hair)
    fmt.Printf("His teeth are ususally %[1]s depending on the coffee.\n", my_teeth)
    fmt.Printf("If I add %[1]d, %[2]d and %[3]d I get %[4]d.\n", my_age, my_height, my_weight, my_age + my_height + my_weight)
}
