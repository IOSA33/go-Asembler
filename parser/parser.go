package parser

import (
	"strings"
)

type CommandType int

const (
	A_COMMAND CommandType = iota // @value
	C_COMMAND                    // dest=comp;jump
	L_COMMAND                    // (label)
)

type Parser struct {
	lines        []string
	currentIndex int
	currentLine  []string
	commandType  CommandType
	symbol       string
	dest         string
	comp         string
	jump         string
}

func NewParser(lines []string) *Parser {
	return &Parser{
		lines:        lines,
		currentIndex: -1,
	}
}

func (p *Parser) HasMoreLines() bool {
	return p.currentIndex+1 < len(p.lines)
}

func (p *Parser) Advance() {
	p.currentIndex++
	line := p.lines[p.currentIndex]

	cleaned := cleanLine(line)
	if cleaned == "" {
		if p.HasMoreLines() {
			p.Advance()
		}
		return
	}

	// Giving for every word tokens
	p.currentLine = strings.Fields(cleaned)

	// Defying which command is it
	firstToken := p.currentLine[0]

	if strings.HasPrefix(firstToken, "@") {
		p.commandType = A_COMMAND
		p.symbol = strings.TrimPrefix(firstToken, "@")
		// If slice is more than one token
		if len(p.currentLine) > 1 {
			p.symbol += strings.Join(p.currentLine[1:], "")
		}
		return
	}

	if p.currentLine[0] == "(" && p.currentLine[len(p.currentLine)-1] == ")" {
		p.commandType = L_COMMAND
		p.symbol = strings.TrimSuffix(strings.TrimPrefix(firstToken, "("), ")")
		return
	} else {
		// Logic for C_COMMAND
		p.commandType = C_COMMAND

		p.dest = ""
		p.comp = ""
		p.jump = ""

		fullCommand := strings.Join(p.currentLine, "")

		// Ищем разделители
		equalIndex := strings.Index(fullCommand, "=")
		semicolonIndex := strings.Index(fullCommand, ";")

		if equalIndex != -1 {
			p.dest = fullCommand[:equalIndex]

			if semicolonIndex != -1 {
				// dest=comp;jump
				p.comp = fullCommand[equalIndex+1 : semicolonIndex]
				p.jump = fullCommand[semicolonIndex+1:]
			} else {
				// dest=comp
				p.comp = fullCommand[equalIndex+1:]
			}
		} else if semicolonIndex != -1 {
			// comp;jump
			p.comp = fullCommand[:semicolonIndex]
			p.jump = fullCommand[semicolonIndex+1:]
		} else {
			// comp
			p.comp = fullCommand
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

func (p *Parser) CommandType() CommandType {
	return p.commandType
}

func (p *Parser) CurrentLine() []string {
	return p.currentLine
}

func (p *Parser) Symbol() string {
	return p.symbol
}

func (p *Parser) Dest() string {
	return p.dest
}

func (p *Parser) Comp() string {
	return p.comp
}

func (p *Parser) Jump() string {
	return p.jump
}
