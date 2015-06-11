package interpreter

//import ("fmt")

func (lexer *ParserLex) Lex(lval *ParserSymType) (ret int) {
	ret = lexer.dLex(lval)
	/*	i := ret - GOTO
		if i >= 0 && i < len(ParserToknames) {
			fmt.Printf("returning %d %s\n", ret, ParserToknames[i])
		}
	*/
	return
}

func (lexer *ParserLex) dLex(lval *ParserSymType) (ret int) {

	//var err error
	var c byte = ' '

	defer func() {
		lexer.buf = nil
		lexer.tokenLen = 0
		lexer.back()
	}()

yystate0:

	goto yystart1

	goto yystate0 // silence unused label error
	goto yystate1 // silence unused label error
yystate1:
	c = lexer.next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate4
	case c == '$' || c >= 'A' && c <= 'C' || c == 'H' || c == 'J' || c == 'K' || c == 'M' || c == 'O' || c == 'Q' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'm' || c >= 'o' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate8
	case c == '(':
		goto yystate9
	case c == ')':
		goto yystate10
	case c == '*':
		goto yystate11
	case c == '+':
		goto yystate12
	case c == ',':
		goto yystate13
	case c == '-':
		goto yystate14
	case c == '.':
		goto yystate15
	case c == '/':
		goto yystate16
	case c == ':':
		goto yystate20
	case c == ';':
		goto yystate21
	case c == '<':
		goto yystate22
	case c == '=':
		goto yystate24
	case c == '>':
		goto yystate25
	case c == 'D':
		goto yystate26
	case c == 'E':
		goto yystate31
	case c == 'F':
		goto yystate34
	case c == 'G':
		goto yystate37
	case c == 'I':
		goto yystate41
	case c == 'L':
		goto yystate47
	case c == 'N':
		goto yystate50
	case c == 'P':
		goto yystate54
	case c == 'R':
		goto yystate59
	case c == 'S':
		goto yystate66
	case c == 'T':
		goto yystate74
	case c == '[':
		goto yystate79
	case c == '\n':
		goto yystate3
	case c == '\t' || c == ' ':
		goto yystate2
	case c == ']':
		goto yystate80
	case c == '^':
		goto yystate81
	case c == 'f':
		goto yystate82
	case c == 'n':
		goto yystate87
	case c == 't':
		goto yystate91
	case c == '{':
		goto yystate93
	case c == '}':
		goto yystate94
	case c >= '0' && c <= '9':
		goto yystate17
	}

yystate2:
	c = lexer.next()
	goto yyrule1

yystate3:
	c = lexer.next()
	switch {
	default:
		goto yyrule6
	case c == '\n':
		goto yystate3
	}

yystate4:
	c = lexer.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate5
	case c == '\\':
		goto yystate6
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate4
	}

yystate5:
	c = lexer.next()
	goto yyrule44

yystate6:
	c = lexer.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate7
	case c == '\\':
		goto yystate6
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate4
	}

yystate7:
	c = lexer.next()
	switch {
	default:
		goto yyrule44
	case c == '"':
		goto yystate5
	case c == '\\':
		goto yystate6
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate4
	}

yystate8:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate9:
	c = lexer.next()
	goto yyrule29

yystate10:
	c = lexer.next()
	goto yyrule30

yystate11:
	c = lexer.next()
	goto yyrule37

yystate12:
	c = lexer.next()
	goto yyrule35

yystate13:
	c = lexer.next()
	goto yyrule27

yystate14:
	c = lexer.next()
	goto yyrule36

yystate15:
	c = lexer.next()
	goto yyrule26

yystate16:
	c = lexer.next()
	goto yyrule38

yystate17:
	c = lexer.next()
	switch {
	default:
		goto yyrule42
	case c == '.':
		goto yystate18
	case c >= '0' && c <= '9':
		goto yystate17
	}

yystate18:
	c = lexer.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate19
	}

yystate19:
	c = lexer.next()
	switch {
	default:
		goto yyrule43
	case c >= '0' && c <= '9':
		goto yystate19
	}

yystate20:
	c = lexer.next()
	goto yyrule25

yystate21:
	c = lexer.next()
	goto yyrule28

yystate22:
	c = lexer.next()
	switch {
	default:
		goto yyrule4
	case c == '>':
		goto yystate23
	}

