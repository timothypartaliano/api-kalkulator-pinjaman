package answer

import "fmt"

func listPrimes(n int) []int {
    primeList := []int{}
    for num := 2; num <= n; num++ {
        isPrime := true
        for i := 2; i < num; i++ {
            if num%i == 0 {
                isPrime = false
                break
            }
        }
        if isPrime {
            primeList = append(primeList, num)
        }
    }
    return primeList
}

func Q2() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    primeNumbers := listPrimes(num)
    fmt.Println(primeNumbers)
}