package interpreter

import (
	"fmt"
	"testing"

	"github.com/Shopify/go-lua"
	"github.com/robertkrimen/otto"
)

func TestInterpreter(t *testing.T) {
	lines, _ := ParseBasicCode("1001 PRINT \"kissa\",\"kissa\"\n1002 GOTO 1001")
	interpreter := NewIntepreter(lines)

	interpreter.Step()
	interpreter.Step()
	interpreter.Step()
}

func TestInterpreterAssignment(t *testing.T) {
	lines, _ := ParseBasicCode("1001 LET a = \"foobar\"\n1002 LET b = a\n1003 LET b = 123\n1004 PRINT \"a = \", a, \", b = \", b")
	interpreter := NewIntepreter(lines)

	interpreter.Step()
	if interpreter.Scope.Get("a") != "foobar" {
		t.Error("Wrong value assigned", interpreter.Scope.Get("a"))
	}

	interpreter.Step()
	if interpreter.Scope.Get("b") != "foobar" {
		t.Error("Wrong value assigned", interpreter.Scope.Get("b"))
	}

	interpreter.Step()
	if interpreter.Scope.Get("b") != int64(123) {
		t.Error("Wrong value assigned", interpreter.Scope.Get("b"))
	}

	interpreter.Step()
}

func TestInterpreterIF(t *testing.T) {
	lines, _ := ParseBasicCode("10 LET a = 123\n20 IF a > 120 THEN 40\n30 LET a = 199\n40 PRINT a")
	interpreter := NewIntepreter(lines)
	interpreter.Run()
	if interpreter.Scope.Get("a") != int64(123) {
		t.Error("If did not work")
	}

	lines, _ = ParseBasicCode("10 LET a = 119\n20 IF a > 120 THEN 40\n30 LET a = 123\n40 PRINT a")
	interpreter = NewIntepreter(lines)
	interpreter.Run()
	if interpreter.Scope.Get("a") != int64(123) {
		t.Error("If did not work")
	}
}

func TestInterpreterDIM(t *testing.T) {
	lines, _ := ParseBasicCode("10 DIM foo(1)\n20 LET foo(0) = 1")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if val, ok := interpreter.Scope.Get("foo").(*ArrayDef); ok {
		if val.data[0] != int64(1) {
			t.Error("Wrong value in array")
		}
	} else {
		t.Error("Wrong type for array")
	}

	lines, _ = ParseBasicCode("10 DIM foo(2,2)\n20 LET foo(0,0) = 1\n30 LET foo(0,1) = 2\n40 LET foo(1,0) = 3\n50 LET foo(1,1) = 4")
	interpreter = NewIntepreter(lines)
	interpreter.Run()

	if val, ok := interpreter.Scope.Get("foo").(*ArrayDef); ok {
		if val.data[0] != int64(1) {
			t.Error("Wrong value in array")
		}
		if val.data[1] != int64(2) {
			t.Error("Wrong value in array")
		}
		if val.data[2] != int64(3) {
			t.Error("Wrong value in array")
		}
		if val.data[3] != int64(4) {
			t.Error("Wrong value in array")
		}
	} else {
		t.Error("Wrong type for array")
	}
}

func TestArrayAccessWithExpression(t *testing.T) {
	lines, _ := ParseBasicCode("10 LET n = 1\n15 DIM f(1)\n20 LET f(n - 1) = 1")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if val, ok := interpreter.Scope.Get("f").(*ArrayDef); ok {
		if val.data[0] != int64(1) {
			t.Error("Wrong value in array")
		}
	} else {
		t.Error("Wrong type for array")
	}
}

func TestInterpreterAdditiveExpression(t *testing.T) {
	lines, _ := ParseBasicCode("10 LET a = 1 + 2\n20 LET b = \"1\" + 1\n30 LET c = \"2\" - 1\n40 LET d =  1 - \"1.5\"")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	fmt.Printf("%+v\n", interpreter.Scope.variables)

	if interpreter.Scope.Get("a") != int64(3) {
		t.Error("Wrong addition result")
	}

	if interpreter.Scope.Get("b") != "11" {
		t.Error("Wrong addition result")
	}

	if interpreter.Scope.Get("c") != float64(1) {
		t.Error("Wrong addition result")
	}

	if interpreter.Scope.Get("d") != float64(-0.5) {
		t.Error("Wrong addition result")
	}
}

