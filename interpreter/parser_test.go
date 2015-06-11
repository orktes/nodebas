package interpreter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParserGoto(t *testing.T) {
	lines, _ := ParseBasicCode("1002 GOTO 1001")
	if val, ok := lines[0].Statements[0].(*GotoStatement); ok {
		if val.LineNumber.GetValue(nil) != int64(1001) {
			t.Error("Wrong line number")
		}
	} else {
		t.Error("Wront statement type returned")
	}

	lines, _ = ParseBasicCode("1002 GO TO 1001")
	if val, ok := lines[0].Statements[0].(*GotoStatement); ok {
		if val.LineNumber.GetValue(nil) != int64(1001) {
			t.Error("Wrong line number")
		}
	} else {
		t.Error("Wront statement type returned")
	}

}

func TestParserPrint(t *testing.T) {
	lines, _ := ParseBasicCode("1001 PRINT \"foobar\"")
	if val, ok := lines[0].Statements[0].(*PrintStatement); ok {
		if val.Args[0].GetValue(nil) != "foobar" {
			t.Error("Wrong argument", val.Args[0])
		}
	} else {
		t.Error("Wront statement type returned")
	}
}

func TestParserPrintWithMultiArgs(t *testing.T) {
	lines, _ := ParseBasicCode("1001 PRINT \"foobar\", \"barfoo\", 123")
	if val, ok := lines[0].Statements[0].(*PrintStatement); ok {
		if val.Args[0].GetValue(nil) != "foobar" {
			t.Error("Wrong argument", val.Args[0].GetValue(nil))
		}
		if val.Args[1].GetValue(nil) != "barfoo" {
			t.Error("Wrong argument", val.Args[1].GetValue(nil))
		}
		if val.Args[2].GetValue(nil) != int64(123) {
			t.Error("Wrong argument", val.Args[2].GetValue(nil))
		}
	} else {
		t.Error("Wront statement type returned")
	}
}

func TestParserPrintWithAssignment(t *testing.T) {
	lines, _ := ParseBasicCode("1001 LET ab = ba")
	if val, ok := lines[0].Statements[0].(*AssignmentStatement); ok {
		if val.Expression.(*IdentifierSymbol).Name != "ba" {
			t.Error("Wrong expression")
		}
		if val.Symbol.(*IdentifierSymbol).Name != "ab" {
			t.Error("Wrong expression")
		}
	} else {
		fmt.Println(reflect.TypeOf(lines[0].Statements[0]))
		t.Error("Wront statement type returned")
	}
}

func TestParserIF(t *testing.T) {
	scope := NewScope()

	scope.Set("a", int64(1))
	scope.Set("b", int64(2))

	lines, _ := ParseBasicCode("1001 IF a = b THEN 1002\n1003 IF a > b THEN 1001\n1006 IF a >= b THEN 1001\n1007 IF a < b THEN 1001\n1008 IF a <= b THEN 1001\n1009 IF a <> b THEN 1001")

	results := []bool{false, false, false, true, true, true}

	for indx, line := range lines {
		stm := line.Statements[0]
		if val, ok := stm.(*IfStatement); ok {
			if val.RelationExpression.GetValue(scope) != results[indx] {
				t.Error("Wrong value", val.RelationExpression.GetValue(scope), indx)
			}
		} else {
			t.Error("Not if")
		}
	}
}

func TestParserDim(t *testing.T) {
	lines, _ := ParseBasicCode("1001 DIM foo(2)\n1002 LET foo(0) = 1\n1003 PRINT foo(0)")
	if val, ok := lines[0].Statements[0].(*DimStatement); ok {
		if val.ArrayDefinitions[0].Symbol.Name != "foo" {
			t.Error("Wrong identifier")
		}

		if val.ArrayDefinitions[0].Dimensions[0].GetValue(nil) != int64(2) {
			t.Error("Wrong size definition")
		}
	} else {
		t.Error("Wront statement type returned")
	}
}

func TestParserAdditiveExpression(t *testing.T) {
	lines, _ := ParseBasicCode("1001 LET a = 1 + 3")
	if val, ok := lines[0].Statements[0].(*AssignmentStatement); ok {
		if val, ok := val.Expression.(*AdditiveExpression); ok {
			if val.Left.GetValue(nil) != int64(1) {
				t.Error("Not the right value")
			}
			if val.Right.GetValue(nil) != int64(3) {
				t.Error("Not the right value")
			}
		}
	}
}

