package symboltable

import "errors"

type SymbolTable struct {
	table map[string]int
}

var predefinedSymbols = map[string]int{
	"R0": 0, "R1": 1, "R2": 2, "R3": 3, "R4": 4,
	"R5": 5, "R6": 6, "R7": 7, "R8": 8, "R9": 9,
	"R10": 10, "R11": 11, "R12": 12, "R13": 13,
	"R14": 14, "R15": 15,
	"SCREEN": 16384,
	"KBD":    24576,
	"SP":     0, "LCL": 1, "ARG": 2, "THIS": 3, "THAT": 4,
}

func NewSymbolTable() *SymbolTable {
	st := &SymbolTable{
		table: make(map[string]int),
	}

	for symbol, address := range predefinedSymbols {
		st.table[symbol] = address
	}

	return st
}

func (st *SymbolTable) AddEntry(symbol string, address int) error {
	if st.table == nil {
		return errors.New("SymbolTable has not been initialized or nil")
	}

	st.table[symbol] = address
	return nil
}

func (st *SymbolTable) GetAddress(symbol string) int {
	if val, ok := st.table[symbol]; ok {
		return val
	}
	return 0
}

func (st *SymbolTable) Contains(symbol string) bool {
	_, ok := st.table[symbol]
	return ok
}
