package main

import (
	"fmt"
	"strings"
)

func Binary(code string) string {
	inst := NewInstruction(code)
	if inst.IsA() {
		addr := inst.Address()
		if addr < 0 {
			return strings.Repeat("x", 16)
		}

        return fmt.Sprintf("%016b", addr)
	}

	return "0000000000000000"
}

func BinaryInt(addr int64) string {
	return fmt.Sprintf("%016b", addr)
}
