package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var commentRegex = regexp.MustCompile(`^(.*)(//.*)$`)
var labelRegex = regexp.MustCompile(`^\(([_a-zA-Z0-9\.\$]+)\)$`)
var symbolRegex = regexp.MustCompile(`^@([a-zA-Z_][_a-zA-Z0-9\.\$]*)$`)
var addressRegex = regexp.MustCompile(`^@([0-9]+)$`)

type Instruction struct {
	Code string
}

func NewInstruction(code string) Instruction {
	// Do all checks to make sure it's a valid instruction
	inst := sanitize(code)

	return Instruction{Code:inst}
}

func (i *Instruction) Empty() bool {
	return i.Code == ""
}

func (i *Instruction) String() string {
	return i.Code
}

func (i *Instruction) IsA() bool {
	return i.Code[0] == '@'
}

func (i *Instruction) IsC() bool {
	// Quick check of first binary bit
	return i.Code[0] != '@'
}

func (i *Instruction) Address() int64 {
	if i.Code[0] == '@' {
		str := addressRegex.ReplaceAllString(i.Code, "$1")
		addr, _ := strconv.ParseInt(str, 10, 16)
		return addr
	}
	return -1
}

func (i *Instruction) Symbol() string {
	if !i.IsA() {
		return ""
	}

	if symbolRegex.MatchString(i.Code) {
		return symbolRegex.ReplaceAllString(i.Code, "$1")
	}

	return ""
}

func (i *Instruction) Label() string {
	// Quick check of first binary bit
	if i.Code[0] == '(' {
        return labelRegex.ReplaceAllString(i.Code, "$1")
	}
	return ""
}

func (i *Instruction) Destination() string {
    if strings.Contains(i.Code, "=") {
    	return strings.Split(i.Code, "=")[0]
	}
    return ""
}

func (i *Instruction) Operation() string {
	if strings.Contains(i.Code, "=") {
		return strings.Split(i.Code, "=")[1]
	}

	return strings.Split(i.Code, ";")[0]
}

func (i *Instruction) Jump() string {
	if strings.Contains(i.Code, ";") {
		return strings.Split(i.Code, ";")[1]
	}
	return ""
}

func (i *Instruction) BinaryC() string {
    if !i.IsC() {
        return "x"
	}

    return fmt.Sprintf("111%s%s%s", OperationMap[i.Operation()], DestinationMap[i.Destination()], JumpMap[i.Jump()])
}

// sanitize cleans up the line so that the parsing of the instruction is more precise
func sanitize(instruction string) string {
	// Strip all whitespace
	code := strings.Join(strings.Fields(instruction), "")

	// Clean comments
	code = commentRegex.ReplaceAllString(code, "$1")

	return code
}
