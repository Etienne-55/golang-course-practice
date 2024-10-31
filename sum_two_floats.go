package main

import (
    "fmt"
)

func main() {
    var float1, float2 float64

    fmt.Print("Enter the first float: ")
    _, err := fmt.Scan(&float1)
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    fmt.Print("Enter the second float: ")
    _, err = fmt.Scan(&float2)
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    sum := float1 + float2

    fmt.Printf("The sum of %.2f and %.2f is %.2f\n", float1, float2, sum)
}
