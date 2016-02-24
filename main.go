package main

import (
	"fmt"
	"os"
)

func main() {
	op1, op2, op := ParseArgs()

	var answer Number
	var overflow bool
	switch op {
	case "+":
		answer, overflow = Sum(op1, op2)
	case "-":
		answer, overflow = Sum(op1, op2.Neg())
	default:
		fmt.Fprintln(os.Stderr, "invalid operation:", op)
		os.Exit(1)
	}

	fmt.Println("Answer:", answer)
	fmt.Println("Overflow:", overflow)
}
