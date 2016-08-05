package main

import "fmt"

func main() {
    var cars int = 100
    var space_in_a_car int = 4
    var drivers int = 30
    var passengers int = 90
    var cars_not_driven int = cars - drivers
    var cars_driven int = drivers
    var carpool_capacity int = cars_driven * space_in_a_car
    var average_passengers_per_car int = passengers / cars_driven

    fmt.Println("There are", cars,"cars available.")
    fmt.Println("There are only", drivers,"drivers available.")
    fmt.Println("There will be", cars_not_driven, "empty cars today.")
    fmt.Println("We can transport", carpool_capacity, "people today.")
    fmt.Println("We have", passengers, "to carpool today.")
    fmt.Println("We need to put about", average_passengers_per_car, "in each car.")
}
