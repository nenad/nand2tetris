// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)

// Put your code here.

@i
M=0

@R2
M=0

(LOOP)
    // Conditional check if we reached i=R1
    @i
    D=M
    @R0
    D=D-M
    @END
    D;JEQ

    // Increment i
    @i
    M=M+1

    // Get value of R1
    @R1
    D=M

    // Sum the second register
    @R2
    M=M+D

    // Repeat
    @LOOP
    0;JMP

(END)
    @END
    0;JMP
