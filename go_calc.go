package main

import "fmt"

func main() {
	const (
		plus  = "+"
		minus = "-"
		multi = "*"
		divis = "/"
		dwr   = "%"
	)

	var a, b int
	var ch string

	fmt.Scan(&a, &ch, &b)

	switch ch {
	case plus:
		fmt.Println(a + b)
	case minus:
		fmt.Println(a - b)
	case multi:
		fmt.Println(a * b)
	case divis:
		if b != 0 {
			fmt.Println(a / b)
		} else {
			fmt.Println("На ноль делить нельзя")
		}
	case dwr:
		if b != 0 {
			fmt.Println(a % b)
		} else {
			fmt.Println("На ноль делить нельзя")
		}
	}
}