yystate23:
	c = lexer.next()
	goto yyrule2

yystate24:
	c = lexer.next()
	goto yyrule5

yystate25:
	c = lexer.next()
	goto yyrule3

yystate26:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'E':
		goto yystate27
	case c == 'I':
		goto yystate29
	}

yystate27:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'F':
		goto yystate28
	}

yystate28:
	c = lexer.next()
	switch {
	default:
		goto yyrule23
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate29:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'M':
		goto yystate30
	}

yystate30:
	c = lexer.next()
	switch {
	default:
		goto yyrule15
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate31:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'N':
		goto yystate32
	}

yystate32:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'D':
		goto yystate33
	}

yystate33:
	c = lexer.next()
	switch {
	default:
		goto yyrule16
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate34:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'O':
		goto yystate35
	}

yystate35:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'R':
		goto yystate36
	}

yystate36:
	c = lexer.next()
	switch {
	default:
		goto yyrule19
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate37:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'O':
		goto yystate38
	}

yystate38:
	c = lexer.next()
	switch {
	default:
		goto yyrule8
	case c == '$' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'T':
		goto yystate39
	}

yystate39:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'O':
		goto yystate40
	}

yystate40:
	c = lexer.next()
	switch {
	default:
		goto yyrule7
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate41:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'F':
		goto yystate42
	case c == 'N':
		goto yystate43
	}

yystate42:
	c = lexer.next()
	switch {
	default:
		goto yyrule14
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate43:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'P':
		goto yystate44
	}

yystate44:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'U':
		goto yystate45
	}

yystate45:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'T':
		goto yystate46
	}

yystate46:
	c = lexer.next()
	switch {
	default:
		goto yyrule13
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate47:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'E':
		goto yystate48
	}

yystate48:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'T':
		goto yystate49
	}

yystate49:
	c = lexer.next()
	switch {
	default:
		goto yyrule21
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate50:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'E':
		goto yystate51
	}

yystate51:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'X':
		goto yystate52
	}

yystate52:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'T':
		goto yystate53
	}

yystate53:
	c = lexer.next()
	switch {
	default:
		goto yyrule20
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate54:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'R':
		goto yystate55
	}

yystate55:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'I':
		goto yystate56
	}

yystate56:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'N':
		goto yystate57
	}

yystate57:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'T':
		goto yystate58
	}

yystate58:
	c = lexer.next()
	switch {
	default:
		goto yyrule12
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate59:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'E':
		goto yystate60
	}

yystate60:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'M':
		goto yystate61
	case c == 'T':
		goto yystate62
	}

yystate61:
	c = lexer.next()
	switch {
	default:
		goto yyrule24
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate62:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'U':
		goto yystate63
	}

yystate63:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'R':
		goto yystate64
	}

yystate64:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'N':
		goto yystate65
	}

yystate65:
	c = lexer.next()
	switch {
	default:
		goto yyrule18
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate66:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'S' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'T':
		goto yystate67
	case c == 'U':
		goto yystate72
	}

yystate67:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'E':
		goto yystate68
	case c == 'O':
		goto yystate70
	}

yystate68:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'P':
		goto yystate69
	}

yystate69:
	c = lexer.next()
	switch {
	default:
		goto yyrule22
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate70:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'P':
		goto yystate71
	}

yystate71:
	c = lexer.next()
	switch {
	default:
		goto yyrule17
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate72:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'B':
		goto yystate73
	}

yystate73:
	c = lexer.next()
	switch {
	default:
		goto yyrule10
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate74:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'H':
		goto yystate75
	case c == 'O':
		goto yystate78
	}

yystate75:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'E':
		goto yystate76
	}

yystate76:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	case c == 'N':
		goto yystate77
	}

yystate77:
	c = lexer.next()
	switch {
	default:
		goto yyrule11
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate78:
	c = lexer.next()
	switch {
	default:
		goto yyrule9
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate79:
	c = lexer.next()
	goto yyrule31

yystate80:
	c = lexer.next()
	goto yyrule32

yystate81:
	c = lexer.next()
	goto yyrule39

yystate82:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate8
	case c == 'a':
		goto yystate83
	}

yystate83:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate8
	case c == 'l':
		goto yystate84
	}

yystate84:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate8
	case c == 's':
		goto yystate85
	}

