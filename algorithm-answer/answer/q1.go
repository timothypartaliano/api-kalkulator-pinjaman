package answer

import "fmt"

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i < n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func Q1() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if isPrime(num) {
        fmt.Println("A Prime Number")
    } else {
        fmt.Println("Not a Prime Number")
    }
}