func TestParserPrintWithAdditiveExpression(t *testing.T) {
	lines, _ := ParseBasicCode("1001 PRINT 1+1")
	if val, ok := lines[0].Statements[0].(*PrintStatement); !ok {
		if _, ok := val.Args[0].(*AdditiveExpression); !ok {
			t.Error("Wront statement type returned")
		}
	}
}

func TestParserPrintWithMultiplicationExpression(t *testing.T) {
	lines, _ := ParseBasicCode("1001 PRINT 1*1")
	if val, ok := lines[0].Statements[0].(*PrintStatement); !ok {
		if _, ok := val.Args[0].(*MultiplicationExpression); !ok {
			t.Error("Wront statement type returned")
		}
	}
}

func TestParserPrintWithDivisionExpression(t *testing.T) {
	lines, _ := ParseBasicCode("1001 PRINT 3/1.5")
	if val, ok := lines[0].Statements[0].(*PrintStatement); !ok {
		if val, ok := val.Args[0].(*DivisionExpression); !ok {
			if _, ok := val.Right.(*Float); !ok {
				t.Error("Wront statement type returned")
			}
		}
	}
}

func TestParserPrintWithPowExpression(t *testing.T) {
	lines, _ := ParseBasicCode("1001 PRINT 2^2")
	if val, ok := lines[0].Statements[0].(*PrintStatement); !ok {
		if val, ok := val.Args[0].(*PowerExpression); !ok {
			if _, ok := val.BaseExpression.(*Integer); !ok {
				t.Error("Wront statement type returned")
			}
			if _, ok := val.ExponentExpression.(*Integer); !ok {
				t.Error("Wront statement type returned")
			}
		}
	}
}

func TestParserDimWithAdditiveExpression(t *testing.T) {
	lines, _ := ParseBasicCode("1001 DIM foo(1+1)")
	if val, ok := lines[0].Statements[0].(*DimStatement); ok {
		if val.ArrayDefinitions[0].Symbol.Name != "foo" {
			t.Error("Wrong identifier")
		}

		if val.ArrayDefinitions[0].Dimensions[0].GetValue(nil) != int64(2) {
			t.Error("Wrong size definition")
		}
	} else {
		t.Error("Wront statement type returned")
	}
}

func TestParserForStatement(t *testing.T) {
	lines, _ := ParseBasicCode("1000 FOR a = 1 TO 10\n1001 PRINT a\n1002 NEXT a")
	if val, ok := lines[0].Statements[0].(*ForStatement); ok {
		if val.ControlVariable.Name != "a" {
			t.Error("Invalid control variable")
		}
		if val.InitialValue.GetValue(nil) != int64(1) {
			t.Error("Invalid initial value")
		}
		if val.Limit.GetValue(nil) != int64(10) {
			t.Error("Invalid limit value")
		}
	} else {
		t.Error("Wront statement type returned")
	}
}

func TestParserGoSubReturnStatement(t *testing.T) {
	lines, _ := ParseBasicCode("1002 GO SUB 1001\n1003 RETURN")
	if val, ok := lines[0].Statements[0].(*GoSubStatement); ok {
		if val.LineNumber.GetValue(nil) != int64(1001) {
			t.Error("Wrong line number")
		}
	} else {
		t.Error("Wront statement type returned")
	}

	if _, ok := lines[1].Statements[0].(*ReturnStatement); !ok {
		t.Error("Wront statement type returned")
	}
}

func TestParserDEFStatement(t *testing.T) {
	lines, _ := ParseBasicCode("1000 DEF FNA = 1 + 1\n1001 DEF FNB(x) = x + 1")
	if val, ok := lines[0].Statements[0].(*DefStatement); ok {
		if val.Expression.GetValue(nil) != int64(2) {
			t.Error("Wrong func definition")
		}
	} else {
		t.Error("Wront statement type returned")
	}

	if val, ok := lines[1].Statements[0].(*DefStatement); ok {
		testScope := NewScope()
		testScope.Set("x", int64(2))
		if val.Expression.GetValue(testScope) != int64(3) {
			t.Error("Wrong func definition")
		}
		if val.Parameter != "x" {
			t.Error("Parameter name should be x")
		}
	} else {
		t.Error("Wront statement type returned")
	}
}

