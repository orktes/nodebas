//line parser.y:2
package interpreter

import __yyfmt__ "fmt"

//line parser.y:3
import (
//"fmt"
)

const GOTO = 57346
const GO = 57347
const TO = 57348
const SUB = 57349
const INTEGER = 57350
const FLOAT = 57351
const NEWLINE = 57352
const PRINT = 57353
const INPUT = 57354
const IF = 57355
const THEN = 57356
const FOR = 57357
const STEP = 57358
const NEXT = 57359
const REM = 57360
const DIM = 57361
const DEF = 57362
const WHITESPACE = 57363
const COLON = 57364
const COMMA = 57365
const DOT = 57366
const SEMICOLON = 57367
const LINESTART = 57368
const STRING_LITERAL = 57369
const BOOLEAN = 57370
const NULL = 57371
const IDENTIFIER = 57372
const RETURN = 57373
const END = 57374
const PLUS = 57375
const MINUS = 57376
const LET = 57377
const CIRCUMFLEX = 57378
const EQUALS = 57379
const LESS_THAN = 57380
const GREATER_THAN = 57381
const NOT_EQUALS = 57382
const LEFT_PARENTHESIS = 57383
const LEFT_BRACKET = 57384
const LEFT_BRACE = 57385
const RIGHT_BRACKET = 57386
const RIGHT_BRACE = 57387
const RIGHT_PARENTHESIS = 57388
const TIMES = 57389
const DIVIDE = 57390

var ParserToknames = []string{
	"GOTO",
	"GO",
	"TO",
	"SUB",
	"INTEGER",
	"FLOAT",
	"NEWLINE",
	"PRINT",
	"INPUT",
	"IF",
	"THEN",
	"FOR",
	"STEP",
	"NEXT",
	"REM",
	"DIM",
	"DEF",
	"WHITESPACE",
	"COLON",
	"COMMA",
	"DOT",
	"SEMICOLON",
	"LINESTART",
	"STRING_LITERAL",
	"BOOLEAN",
	"NULL",
	"IDENTIFIER",
	"RETURN",
	"END",
	"PLUS",
	"MINUS",
	"LET",
	"CIRCUMFLEX",
	"EQUALS",
	"LESS_THAN",
	"GREATER_THAN",
	"NOT_EQUALS",
	"LEFT_PARENTHESIS",
	"LEFT_BRACKET",
	"LEFT_BRACE",
	"RIGHT_BRACKET",
	"RIGHT_BRACE",
	"RIGHT_PARENTHESIS",
	"TIMES",
	"DIVIDE",
}
var ParserStatenames = []string{}

const ParserEofCode = 1
const ParserErrCode = 2
const ParserMaxDepth = 200

//line parser.y:459

//line yacctab:1
var ParserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const ParserNprod = 92
const ParserPrivate = 57344

var ParserTokenNames []string
var ParserStates []string

const ParserLast = 236

