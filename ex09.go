package main

// Here's some new strange stuff, remember type it exactly.
import "fmt"
func main() {
    days := "Mon Tue Wed Thur Fri Sat Sun"
    months := "Jan\nFeb\nMar\nApr\nMay\nJun\nJul\nAug"

    fmt.Println("Here are the days:", days)
    fmt.Println("Here are the months:", months)
   
    
    longstring :=
`
There's something going on here. 
With the back ticks.
We'll be able to type as much as we like.
Even 4 lines if we want, or 5, or 8.`

    fmt.Println(longstring)
}
