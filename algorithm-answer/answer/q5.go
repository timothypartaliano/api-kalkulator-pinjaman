package answer

import "fmt"

func swapWithoutTemp(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func Q5() {
	a := 7
	b := 10

	fmt.Println("Input:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	a, b = swapWithoutTemp(a, b)

	fmt.Println("Output:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}