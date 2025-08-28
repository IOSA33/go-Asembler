package parser

import (
	"strings"
)

type CommandType int

const (
	A_COMMAND CommandType = iota // @value
	C_COMMAND                    // Dest=comp;jump
	L_COMMAND                    // (label)
)

type Parser struct {
	Lines        []string
	CurrentIndex int
	CurrentLine  []string
	CommandType  CommandType
	Symbol       string
	Dest         string
	Comp         string
	Jump         string
}

func NewParser(lines []string) *Parser {
	return &Parser{
		Lines:        lines,
		CurrentIndex: -1,
	}
}

func (p *Parser) HasMoreLines() bool {
	return p.CurrentIndex+1 < len(p.Lines)
}

func (p *Parser) Advance() {
	p.CurrentIndex++
	line := p.Lines[p.CurrentIndex]

	cleaned := cleanLine(line)
	if cleaned == "" {
			if p.HasMoreLines() {
				p.Advance()
			}
			return
	}

	// Giving for every word tokens
	p.CurrentLine = strings.Fields(cleaned)

	// Defying which command is it
	firstToken := p.CurrentLine[0]

	if strings.HasPrefix(firstToken, "@") {
		p.CommandType = A_COMMAND
		p.Symbol = strings.TrimPrefix(firstToken, "@")
		// If slice is more than one token
		if len(p.CurrentLine) > 1 {
			p.Symbol += strings.Join(p.CurrentLine[1:], "")
		}
		return
	}

	if p.CurrentLine[0] == "(" && p.CurrentLine[len(p.CurrentLine)-1] == ")" {
		p.CommandType = L_COMMAND
		p.Symbol = strings.TrimSuffix(strings.TrimPrefix(firstToken, "("), ")")
		return
	} else {
		// Logic for C_COMMAND
		p.CommandType = C_COMMAND

		p.Dest = ""
		p.Comp = ""
		p.Jump = ""

		fullCommand := strings.Join(p.CurrentLine, "")

		// Defying Index of "=" and ";"
		equalIndex := strings.Index(fullCommand, "=")
		semicolonIndex := strings.Index(fullCommand, ";")

		if equalIndex != -1 {
			p.Dest = fullCommand[:equalIndex]

			if semicolonIndex != -1 {
				// Dest=comp;jump
				p.Comp = fullCommand[equalIndex+1 : semicolonIndex]
				p.Jump = fullCommand[semicolonIndex+1:]
			} else {
				// Dest=comp
				p.Comp = fullCommand[equalIndex+1:]
			}
		} else if semicolonIndex != -1 {
			// comp;jump
			p.Comp = fullCommand[:semicolonIndex]
			p.Jump = fullCommand[semicolonIndex+1:]
		} else {
			// comp
			p.Comp = fullCommand
		}

	}
}

func cleanLine(line string) string {
	// deleting comment after "//"
	if idx := strings.Index(line, "//"); idx != -1 {
		line = line[:idx]
	}
	// Deleting spaces
	return strings.TrimSpace(line)
}
