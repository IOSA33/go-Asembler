package main

import (
	"assembler/code"
	"assembler/parser"
	"assembler/symboltable"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/main.go hello.asm")
		return
	}

	file, err := os.ReadFile(os.Args[1])
	// file, err := os.ReadFile("hello.asm")
	if err != nil {
		fmt.Printf("Error: %d", err)
		return
	}

	lines := strings.Split(string(file), "\n")

	parserTest := parser.NewParser(lines)
	symtable := symboltable.NewSymbolTable()

	fileName := os.Args[1]

	f, err := os.Create(fileName + ".hack")
	//f, err := os.Create("hello.hack")
	if err != nil {
		fmt.Printf("Error: %d", err)
	}
	defer f.Close()

	programCounter := 0

	// First Pass
	for parserTest.HasMoreLines() {
		parserTest.Advance()

		switch parserTest.CommandType {
		case parser.A_COMMAND, parser.C_COMMAND:
			programCounter++
		case parser.L_COMMAND:
			fmt.Printf("COMMAND: %v\n", parserTest.CommandType)
			_, err := symtable.AddEntry(parserTest.Symbol, programCounter)
			check(err)
		}
	}

	// Creating NewParser for second Pass
	parserTest2 := parser.NewParser(lines)
	programCounterA := 16

	// Second Pass
	for parserTest2.HasMoreLines() {
		parserTest2.Advance()

		dest := parserTest2.Dest
		comp := parserTest2.Comp
		jump := parserTest2.Jump

		destCode := code.Dest(dest)
		compCode := code.Comp(comp)
		jumpCode := code.Jump(jump)

		fmt.Printf("Input line: %s\n", parserTest2.CurrentLine)

		switch parserTest2.CommandType {
		case parser.A_COMMAND:
			fmt.Printf("Symbol: %s\n", parserTest2.Symbol)
			// Checks if it is an int
			if typeInt, err := strconv.Atoi(parserTest2.Symbol); err == nil {
				fmt.Println("---------- Converted")
				binaryCommandA := fmt.Sprintf("%016b", typeInt)
				fmt.Println(binaryCommandA)
				_, err := f.WriteString(binaryCommandA + "\n")
				check(err)
				fmt.Println("---------- Converted")
			} else {
				// If it contains in
				if ok := symtable.Contains(parserTest2.Symbol); ok {
					fmt.Println("---------- Converted 2")
					addressA := symtable.GetAddress(parserTest2.Symbol)
					fmt.Println(addressA)
					binaryCommandA1 := fmt.Sprintf("%016b", addressA)
					fmt.Println(binaryCommandA1)

					_, err := f.WriteString(binaryCommandA1 + "\n")
					check(err)

					fmt.Println("---------- Converted 2")
				} else {
					// If it doesn't contain in symbol table we add the entry to map
					fmt.Println("---------- Converted 3")
					valueAddress, err := symtable.AddEntry(parserTest2.Symbol, programCounterA)
					check(err)

					binaryCommandA2 := fmt.Sprintf("%016b", valueAddress)
					fmt.Println(binaryCommandA2)
					_, err = f.WriteString(binaryCommandA2 + "\n")
					check(err)

					programCounterA++
					fmt.Println("---------- Converted 3")
				}

			}
		case parser.C_COMMAND:
			fmt.Printf("Dest: '%s', Comp: '%s', Jump: '%s'\n", parserTest2.Dest, parserTest2.Comp, parserTest2.Jump)
			fmt.Printf("CompCode: %s\n", compCode)
			fmt.Printf("DestCode: %s\n", destCode)
			fmt.Printf("JumpCode: %s\n", jumpCode)

			binaryCommand := "111" + compCode + destCode + jumpCode

			fmt.Printf("Binary command: %s\n", binaryCommand)

			_, err := f.WriteString(binaryCommand + "\n")
			check(err)
		case parser.L_COMMAND:
			fmt.Println("L_COMMAND")
		}
		fmt.Println("------------------------------ ForLoop")
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