var ParserAct = []int{

	68, 54, 94, 147, 112, 101, 66, 102, 67, 101,
	149, 102, 136, 114, 101, 126, 102, 115, 135, 8,
	113, 105, 104, 37, 103, 95, 72, 74, 145, 77,
	78, 79, 141, 80, 90, 128, 63, 125, 122, 92,
	81, 82, 90, 85, 86, 127, 89, 108, 110, 111,
	109, 84, 91, 150, 89, 52, 83, 87, 88, 84,
	91, 90, 101, 97, 102, 98, 99, 43, 96, 106,
	85, 86, 44, 89, 64, 65, 100, 5, 84, 91,
	3, 70, 35, 6, 87, 88, 117, 118, 119, 120,
	121, 116, 123, 71, 36, 146, 1, 75, 69, 76,
	13, 93, 50, 51, 130, 131, 132, 133, 134, 129,
	48, 49, 47, 90, 138, 139, 46, 140, 45, 137,
	53, 90, 85, 86, 42, 89, 41, 40, 144, 143,
	84, 91, 39, 89, 107, 20, 87, 88, 84, 91,
	21, 73, 12, 90, 87, 88, 18, 148, 16, 17,
	151, 152, 85, 86, 11, 89, 15, 14, 10, 90,
	84, 91, 19, 142, 9, 7, 87, 88, 85, 86,
	4, 89, 2, 0, 0, 90, 84, 91, 0, 0,
	0, 124, 87, 88, 85, 86, 0, 89, 56, 57,
	0, 0, 84, 91, 0, 0, 0, 0, 87, 88,
	0, 0, 0, 0, 22, 23, 0, 55, 58, 59,
	63, 24, 0, 26, 62, 28, 0, 29, 33, 27,
	31, 38, 61, 60, 0, 0, 0, 0, 0, 0,
	34, 30, 32, 0, 0, 25,
}
var ParserPact = []int{

	69, -1000, 69, -1000, 200, -1000, -1000, 72, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 180, 68, 180, 180, 180, 6, 6, 6,
	-1000, 6, -1000, 13, 180, -1000, 200, 151, 180, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-5, 180, 180, -1000, 180, 180, 39, -1000, 151, -13,
	-15, -16, 151, 55, 10, -1000, -1000, -37, -17, -1000,
	-24, -1000, 39, -1000, 180, 180, 180, 180, 180, 180,
	8, 180, 135, -8, -1000, 23, -9, 97, 151, 151,
	180, -1000, -1000, 180, 180, 180, 180, 180, -1000, -1000,
	-19, -25, 180, 180, 180, 6, -14, 97, 97, 18,
	18, 18, -1000, 119, -1000, -1000, -5, 180, -1000, -1000,
	151, 151, 151, 151, 151, -1000, -1000, -18, 89, 151,
	-43, -1000, -1000, -1000, 151, -1000, 180, -27, 37, 180,
	180, 151, 151,
}
var ParserPgo = []int{

	0, 172, 80, 170, 19, 165, 164, 162, 158, 157,
	156, 154, 149, 148, 146, 142, 141, 140, 135, 134,
	6, 8, 72, 67, 132, 127, 126, 124, 0, 120,
	1, 118, 116, 112, 111, 110, 103, 102, 101, 2,
	100, 55, 99, 97, 96, 76,
}
var ParserR1 = []int{

	0, 44, 1, 1, 2, 3, 5, 5, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 17, 18, 7, 40, 14, 14, 6, 6, 13,
	12, 8, 11, 11, 11, 15, 9, 9, 10, 16,
	19, 19, 19, 19, 19, 19, 22, 41, 43, 42,
	24, 24, 25, 26, 27, 28, 28, 28, 28, 28,
	28, 28, 28, 28, 28, 28, 28, 28, 28, 28,
	28, 29, 30, 31, 32, 33, 35, 34, 36, 37,
	23, 23, 38, 38, 38, 39, 20, 20, 20, 45,
	45, 21,
}
var ParserR2 = []int{

	0, 1, 1, 2, 3, 1, 1, 3, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 1, 2, 4, 7, 2, 3, 3,
	1, 2, 4, 4, 4, 4, 6, 8, 2, 3,
	1, 1, 1, 2, 1, 2, 1, 4, 1, 4,
	3, 3, 3, 3, 3, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 1, 1, 1, 1, 1, 1, 3, 3,
	3, 4, 0, 1, 3, 3, 0, 1, 3, 1,
	1, 1,
}
var ParserChk = []int{

	-1000, -44, -1, -2, -3, 8, -2, -5, -4, -6,
	-8, -11, -15, -40, -9, -10, -13, -12, -14, -7,
	-18, -17, 4, 5, 11, 35, 13, 19, 15, 17,
	31, 20, 32, 18, 30, 10, 22, -28, 41, -24,
	-25, -26, -27, -23, -22, -31, -32, -33, -35, -34,
	-37, -36, -41, -29, -30, 27, 8, 9, 28, 29,
	43, 42, 34, 30, 6, 7, -20, -21, -28, -22,
	-41, -23, -28, -16, -28, -43, -42, -30, -30, -30,
	-30, 27, -20, -4, 41, 33, 34, 47, 48, 36,
	24, 42, -28, -38, -39, 30, -20, -28, -28, -28,
	-45, 23, 25, 37, 37, 37, 14, -19, 37, 40,
	38, 39, 41, 37, 37, 41, -20, -28, -28, -28,
	-28, -28, 30, -28, 46, 45, 23, 22, 44, -21,
	-28, -28, -28, -28, -28, 37, 37, -20, -28, -28,
	-30, 46, 44, -39, -28, 46, 6, 46, -28, 37,
	16, -28, -28,
}
var ParserDef = []int{

	0, -2, 1, 2, 0, 5, 3, 0, 6, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	19, 20, 0, 0, 86, 0, 0, 0, 0, 0,
	30, 0, 23, 0, 86, 4, 0, 27, 0, 56,
	57, 58, 59, 60, 61, 62, 63, 64, 65, 66,
	67, 68, 69, 70, 46, 73, 74, 75, 76, 77,
	82, 86, 0, 72, 0, 0, 31, 87, 91, 61,
	69, 60, 0, 0, 0, 24, 48, 0, 0, 38,
	0, 22, 21, 7, 86, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 83, 0, 0, 71, 28, 29,
	0, 89, 90, 0, 0, 0, 0, 0, 40, 41,
	42, 44, 86, 0, 0, 0, 0, 50, 51, 52,
	53, 54, 80, 0, 55, 79, 0, 0, 78, 88,
	32, 33, 34, 35, 39, 43, 45, 0, 0, 25,
	0, 47, 81, 84, 85, 49, 0, 0, 36, 0,
	0, 26, 37,
}
var ParserTok1 = []int{

	1,
}
var ParserTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48,
}
var ParserTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var ParserDebug = 0

