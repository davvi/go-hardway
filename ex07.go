package main

import ("fmt"
        "strings"
)

func main() {
    fmt.Println("Mary had a little lamb")
    fmt.Println("It's fleece was white as snow")
    fmt.Println("And everywhere that Mary went.")
    // You can do operations on strings like you can in python.. so you can do the following.
    y := strings.Repeat("*", 10)
    
    fmt.Println(y)

    end1 := "C"
    end2 := "h"
    end3 := "e"
    end4 := "e"
    end5 := "s"
    end6 := "e"
    end7 := "B"
    end8 := "u"
    end9 := "r"
    end10 := "g"
    end11 := "e"
    end12 := "r"
    
    fmt.Println(end1 + end2 + end3 + end4 + end5 + end6 + end7 + end8 + end9 + end10 + end11 + end12)
    
}