func TestParserFunctionCallStatement(t *testing.T) {
	lines, _ := ParseBasicCode("1000 LET a = FNF()")
	if val, ok := lines[0].Statements[0].(*AssignmentStatement); ok {
		if _, ok := val.Expression.(*FunctionCall); !ok {
			t.Error("Should be a function call")
		}
	} else {
		t.Error("Wront statement type returned")
	}
}

func TestParserFunctionEndStatement(t *testing.T) {
	lines, _ := ParseBasicCode("1000 END")
	if _, ok := lines[0].Statements[0].(*EndStatement); !ok {
		t.Error("Wront statement type returned")
	}
}

func TestParserOperationStatement(t *testing.T) {
	lines, _ := ParseBasicCode("1000 BARFOO 1")
	if _, ok := lines[0].Statements[0].(*OperationStatement); !ok {
		t.Error("Wront statement type returned")
	}
}

func TestParserRemarkStatement(t *testing.T) {
	lines, _ := ParseBasicCode("1000 REM \"tässä saa olla mitä vaan\"")
	if _, ok := lines[0].Statements[0].(*RemarkStatement); !ok {
		t.Error("Wront statement type returned")
	}
}

func TestParserMultipleStatementsOnSameLine(t *testing.T) {
	lines, _ := ParseBasicCode("1000 LET a = 1 : LET b = 2")
	if _, ok := lines[0].Statements[0].(*AssignmentStatement); !ok {
		t.Error("Wront statement type returned")
	}
	if _, ok := lines[0].Statements[1].(*AssignmentStatement); !ok {
		t.Error("Wront statement type returned")
	}
}

func TestParserMap(t *testing.T) {
	lines, _ := ParseBasicCode("1000 LET a = {}\n1010 LET b = {a: 123}\n1020 LET b = {a: 123, b: \"foo\"}")
	if _, ok := lines[0].Statements[0].(*AssignmentStatement); !ok {
		t.Error("Wront statement type returned")
	}
	if _, ok := lines[1].Statements[0].(*AssignmentStatement); !ok {
		t.Error("Wront statement type returned")
	}
	if _, ok := lines[2].Statements[0].(*AssignmentStatement); !ok {
		t.Error("Wront statement type returned")
	}
}

func TestParserMapAccess(t *testing.T) {
	lines, _ := ParseBasicCode("1010 LET b = a[\"a\"] : LET c = a.a : LET d = a[123] : LET a[123] = 123")
	if _, ok := lines[0].Statements[0].(*AssignmentStatement); !ok {
		t.Error("Wront statement type returned")
	}
	if _, ok := lines[0].Statements[1].(*AssignmentStatement); !ok {
		t.Error("Wront statement type returned")
	}
}

func TestParserArray(t *testing.T) {
	lines, _ := ParseBasicCode("1000 LET a = []\n1010 LET b = [123]\n1020 LET b = [123, \"foo\"]")
	if _, ok := lines[0].Statements[0].(*AssignmentStatement); !ok {
		t.Error("Wront statement type returned")
	}
	if _, ok := lines[1].Statements[0].(*AssignmentStatement); !ok {
		t.Error("Wront statement type returned")
	}
	if _, ok := lines[2].Statements[0].(*AssignmentStatement); !ok {
		t.Error("Wront statement type returned")
	}
}

func TestParserFunctionStuff(t *testing.T) {
	ParseBasicCode("2080 LET f(b(n+1)) = c(f(n) + f(n-1))")
}

func TestParserTestBoolean(t *testing.T) {
	lines, _ := ParseBasicCode("2080 LET booleanValue = true")
	if _, ok := lines[0].Statements[0].(*AssignmentStatement); !ok {
		t.Error("Wront statement type returned")
	}
}

func TestParserUnaryMinus(t *testing.T) {
	lines, _ := ParseBasicCode(`10 LET a = -1
20 LET b = -a`)
	if stm, ok := lines[0].Statements[0].(*AssignmentStatement); ok {
		if _, ok := stm.Expression.(*UnaryMinusExpression); !ok {
			t.Error("Wrong expression type returned")
		}
	} else {
		t.Error("Wront statement type returned")
	}
}

func TestParserErrorMessages(t *testing.T) {
	//ParserDebug = 2
	_, err := ParseBasicCode("foo")
	fmt.Println(err.Error())
}
