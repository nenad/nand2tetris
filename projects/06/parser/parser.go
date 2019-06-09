package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type Parser struct {
	// Current points to the current instruction
	Current int64
	// Labels stores the labels and their memory address
	Labels map[string]string
	// Vars stores the variables and their memory address
	Vars map[string]int64
}

// Parse processes provided assembly code and returns byte array consisted
// of binary representation of the assembly code. An error is returned if
// there was an issue parsing the provided data.
func (p *Parser) Parse(data io.Reader) ([]byte, error) {
	p.Current = 0
	p.Labels = map[string]string{}
	p.Vars = map[string]int64{}

	s := bufio.NewScanner(data)
	buf := bytes.Buffer{}
	// First pass to sanitize the file and populate Label map
	for s.Scan() {
		instruction := NewInstruction(s.Text())
		if instruction.Empty() {
			continue
		}

		// Check if current line is label
		label := instruction.Label()
		if label != "" {
			p.Labels[label] = "@" + strconv.FormatInt(p.Current, 10)
			continue
		}

		buf.WriteString(instruction.String() + "\n")

		p.Current++
	}

	p.Current = 0
	s = bufio.NewScanner(&buf)
	binBuf := bytes.Buffer{}
	for s.Scan() {
		inst := NewInstruction(s.Text())
		// fmt.Println(inst.String())

		// Symbol handling
		symbol := inst.Symbol()
		if symbol != "" {
			// Check if it's in the symbol map
			if val, ok := SymbolMap[symbol]; ok {
				// fmt.Printf("Symbol %s\n", symbol)
				binBuf.WriteString(BinaryInt(val) + "\n")
				continue
			}

			// Check if it's a label
			if val, ok := p.Labels[symbol]; ok {
				// fmt.Printf("Label %s\n", symbol)
				binBuf.WriteString(Binary(val) + "\n")
				continue
			}

			// Variable handling
			addr := p.getVarAddress(symbol)
            binBuf.WriteString(BinaryInt(addr) + "\n")
			// fmt.Printf("Var %s\n", symbol)
			continue
		}

		// A instruction handling
		addr := inst.Address()
		if addr != -1 {
			binBuf.WriteString(fmt.Sprintf("%016b\n", addr))
			continue
		}

		// C instruction handling
		binBuf.WriteString(inst.BinaryC() + "\n")
	}

	// fmt.Println(p.Labels)
	// fmt.Println(p.Vars)

	return binBuf.Bytes(), nil
}

func (p *Parser) getVarAddress(name string) int64 {
    if addr, ok := p.Vars[name]; ok {
    	return addr
	}

    // First variable addressable space
    max := int64(16)

    // If we already declared vars
    for _, v := range p.Vars {
        if v >= max {
        	max = v + 1
		}
	}

    // fmt.Printf("max %s - %d\n", name, max)

    // If the current instruction has passed the maximum variable address take that instead
    if p.Current > max {
		// fmt.Printf("current %d, %s - %d\n", p.Current, name, max)
    	max = p.Current
	}

    p.Vars[name] = max

    return max
}