yystate85:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate8
	case c == 'e':
		goto yystate86
	}

yystate86:
	c = lexer.next()
	switch {
	default:
		goto yyrule41
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate87:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate8
	case c == 'u':
		goto yystate88
	}

yystate88:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate8
	case c == 'l':
		goto yystate89
	}

yystate89:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate8
	case c == 'l':
		goto yystate90
	}

yystate90:
	c = lexer.next()
	switch {
	default:
		goto yyrule40
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate8
	}

yystate91:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate8
	case c == 'r':
		goto yystate92
	}

yystate92:
	c = lexer.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate8
	case c == 'u':
		goto yystate85
	}

yystate93:
	c = lexer.next()
	goto yyrule33

yystate94:
	c = lexer.next()
	goto yyrule34

yyrule1: // [ \t]
	{
		lexer.tokenLen--
		if nil != lexer.buf {
			lexer.buf = lexer.buf[len(lexer.buf)-1:]
		}
		goto yystate0
	}
yyrule2: // "<>"
	{
		return NOT_EQUALS
	}
yyrule3: // ">"
	{
		return GREATER_THAN
	}
yyrule4: // "<"
	{
		return LESS_THAN
	}
yyrule5: // "="
	{
		return EQUALS
	}
yyrule6: // [\n]+
	{
		return NEWLINE
	}
yyrule7: // "GOTO"
	{
		return GOTO
	}
yyrule8: // "GO"
	{
		return GO
	}
yyrule9: // "TO"
	{
		return TO
	}
yyrule10: // "SUB"
	{
		return SUB
	}
yyrule11: // "THEN"
	{
		return THEN
	}
yyrule12: // "PRINT"
	{
		return PRINT
	}
yyrule13: // "INPUT"
	{
		return INPUT
	}
yyrule14: // "IF"
	{
		return IF
	}
yyrule15: // "DIM"
	{
		return DIM
	}
yyrule16: // "END"
	{
		return END
	}
yyrule17: // "STOP"
	{
		return END
	}
yyrule18: // "RETURN"
	{
		return RETURN
	}
yyrule19: // "FOR"
	{
		return FOR
	}
yyrule20: // "NEXT"
	{
		return NEXT
	}
yyrule21: // "LET"
	{
		return LET
	}
yyrule22: // "STEP"
	{
		return STEP
	}
yyrule23: // "DEF"
	{
		return DEF
	}
yyrule24: // "REM"
	{
		return REM
	}
yyrule25: // ":"
	{
		return COLON
	}
yyrule26: // "."
	{
		return DOT
	}
yyrule27: // ","
	{
		return COMMA
	}
yyrule28: // ";"
	{
		return SEMICOLON
	}
yyrule29: // "("
	{
		return LEFT_PARENTHESIS
	}
yyrule30: // ")"
	{
		return RIGHT_PARENTHESIS
	}
yyrule31: // "["
	{
		return LEFT_BRACKET
	}
yyrule32: // "]"
	{
		return RIGHT_BRACKET
	}
yyrule33: // "{"
	{
		return LEFT_BRACE
	}
yyrule34: // "}"
	{
		return RIGHT_BRACE
	}
yyrule35: // "+"
	{
		return PLUS
	}
yyrule36: // "-"
	{
		return MINUS
	}
yyrule37: // "*"
	{
		return TIMES
	}
yyrule38: // "/"
	{
		return DIVIDE
	}
yyrule39: // "^"
	{
		return CIRCUMFLEX
	}
yyrule40: // "null"
	{
		lval.item = nil
		return NULL
		goto yystate0
	}
yyrule41: // "true"|"false"
	{
		lval.item = lexer.getString() == "true"
		return BOOLEAN
		goto yystate0
	}
yyrule42: // {integer}
	{
		lval.item, _ = lexer.getInt()
		return INTEGER
		goto yystate0
	}
yyrule43: // {float}
	{
		lval.item, _ = lexer.getFloat()
		return FLOAT
		goto yystate0
	}
yyrule44: // \"(\\.|[^"])*\"
	{
		lval.item = lexer.getString()
		return STRING_LITERAL
		goto yystate0
	}
yyrule45: // [$_a-zA-Z]+
	{
		lval.item = lexer.getString()
		return IDENTIFIER
		goto yystate0
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized

	return -1

}
