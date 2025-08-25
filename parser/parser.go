package parser

import (
	"strings"
)

type CommandType int

const (
	A_COMMAND CommandType = iota // @value
	C_COMMAND
	L_COMMAND
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

func NewParse(lines []string) *Parser {
	return &Parser{
		lines:        lines,
		currentIndex: -1,
	}
}

func (p *Parser) hasMoreLines(lines []string) bool {
	return p.currentIndex+1 < len(lines)
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

		for i, word := range p.currentLine {
			equal := strings.Index(word, "=")
			doubleDot := strings.Index(word, ";")

			if equal == -1 {
				p.dest = ""
				p.jump = strings.Join(p.currentLine[doubleDot+1:], " ")
				p.comp = strings.Join(p.currentLine[:doubleDot], " ")
			}

			if word == "=" {
				p.dest = strings.Join(p.currentLine[:i], " ")
			}

			if doubleDot == -1 {
				p.comp = strings.Join(p.currentLine[equal+1:], " ")
				p.jump = ""
				return
			}

			if word == ";" {
				p.comp = strings.Join(p.currentLine[equal+1:i], " ")
				p.jump = strings.Join(p.currentLine[doubleDot+1:], " ")
			}
		}

	}
}

func (p *Parser) CommandType() CommandType {
	return p.commandType
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
