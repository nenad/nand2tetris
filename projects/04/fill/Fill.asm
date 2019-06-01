// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed.
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.



// Check keyboard input
// If YES
//      if screen is not colored, color
//      else jump to LOOP
// If NO
//      if screen is not colored, jump to LOOP
//      else color screen

// @colored | @keyboard | result
//      0   |   0       |   jump to loop
//      0   |   1       |   color & set colored=1
//      1   |   0       |   remove color & set colored=0
//      1   |   1       |   jump to loop

// SCREEN is 2048 addressed wide

@colored
M=0
@pressed
M=0

(LOOP)

    @KBD
    D=M

    @NOT_PRESSED
    D;JEQ

    @KBD
    D=1

    (NOT_PRESSED)
    // Check if we should skip screen changing
    @colored
    D=D+M
    @LOOP
    D-1;JNE

    // Point @i to first screen address
    @SCREEN
    D=A
    @i
    M=D

    // Set end address
    // 8192 addresses for the screen
    @8192
    D=D+A
    @end
    M=D

    (COLOR)
        @colored
        D=M
        @BLACK
        D;JEQ
        @WHITE
        D;JGT

        (BLACK)
        @i
        D=M
        A=D
        M=-1
        @EXIT_COLOR
        0;JMP

        (WHITE)
        @i
        D=M
        A=D
        M=0
        @EXIT_COLOR
        0;JMP

        (EXIT_COLOR)
            // Should we color more or should we exit
            @i
            MD=M+1
            @end
            D=M-D
            @COLOR
            D;JNE

            @colored
            D=M

            @COLOR_TRUE
            D;JEQ
            @COLOR_FALSE
            D;JNE

            (COLOR_TRUE)
            @colored
            M=1
            @LOOP
            0;JMP

            (COLOR_FALSE)
            @colored
            M=0
            @LOOP
            0;JMP
