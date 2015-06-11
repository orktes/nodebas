%{

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

	defer func(){
		lexer.buf = nil
		lexer.tokenLen = 0
		lexer.back()
	}()
%}

%yyc c
%yyn c = lexer.next()

digit [0-9]
integer {digit}+
float {integer}"."{integer}

%%

[ \t]           		lexer.tokenLen--;if nil!=lexer.buf {
                    lexer.buf = lexer.buf[len(lexer.buf)-1:]
		              }

"<>"					return NOT_EQUALS;
">"						return GREATER_THAN;
"<"						return LESS_THAN;
"="						return EQUALS;
[\n]+   		  return NEWLINE;
"GOTO"        return GOTO;
"GO"          return GO;
"TO"					return TO;
"SUB"					return SUB;
"THEN"				return THEN;
"PRINT"       return PRINT;
"INPUT"       return INPUT;
"IF"          return IF;
"DIM"         return DIM;
"END"					return END;
"STOP"				return END;
"RETURN"			return RETURN;
"FOR"					return FOR;
"NEXT"				return NEXT;
"LET"					return LET;
"STEP"				return STEP;
"DEF"					return DEF;
"REM"					return REM;
":"           return COLON;
"."						return DOT;
","						return COMMA;
";"						return SEMICOLON;
"("						return LEFT_PARENTHESIS;
")"						return RIGHT_PARENTHESIS;
"["						return LEFT_BRACKET;
"]"						return RIGHT_BRACKET;
"{"						return LEFT_BRACE;
"}"						return RIGHT_BRACE;
"+"						return PLUS;
"-"						return MINUS;
"*"						return TIMES;
"/"						return DIVIDE;
"^"						return CIRCUMFLEX;

"null"				  lval.item = nil; return NULL;
"true"|"false"  lval.item = lexer.getString() == "true"; return BOOLEAN;
{integer}       lval.item, _ = lexer.getInt(); return INTEGER;
{float}         lval.item, _ = lexer.getFloat(); return FLOAT;
\"(\\.|[^"])*\" lval.item = lexer.getString(); return STRING_LITERAL;

[$_a-zA-Z]+     lval.item = lexer.getString(); return IDENTIFIER;


%%

return -1

}
