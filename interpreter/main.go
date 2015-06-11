package interpreter

import (
	"sort"
	"strings"
)

//go:generate go tool yacc -p Parser -o parser.go parser.y
//go:generate $GOPATH/bin/golex -o lexer.go parser.l

func ParseBasicCode(code string) (lines Lines, err error) {
	// TODO extra \ln should be handle in the parser or lexer
	if !strings.HasSuffix(code, "\n") {
		code = code + "\n"
	}
	ParserParse(&ParserLex{
		S: code,
		cb: func(res Lines) {
			lines = res
		},
		err: func(errObj error) {
			err = errObj
		},
	})

	sort.Sort(lines)

	return
}

func EvalBasicCode(code string) error {
	lines, err := ParseBasicCode(code)
	if err != nil {
		return err
	}
	interpreter := NewIntepreter(lines)
	return interpreter.Run()
}