func TestInterpreterFunctionCalls(t *testing.T) {
	lines, _ := ParseBasicCode("10 DEF FNF = 1 + 1\n20 LET a = FNF()")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if interpreter.Scope.Get("a") != int64(2) {
		t.Error("Wrong result for function")
	}

	lines, _ = ParseBasicCode("10 DEF FNF(x) = x + 1\n20 LET a = FNF(2)")
	interpreter = NewIntepreter(lines)
	interpreter.Run()

	if interpreter.Scope.Get("a") != int64(3) {
		t.Error("Wrong result for function")
	}
}

func TestInterpreterDivisionExpression(t *testing.T) {
	lines, _ := ParseBasicCode("1001 LET a = 3/1.5/4")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if interpreter.Scope.Get("a") != float64(0.5) {
		t.Error("Wrong result for division")
	}
}

func TestInterpreterMultiplicationExpression(t *testing.T) {
	lines, _ := ParseBasicCode("1001 LET a = 1*2*3*4")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if interpreter.Scope.Get("a") != int64(1*2*3*4) {
		t.Error("Wrong result for division")
	}
}

func TestInterpreterPowerExpression(t *testing.T) {
	lines, _ := ParseBasicCode("1001 LET a = 2^2")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if interpreter.Scope.Get("a") != float64(4) {
		t.Error("Wrong result for exponent")
	}
}

func TestInterpreterAlgebraPrecedence(t *testing.T) {
	lines, _ := ParseBasicCode("1001 LET a = 2+1*2\n1002 LET b = (2+1)*2")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if interpreter.Scope.Get("a") != int64(4) {
		t.Error("Wrong result for exponent")
	}

	if interpreter.Scope.Get("b") != int64(6) {
		t.Error("Wrong result for exponent")
	}

}

func TestInterpreterForLoop(t *testing.T) {
	lines, _ := ParseBasicCode("10 LET b = 0\n20 FOR a = 1 TO 30\n40 LET b = b + a\n60 NEXT a")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if interpreter.Scope.Get("b") != float64(465) {
		t.Error("For loop iteration failed")
	}

}

func TestInterpreterSubRoutine(t *testing.T) {
	lines, _ := ParseBasicCode("10 GO SUB 40\n20 LET b = a\n30 GOTO 80\n40 LET a = 100\n50 RETURN\n80 PRINT b")
	interpreter := NewIntepreter(lines)
	interpreter.Run()
}

func TestInterpreterEndStatement(t *testing.T) {
	lines, _ := ParseBasicCode("10 END\n20 LET a = 1")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if interpreter.Scope.Get("a") == int64(1) {
		t.Error("No exit")
	}
}

func TestInterpreterDefStatement(t *testing.T) {
	lines, _ := ParseBasicCode("10 DEF FNF = 1 + 1\n20 LET y = 100\n30 LET x = 100\n40 DEF FNB(x) = x + y")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if fooFN, ok := interpreter.Scope.Get("FNF").(func(args []interface{}) interface{}); ok {
		if fooFN([]interface{}{}) != int64(2) {
			t.Error("Wrong result for function call")
		}
	} else {
		t.Error("foo has wrong type")
	}

	if barFN, ok := interpreter.Scope.Get("FNB").(func(args []interface{}) interface{}); ok {
		if barFN([]interface{}{int64(1)}) != int64(101) {
			t.Error("Wrong result for function call")
		}
	} else {
		t.Error("bar has wrong type")
	}
}

func TestInterpreterMultipleStatementsOneLine(t *testing.T) {
	lines, _ := ParseBasicCode("1000 LET a = 1 : LET b = 2")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if interpreter.Scope.Get("a") != int64(1) {
		t.Error("Wrong value for a")
	}
	if interpreter.Scope.Get("b") != int64(2) {
		t.Error("Wrong value for b")
	}
}

func TestInterpreterInlineMap(t *testing.T) {
	lines, _ := ParseBasicCode("1000 LET a = {foo: \"bar\"}")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if (interpreter.Scope.Get("a").(map[string]interface{}))["foo"] != "bar" {
		t.Error("Invalid map property")
	}
}

func TestInterpreterMapAccess(t *testing.T) {
	lines, _ := ParseBasicCode("1000 LET a = {foo: \"bar\"} : LET b = a.foo")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if interpreter.Scope.Get("b") != "bar" {
		t.Error("Invalid map property")
	}
}

func TestInterpreterMapAccessInline(t *testing.T) {
	lines, _ := ParseBasicCode("1000 LET b = ({foo: \"bar\"}).foo : LET c = ({foo:{bar:\"baz\"}}).foo.bar")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if interpreter.Scope.Get("b") != "bar" {
		t.Error("Invalid map property")
	}
	if interpreter.Scope.Get("c") != "baz" {
		t.Error("Invalid map property")
	}
}

