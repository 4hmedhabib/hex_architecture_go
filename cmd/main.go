package main

import (
	"fmt"
	"hex_arch_go/internal/adapters/core/arithmetic"
	"hex_arch_go/internal/ports"
)

func main() {
	// ports
	var core ports.ArithmeticPort

	core = arithmetic.NewAdapter()
	result, err := core.Subtraction(1, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
