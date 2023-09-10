package main

import "fmt"

func main() {
    var num1, num2 int
    var operator string

    fmt.Print("Enter first number: ")
    fmt.Scanln(&num1)
    fmt.Print("Enter second number: ")
    fmt.Scanln(&num2)
    fmt.Print("Enter operator (+, -, *, /): ")
    fmt.Scanln(&operator)

    if operator == "+" {
        fmt.Printf( num1+num2)
    } else if operator == "-" {
        fmt.Printf( num1-num2)
    } else if operator == "*" {
        fmt.Printf( num1*num2)
    } else if operator == "/" {
        if num2 != 0 {
            fmt.Printf( float64(num1)/float64(num2))
        } else {
            fmt.Println("Error: division by zero")
        }
    } else {
        fmt.Println("Invalid operator")
    }
}