type ParserLexer interface {
	Lex(lval *ParserSymType) int
	Error(s string)
}

const ParserFlag = -1000

func ParserTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(ParserToknames) {
		if ParserToknames[c-4] != "" {
			return ParserToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func ParserStatname(s int) string {
	if s >= 0 && s < len(ParserStatenames) {
		if ParserStatenames[s] != "" {
			return ParserStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func Parserlex1(lex ParserLexer, lval *ParserSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = ParserTok1[0]
		goto out
	}
	if char < len(ParserTok1) {
		c = ParserTok1[char]
		goto out
	}
	if char >= ParserPrivate {
		if char < ParserPrivate+len(ParserTok2) {
			c = ParserTok2[char-ParserPrivate]
			goto out
		}
	}
	for i := 0; i < len(ParserTok3); i += 2 {
		c = ParserTok3[i+0]
		if c == char {
			c = ParserTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = ParserTok2[1] /* unknown char */
	}
	if ParserDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", ParserTokname(c), uint(char))
	}
	return c
}

func ParserParse(Parserlex ParserLexer) int {
	var Parsern int
	var Parserlval ParserSymType
	var ParserVAL ParserSymType
	ParserS := make([]ParserSymType, ParserMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Parserstate := 0
	Parserchar := -1
	Parserp := -1
	goto Parserstack

ret0:
	return 0

ret1:
	return 1

Parserstack:
	/* put a state and value onto the stack */
	if ParserDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", ParserTokname(Parserchar), ParserStatname(Parserstate))
	}

	Parserp++
	if Parserp >= len(ParserS) {
		nyys := make([]ParserSymType, len(ParserS)*2)
		copy(nyys, ParserS)
		ParserS = nyys
	}
	ParserS[Parserp] = ParserVAL
	ParserS[Parserp].yys = Parserstate

Parsernewstate:
	Parsern = ParserPact[Parserstate]
	if Parsern <= ParserFlag {
		goto Parserdefault /* simple state */
	}
	if Parserchar < 0 {
		Parserchar = Parserlex1(Parserlex, &Parserlval)
	}
	Parsern += Parserchar
	if Parsern < 0 || Parsern >= ParserLast {
		goto Parserdefault
	}
	Parsern = ParserAct[Parsern]
	if ParserChk[Parsern] == Parserchar { /* valid shift */
		Parserchar = -1
		ParserVAL = Parserlval
		Parserstate = Parsern
		if Errflag > 0 {
			Errflag--
		}
		goto Parserstack
	}

Parserdefault:
	/* default state action */
	Parsern = ParserDef[Parserstate]
	if Parsern == -2 {
		if Parserchar < 0 {
			Parserchar = Parserlex1(Parserlex, &Parserlval)
		}

		/* look through exception table */
		xi := 0
		for {
			if ParserExca[xi+0] == -1 && ParserExca[xi+1] == Parserstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Parsern = ParserExca[xi+0]
			if Parsern < 0 || Parsern == Parserchar {
				break
			}
		}
		Parsern = ParserExca[xi+1]
		if Parsern < 0 {
			goto ret0
		}
	}
	if Parsern == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Parserlex.Error("syntax error")
			Nerrs++
			if ParserDebug >= 1 {
				__yyfmt__.Printf("%s", ParserStatname(Parserstate))
				__yyfmt__.Printf(" saw %s\n", ParserTokname(Parserchar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Parserp >= 0 {
				Parsern = ParserPact[ParserS[Parserp].yys] + ParserErrCode
				if Parsern >= 0 && Parsern < ParserLast {
					Parserstate = ParserAct[Parsern] /* simulate a shift of "error" */
					if ParserChk[Parserstate] == ParserErrCode {
						goto Parserstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if ParserDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", ParserS[Parserp].yys)
				}
				Parserp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if ParserDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", ParserTokname(Parserchar))
			}
			if Parserchar == ParserEofCode {
				goto ret1
			}
			Parserchar = -1
			goto Parsernewstate /* try again in the same state */
		}
	}

	/* reduction by production Parsern */
	if ParserDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Parsern, ParserStatname(Parserstate))
	}

	Parsernt := Parsern
	Parserpt := Parserp
	_ = Parserpt // guard against "declared and not used"

	Parserp -= ParserR2[Parsern]
	ParserVAL = ParserS[Parserp+1]

	/* consult goto table to find next state */
	Parsern = ParserR1[Parsern]
	Parserg := ParserPgo[Parsern]
	Parserj := Parserg + ParserS[Parserp].yys + 1

	if Parserj >= ParserLast {
		Parserstate = ParserAct[Parserg]
	} else {
		Parserstate = ParserAct[Parserj]
		if ParserChk[Parserstate] != -Parsern {
			Parserstate = ParserAct[Parserg]
		}
	}
	// dummy call; replaced with literal code
	switch Parsernt {

	case 1:
		//line parser.y:123
		{
			if v, ok := Parserlex.(*ParserLex); ok {
				v.cb(ParserS[Parserpt-0].item.(Lines))
			}
		}
	case 2:
		//line parser.y:130
		{
			ParserVAL.item = Lines{ParserS[Parserpt-0].item.(Line)}
		}
	case 3:
		//line parser.y:134
		{
			ParserVAL.item = append(ParserS[Parserpt-1].item.(Lines), ParserS[Parserpt-0].item.(Line))
		}
	case 4:
		//line parser.y:139
		{
			ParserVAL.item = Line{Number: ParserS[Parserpt-2].item.(LineNumber), Statements: ParserS[Parserpt-1].item.(Statements)}
		}
	case 5:
		//line parser.y:144
		{
			ParserVAL.item = LineNumber(ParserS[Parserpt-0].item.(int64))
		}
	case 6:
		//line parser.y:149
		{
			ParserVAL.item = Statements{ParserS[Parserpt-0].item.(Statement)}
		}
	case 7:
		//line parser.y:153
		{
			ParserVAL.item = append(ParserS[Parserpt-2].item.(Statements), ParserS[Parserpt-0].item.(Statement))
		}
	case 8:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 9:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 10:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 11:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 12:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 13:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 14:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 15:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 16:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 17:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 18:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 19:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 20:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 21:
		//line parser.y:173
		{
			ParserVAL.item = &OperationStatement{Operation: ParserS[Parserpt-1].item.(string), Args: ParserS[Parserpt-0].item.(Args)}
		}
	case 22:
		//line parser.y:178
		{
			ParserVAL.item = &RemarkStatement{}
		}
	case 23:
		//line parser.y:183
		{
			ParserVAL.item = &EndStatement{}
		}
	case 24:
		//line parser.y:188
		{
			ParserVAL.item = &DimStatement{ArrayDefinitions: ParserS[Parserpt-0].item.([]*ArrayDefinition)}
		}
	case 25:
		//line parser.y:193
		{
			ParserVAL.item = &DefStatement{
				Identifier: ParserS[Parserpt-2].item.(*IdentifierSymbol),
				Expression: ParserS[Parserpt-0].item.(Expression),
			}
		}
	case 26:
		//line parser.y:200
		{

			parameter := ParserS[Parserpt-3].item.(*IdentifierSymbol)
			ParserVAL.item = &DefStatement{
				Identifier: ParserS[Parserpt-5].item.(*IdentifierSymbol),
				Parameter:  parameter.Name,
				Expression: ParserS[Parserpt-0].item.(Expression),
			}
		}
	case 27:
		//line parser.y:211
		{
			ParserVAL.item = &GotoStatement{LineNumber: ParserS[Parserpt-0].item.(Expression)}
		}
	case 28:
		//line parser.y:215
		{
			ParserVAL.item = &GotoStatement{LineNumber: ParserS[Parserpt-0].item.(Expression)}
		}
	case 29:
		//line parser.y:220
		{
			ParserVAL.item = &GoSubStatement{LineNumber: ParserS[Parserpt-0].item.(Expression)}
		}
	case 30:
		//line parser.y:225
		{
			ParserVAL.item = &ReturnStatement{}
		}
	case 31:
		//line parser.y:230
		{
			ParserVAL.item = &PrintStatement{Args: ParserS[Parserpt-0].item.(Args)}
		}
	case 32:
		//line parser.y:235
		{
			ParserVAL.item = &AssignmentStatement{Symbol: ParserS[Parserpt-2].item.(Symbol), Expression: ParserS[Parserpt-0].item.(Expression)}
		}
	case 33:
		//line parser.y:239
		{
			ParserVAL.item = &AssignmentStatement{Symbol: ParserS[Parserpt-2].item.(Symbol), Expression: ParserS[Parserpt-0].item.(Expression)}
		}
	case 34:
		//line parser.y:243
		{
			ParserVAL.item = &AssignmentStatement{Symbol: ParserS[Parserpt-2].item.(Symbol), Expression: ParserS[Parserpt-0].item.(Expression)}
		}
	case 35:
		//line parser.y:248
		{
			ParserVAL.item = &IfStatement{RelationExpression: ParserS[Parserpt-2].item.(*RelationExpression), LineNumber: ParserS[Parserpt-0].item.(Expression)}
		}
	case 36:
		//line parser.y:253
		{
			ParserVAL.item = &ForStatement{
				ControlVariable: ParserS[Parserpt-4].item.(*IdentifierSymbol),
				InitialValue:    ParserS[Parserpt-2].item.(Expression),
				Limit:           ParserS[Parserpt-0].item.(Expression),
			}
		}
	case 37:
		//line parser.y:261
		{
			ParserVAL.item = &ForStatement{
				ControlVariable: ParserS[Parserpt-6].item.(*IdentifierSymbol),
				InitialValue:    ParserS[Parserpt-4].item.(Expression),
				Limit:           ParserS[Parserpt-2].item.(Expression),
				Increment:       ParserS[Parserpt-0].item.(Expression),
			}
		}
	case 38:
		//line parser.y:271
		{
			ParserVAL.item = &NextStatement{
				ControlVariable: ParserS[Parserpt-0].item.(*IdentifierSymbol),
			}
		}
	case 39:
		//line parser.y:278
		{
			ParserVAL.item = &RelationExpression{Left: ParserS[Parserpt-2].item.(Expression), Comparator: ParserS[Parserpt-1].item.(Comparator), Right: ParserS[Parserpt-0].item.(Expression)}
		}
	case 40:
		//line parser.y:283
		{
			ParserVAL.item = &EqualsComparator{}
		}
	case 41:
		//line parser.y:287
		{
			ParserVAL.item = &EqualsComparator{Negation: true}
		}
	case 42:
		//line parser.y:291
		{
			ParserVAL.item = &LessThanComparator{}
		}
	case 43:
		//line parser.y:295
		{
			ParserVAL.item = &LessThanComparator{OrEqual: true}
		}
	case 44:
		//line parser.y:299
		{
			ParserVAL.item = &GreaterThanComparator{}
		}
	case 45:
		//line parser.y:303
		{
			ParserVAL.item = &GreaterThanComparator{OrEqual: true}
		}
	case 46:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 47:
		//line parser.y:310
		{
			ParserVAL.item = &FunctionCall{Expression: ParserS[Parserpt-3].item.(Expression), Args: ParserS[Parserpt-1].item.(Args)}
		}
	case 48:
		//line parser.y:315
		{
			ParserVAL.item = []*ArrayDefinition{ParserS[Parserpt-0].item.(*ArrayDefinition)}
		}
	case 49:
		//line parser.y:320
		{
			ParserVAL.item = &ArrayDefinition{Symbol: ParserS[Parserpt-3].item.(*IdentifierSymbol), Dimensions: ParserS[Parserpt-1].item.(Args)}
		}
	case 50:
		//line parser.y:325
		{
			ParserVAL.item = &AdditiveExpression{Left: ParserS[Parserpt-2].item.(Expression), Minus: false, Right: ParserS[Parserpt-0].item.(Expression)}
		}
	case 51:
		//line parser.y:329
		{
			ParserVAL.item = &AdditiveExpression{Left: ParserS[Parserpt-2].item.(Expression), Minus: true, Right: ParserS[Parserpt-0].item.(Expression)}
		}
	case 52:
		//line parser.y:334
		{
			ParserVAL.item = &MultiplicationExpression{Left: ParserS[Parserpt-2].item.(Expression), Right: ParserS[Parserpt-0].item.(Expression)}
		}
	case 53:
		//line parser.y:339
		{
			ParserVAL.item = &DivisionExpression{Left: ParserS[Parserpt-2].item.(Expression), Right: ParserS[Parserpt-0].item.(Expression)}
		}
	case 54:
		//line parser.y:344
		{
			ParserVAL.item = &PowerExpression{BaseExpression: ParserS[Parserpt-2].item.(Expression), ExponentExpression: ParserS[Parserpt-0].item.(Expression)}
		}
	case 55:
		//line parser.y:349
		{
			ParserVAL.item = ParserS[Parserpt-1].item
		}
	case 56:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 57:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 58:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 59:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 60:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 61:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 62:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 63:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 64:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 65:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 66:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 67:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 68:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 69:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 70:
		ParserVAL.item = ParserS[Parserpt-0].item
	case 71:
		//line parser.y:369
		{
			ParserVAL.item = &UnaryMinusExpression{Expression: ParserS[Parserpt-0].item.(Expression)}
		}
	case 72:
		//line parser.y:374
		{
			ParserVAL.item = &IdentifierSymbol{Name: ParserS[Parserpt-0].item.(string)}
		}
	case 73:
		//line parser.y:379
		{
			str := ParserS[Parserpt-0].item.(string)
			ParserVAL.item = &StringLiteral{Value: str[1 : len(str)-1]}
		}
	case 74:
		//line parser.y:385
		{
			ParserVAL.item = &Integer{Value: ParserS[Parserpt-0].item.(int64)}
		}
	case 75:
		//line parser.y:390
		{
			ParserVAL.item = &Float{Value: ParserS[Parserpt-0].item.(float64)}
		}
	case 76:
		//line parser.y:395
		{
			ParserVAL.item = &Boolean{Value: ParserS[Parserpt-0].item.(bool)}
		}
	case 77:
		//line parser.y:400
		{
			ParserVAL.item = &Null{}
		}
	case 78:
		//line parser.y:405
		{
			ParserVAL.item = &Array{Items: ParserS[Parserpt-1].item.(Args)}
		}
	case 79:
		//line parser.y:410
		{
			ParserVAL.item = &Map{Properties: ParserS[Parserpt-1].item.([]*MapProperty)}
		}
	case 80:
		//line parser.y:415
		{
			ParserVAL.item = &MapPropertyExpression{Map: ParserS[Parserpt-2].item.(Expression), Property: &StringLiteral{Value: ParserS[Parserpt-0].item.(string)}}
		}
	case 81:
		//line parser.y:419
		{
			ParserVAL.item = &MapPropertyExpression{Map: ParserS[Parserpt-3].item.(Expression), Property: ParserS[Parserpt-1].item.(Expression)}
		}
	case 82:
		//line parser.y:424
		{
			ParserVAL.item = []*MapProperty{}
		}
	case 83:
		//line parser.y:428
		{
			ParserVAL.item = []*MapProperty{ParserS[Parserpt-0].item.(*MapProperty)}
		}
	case 84:
		//line parser.y:432
		{
			ParserVAL.item = append(ParserS[Parserpt-2].item.([]*MapProperty), ParserS[Parserpt-0].item.(*MapProperty))
		}
	case 85:
		//line parser.y:437
		{
			ParserVAL.item = &MapProperty{Name: ParserS[Parserpt-2].item.(string), Value: ParserS[Parserpt-0].item.(Expression)}
		}
	case 86:
		//line parser.y:442
		{
			ParserVAL.item = Args{}
		}
	case 87:
		//line parser.y:446
		{
			ParserVAL.item = Args{ParserS[Parserpt-0].item.(Expression)}
		}
	case 88:
		//line parser.y:450
		{
			ParserVAL.item = append(ParserS[Parserpt-2].item.(Args), ParserS[Parserpt-0].item.(Expression))
		}
	case 91:
		ParserVAL.item = ParserS[Parserpt-0].item
	}
	goto Parserstack /* stack new state and value */
}