func TestInterpreterInlineArray(t *testing.T) {
	lines, _ := ParseBasicCode("1000 LET a = [\"bar\"]")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if (interpreter.Scope.Get("a").(*ArrayDef)).data[0] != "bar" {
		t.Error("Invalid array property")
	}
}

func TestInterpreterInlineArrayAccess(t *testing.T) {
	lines, _ := ParseBasicCode("1000 LET a = [\"bar\"](0)")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if interpreter.Scope.Get("a") != "bar" {
		t.Error("Invalid array property")
	}
}

func TestInterpreterCustomOperator(t *testing.T) {
	lines, _ := ParseBasicCode("1000 DIM a(1)\n1100 BARFOO a(0), b, 3")
	interpreter := NewIntepreter(lines)

	scope := interpreter.Scope

	scope.SetOperator("BARFOO", func(interpreter *Interpreter, args Args) (int, error) {
		a := args[0]
		b := args[1]
		int1 := args[2]

		(a.(Symbol)).SetValue(scope, int64(1*int1.GetValue(scope).(int64)))
		(b.(Symbol)).SetValue(scope, int64(2*int1.GetValue(scope).(int64)))

		return 0, nil
	})
	interpreter.Run()

	if interpreter.Scope.Get("a").(*ArrayDef).data[0] != int64(3) {
		t.Error("Something went wrong")
	}
	if interpreter.Scope.Get("b") != int64(6) {
		t.Error("Something went wrong")
	}
}

func TestInterpreterUnaryMinux(t *testing.T) {
	lines, _ := ParseBasicCode("2030 LET a = -1 : LET b = -a : LET c = -(a + b + 1) : LET d = -({foo: \"123\"}).foo")
	interpreter := NewIntepreter(lines)
	interpreter.Run()

	if interpreter.Scope.Get("a") != int64(-1) {
		t.Error("UnaryMinus problem!!!")
	}
	if interpreter.Scope.Get("b") != int64(1) {
		t.Error("UnaryMinus problem!!!")
	}
	if interpreter.Scope.Get("c") != int64(-1) {
		t.Error("UnaryMinus problem!!!")
	}
	if interpreter.Scope.Get("d") != float64(-123) {
		t.Errorf("UnaryMinus problem!!! %+v", interpreter.Scope.Get("d"))
	}
}

func TestInterpreterFibonassi(t *testing.T) {
	lines, _ := ParseBasicCode("2030 DIM f(51)\n2040 LET f(0) = 0\n2050 LET f(1) = 1\n2060 LET n = 1\n2080 LET f(n+1) = f(n) + f(n-1)\n2090 LET n = n + 1\n2100 PRINT f(n)\n2120 IF n < 45 THEN 2080")
	interpreter := NewIntepreter(lines)
	interpreter.Run()
}

func BenchmarkInterpreterGOTOFibonassi(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		lines, _ := ParseBasicCode("2030 DIM f(51)\n2040 LET f(0) = 0\n2050 LET f(1) = 1\n2060 LET n = 1\n2080 LET f(n+1) = f(n) + f(n-1)\n2090 LET n = n + 1\n2120 IF n < 45 THEN 2080")
		interpreter := NewIntepreter(lines)
		interpreter.Run()
	}
}

func BenchmarkInterpreterFORFibonassi(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		lines, _ := ParseBasicCode("2030 DIM f(51)\n2040 LET f(0) = 0\n2050 LET f(1) = 1\n2060 FOR n = 2 TO 45\n2080 LET f(n) = f(n-1) + f(n-2)\n2090 NEXT n")
		interpreter := NewIntepreter(lines)
		interpreter.Run()
	}

}

func BenchmarkOttoJSFibonassi(b *testing.B) {
	// run the Fib function b.N times
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm := otto.New()
		vm.Run(`
			var i;
var fib = []; //Initialize array!

fib[0] = 0;
fib[1] = 1;
for(i=2; i<=45; i++)
{
    // Next fibonacci number = previous + one before previous
    // Translated to JavaScript:
    fib[i] = fib[i-2] + fib[i-1];
}
		`)
	}

}

func BenchmarkLuaFibonassi(b *testing.B) {
	// run the Fib function b.N times
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l := lua.NewState()
		lua.DoString(l, `
function fastfib(n)
	fibs={1,1}

	for i=3,n do
		fibs[i]=fibs[i-1]+fibs[i-2]
	end

	return fibs[n]
end

fastfib(45)
			`)
	}

}
