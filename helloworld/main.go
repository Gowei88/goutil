package main

import (
	"github.com/myuser/calculator"
	"rsc.io/quote"
	"strconv"
)

func main() {
	//number1, _ := strconv.Atoi(os.Args[1])
	//number2, _ := strconv.Atoi(os.Args[2])
	//fmt.Println("Sum", number1+number2)
	//sum, _ := sum(os.Args[1], os.Args[2])
	//println("sum", sum)
	//println("mul", mul)
	//firstName := "John"
	//updateName(&firstName)
	//println(firstName)
	total := calculator.Sum(5, 3)
	println(total)
	println("Version", calculator.Version)
	println(quote.Hello())
}
func sum(number1 string, number2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mul = int1 * int2
	return
}
func updateName(name *string) {
	*name = "David"
}
