package main

import (
	"assembler/code"
	"assembler/parser"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/main.go hello.asm")
		return
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %d", err)
		return
	}

	lines := strings.Split(string(file), "\n")

	parser_test := parser.NewParser(lines)

	for parser_test.HasMoreLines() {
		parser_test.Advance()

		dest := parser_test.Dest
		comp := parser_test.Comp
		jump := parser_test.Jump

		destCode := code.Dest(dest)
		compCode := code.Comp(comp)
		jumpCode := code.Jump(jump)

		fmt.Printf("Input line: %s\n", parser_test.CurrentLine)
		fmt.Printf("CommandType: %d\n", parser_test.CommandType)

		switch parser_test.CommandType {
		case parser.A_COMMAND, parser.L_COMMAND:
			fmt.Printf("Symbol: %s\n", parser_test.Symbol)
		case parser.C_COMMAND:
			fmt.Printf("Dest: '%s', Comp: '%s', Jump: '%s'\n", parser_test.Dest, parser_test.Comp, parser_test.Jump)
			fmt.Printf("CompCode: %s\n", compCode)
			fmt.Printf("DestCode: %s\n", destCode)
			fmt.Printf("JumpCode: %s\n", jumpCode)
			binaryCommand := "111" + compCode + destCode + jumpCode
			fmt.Printf("Binary command: %s\n", binaryCommand)
		}
		fmt.Println("-----")
	}

}
