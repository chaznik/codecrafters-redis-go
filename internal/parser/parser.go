package parser

import (
	"bufio"
	"io"
)

const (
	RESPSimpleString = '+'
	RESPSimpleError  = '-'
	RESPInteger      = ':'
	RESPBulkString   = '$'
	RESPArray        = '*'
)

type Value struct {
	TypeOf byte
	Str    string
	Num    int
	Array  []Value
}

type Parser struct {
	r *bufio.Reader
}

func NewParser(r io.Reader) *Parser {
	return &Parser{r: bufio.NewReader(r)}
}
