package interpreter

import (
	"errors"
	"strconv"
)

type ParserLex struct {
	S        string
	pos      int
	buf      []byte
	tokenLen int
	cb       func(lines Lines)
	err      func(err error)
}

func (lexer *ParserLex) peek() (bret byte) {
	bret = lexer.next()
	lexer.pos--
	return
}

func (lexer *ParserLex) back() {
	if 0 < lexer.pos {
		lexer.pos--
	}
	return
}

func (lexer *ParserLex) next() (bret byte) {

	if lexer.pos < len(lexer.S) {
		bret = byte(lexer.S[lexer.pos])
		lexer.buf = append(lexer.buf, bret)
	} else {
		bret = 0
	}
	lexer.pos++
	lexer.tokenLen++

	return
}

func (lexer *ParserLex) data() (bb []byte) {
	if lexer.pos < len(lexer.S) {
		bb = lexer.buf[:len(lexer.buf)-1]
	} else {
		bb = lexer.buf
	}

	if len(bb) > lexer.tokenLen {
		bb = bb[:lexer.tokenLen]
	}
	return
}

func (lexer *ParserLex) Error(s string) {
	lexer.err(errors.New(s))
}

func (lexer *ParserLex) getInt() (n int64, err error) {
	s := string(lexer.data())
	if n, err = strconv.ParseInt(s, 10, 64); nil != err {
		lexer.Error(err.Error() + ";s=" + s)
	}
	return
}

func (lexer *ParserLex) getFloat() (n float64, err error) {
	s := string(lexer.data())
	if n, err = strconv.ParseFloat(s, 64); nil != err {
		lexer.Error(err.Error() + ";s=" + s)
	}
	return
}

func (lexer *ParserLex) getString() (s string) {
	s = string(lexer.data())
	return
}
