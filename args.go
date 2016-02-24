package main

import (
	"fmt"
	"os"
	"strconv"
)

func ParseArgs() (op1 Number, op2 Number, op string) {
	var err error
	bits := 32

	if len(os.Args) < 2 {
		dieUsage()
	}

	opArgs := os.Args[1:]

	if opArgs[0] == "-bits" {
		if len(opArgs) < 3 {
			dieUsage()
		}
		bits, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		opArgs = opArgs[2:]
	}

	if len(opArgs) != 1 && len(opArgs) != 3 {
		dieUsage()
	}

	op1, err = ParseNumber(opArgs[0], bits)
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid number:", opArgs[0])
		os.Exit(1)
	}

	op = "+"
	op2 = Number{false}.SignExtend(bits)

	if len(opArgs) == 3 {
		op = opArgs[1]
		op2, err = ParseNumber(opArgs[2], bits)
		if err != nil {
			fmt.Fprintln(os.Stderr, "invalid number:", opArgs[2])
			os.Exit(1)
		}
	}

	return
}

func dieUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [flags] <operand> [<operation> <operand>]\n"+
		"\nAvailable flags are:\n\n"+
		"-bits <int>      the bit size to use for operands", os.Args[0])
	os.Exit(1)
}
