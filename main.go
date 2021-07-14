package main

import (
	"fmt"
	"goboy/gb"
)

func main() {
	var cpu gb.CPU
	cpu.ResetRegisters()
	fmt.Println(&cpu)
}
