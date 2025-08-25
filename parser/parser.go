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
	if len(line) == 0 {
		return
	}

	p.currentLine = strings.Fields(line)
	for i, word := range p.currentLine {
		if word == "/" {
			p.currentLine = p.currentLine[:i]
		}
	}

	if p.currentLine[0] == "@" {
		p.commandType = A_COMMAND
		p.symbol = strings.Join(p.currentLine[1:], "")
	}

	if p.currentLine[0] == "(" && p.currentLine[len(p.currentLine)-1] == ")" {
		p.commandType = L_COMMAND
		p.symbol = strings.Join(p.currentLine[1:len(p.currentLine)-1], "")
	} else {
		p.commandType = C_COMMAND
		p.symbol = strings.Join(p.currentLine, " ")

		// Checks if in the slice we have "="
		hasEqual := false
		equalIndex := -1
		for i, token := range p.currentLine {
			if token == "=" {
				hasEqual = true
				equalIndex = i
				break
			}
		}

		// Checks if in the slice we have ";"
		hasSemicolon := false
		semicolonIndex := -1
		for i, token := range p.currentLine {
			if token == ";" {
				hasSemicolon = true
				semicolonIndex = i
				break
			}
		}

		// If slice has "="
		if hasEqual {
			p.dest = strings.Join(p.currentLine[:equalIndex], " ")

			// If slice has "=" and ";"
			if hasSemicolon {
				p.comp = strings.Join(p.currentLine[equalIndex+1:semicolonIndex], " ")
				p.jump = strings.Join(p.currentLine[semicolonIndex+1:], " ")
			} else {
				// If slice has "=" and dont ";"
				p.comp = strings.Join(p.currentLine[equalIndex+1:], " ")
				p.jump = ""
			}

		} else if hasSemicolon {
			// If slice doesnt have "=" but has ";"
			p.dest = ""
			p.comp = strings.Join(p.currentLine[:semicolonIndex], " ")
			p.jump = strings.Join(p.currentLine[semicolonIndex+1:], " ")
		} else {
			p.dest = ""
			p.comp = strings.Join(p.currentLine, " ")
			p.jump = ""
		}

	}
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
