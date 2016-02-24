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

	for _, base := range []int{2, 8, 10, 16} {
		fmt.Println("Signed answer in base", base, "is  ", answer.BaseStringSigned(base))
		fmt.Println("Unsigned answer in base", base, "is", answer.BaseStringUnsigned(base))
	}
	fmt.Println("Overflow:", overflow)
}
