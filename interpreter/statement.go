package interpreter

import (
	"errors"
	"fmt"
)

type (
	Statements []Statement
	Statement  interface {
		Eval(interpreter *Interpreter) (int, error)
	}
	EndStatement struct {
	}
	RemarkStatement struct {
	}
	GotoStatement struct {
		LineNumber Expression
	}
	GoSubStatement struct {
		LineNumber Expression
	}
	ReturnStatement struct {
	}
	DefStatement struct {
		Identifier *IdentifierSymbol
		Parameter  string
		Expression Expression
	}
	ForStatement struct {
		ControlVariable *IdentifierSymbol
		InitialValue    Expression
		Limit           Expression
		Increment       Expression
	}
	NextStatement struct {
		ControlVariable *IdentifierSymbol
	}
	AssignmentStatement struct {
		Symbol     Symbol
		Expression Expression
	}
	DimStatement struct {
		ArrayDefinitions []*ArrayDefinition
	}
	IfStatement struct {
		RelationExpression *RelationExpression
		LineNumber         Expression
	}
	Args               []Expression
	OperationStatement struct {
		Operation string
		Args      Args
	}
	PrintStatement struct {
		Args Args
	}
)

type EmptyStatement struct {
}

func (_ *EmptyStatement) Eval(_ *Interpreter) (int, error) {
	return 0, nil
}

func (_ *RemarkStatement) Eval(_ *Interpreter) (int, error) {
	return 0, nil
}

func (stm *OperationStatement) Eval(interpreter *Interpreter) (int, error) {
	if operator, ok := interpreter.Scope.operators[stm.Operation]; ok {
		return operator(interpreter, stm.Args)
	}
	return 0, fmt.Errorf("No such operator %s", stm.Operation)
}

func (_ *EndStatement) Eval(interpreter *Interpreter) (int, error) {
	interpreter.stopped = true
	return 0, nil
}

func (stm *DimStatement) Eval(interpreter *Interpreter) (int, error) {
	scope := interpreter.Scope

	for _, arrs := range stm.ArrayDefinitions {
		dims := make([]int64, len(arrs.Dimensions))
		length := int64(0)

		ieLength := len(arrs.Dimensions)
		iexp := arrs.Dimensions[0]
		// TODO support float index
		if val, ok := iexp.GetValue(scope).(int64); ok {
			length = val
			dims[0] = val
		}

		for i := 1; i < ieLength; i++ {
			iexp := arrs.Dimensions[i]
			// TODO support float index
			if val, ok := iexp.GetValue(scope).(int64); ok {
				length += length * val
				dims[i] = val
			} else {
				return 0, errors.New("Index not a number")
			}
		}

		data := make([]interface{}, length)

		arrs.Symbol.SetValue(scope, &ArrayDef{dims: dims, data: data, length: length})
	}

	return 0, nil
}

func (stm *GotoStatement) Eval(interpreter *Interpreter) (int, error) {
	// TODO check that line is inside correct range

	val := stm.LineNumber.GetValue(interpreter.Scope)

	lineNumberInt, err := getIntValue(val)
	if err != nil {
		return 0, err
	}

	interpreter.Goto(lineNumberInt)
	return 0, nil
}

func (stm *PrintStatement) Eval(interpreter *Interpreter) (int, error) {
	for _, arg := range stm.Args {
		fmt.Print(arg.GetValue(interpreter.Scope))
	}
	fmt.Print("\n")
	return 0, nil
}

func (stm *ForStatement) Eval(interpreter *Interpreter) (result int, err error) {
	var limitFloat float64
	var initialFloat float64

	scope := interpreter.Scope
	increment := float64(1)

	if stm.Increment != nil {
		incrementValue := stm.Increment.GetValue(scope)
		increment, err = getFloatValue(incrementValue)
		if err != nil {
			// TODO proper error message
			return
		}
	}

	limitValue := stm.Limit.GetValue(scope)
	limitFloat, err = getFloatValue(limitValue)
	if err != nil {
		return
	}

	initialValue := stm.InitialValue.GetValue(scope)
	initialFloat, err = getFloatValue(initialValue)
	if err != nil {
		return
	}

	control := stm.ControlVariable

	fl := &ForLoop{
		BodyLineNumber:  LineNumber(interpreter.gotoLine),
		Increment:       increment,
		Limit:           limitFloat,
		ControlVariable: control,
		Reverse:         (initialFloat > limitFloat),
	}

	control.SetValue(scope, initialValue)
	interpreter.forLoops[control.Name] = fl

	return
}

func (stm *NextStatement) Eval(interpreter *Interpreter) (int, error) {
	scope := interpreter.Scope
	control := stm.ControlVariable
	fl := interpreter.forLoops[control.Name]

	if fl == nil {
		// TODO proper errors message
		return 0, errors.New("NEXT " + control.Name + " but no FOR command")
	}

	controlValue := control.GetValue(scope)
	controlValueFloat, err := getFloatValue(controlValue)
	if err != nil {
		return 0, err
	}

	controlValueFloat += fl.Increment

	if (!fl.Reverse && controlValueFloat <= fl.Limit) ||
		(fl.Reverse && controlValueFloat >= fl.Limit) {
		control.SetValue(scope, controlValueFloat)
		interpreter.gotoLine = int(fl.BodyLineNumber)
	}

	return 0, nil
}

func (stm *GoSubStatement) Eval(interpreter *Interpreter) (int, error) {
	val := stm.LineNumber.GetValue(interpreter.Scope)

	lineNumberInt, err := getIntValue(val)
	if err != nil {
		return 0, err
	}

	interpreter.returnLines = append(interpreter.returnLines, interpreter.gotoLine)
	interpreter.Goto(lineNumberInt)

	return 0, nil
}

func (stm *ReturnStatement) Eval(interpreter *Interpreter) (int, error) {
	returnLength := len(interpreter.returnLines)
	if returnLength == 0 {
		return 0, errors.New("RETURN but no GO SUB")
	}

	pos := returnLength - 1
	returnLine := interpreter.returnLines[pos]
	interpreter.returnLines = interpreter.returnLines[:pos]

	interpreter.gotoLine = returnLine

	return 0, nil
}

func (stm *DefStatement) Eval(interpreter *Interpreter) (int, error) {
	scope := interpreter.Scope
	stm.Identifier.SetValue(scope, func(args []interface{}) interface{} {

		if len(stm.Parameter) != 0 {
			scope = scope.SubScope()
			if len(args) == 0 {
				// TODO proper error message
				panic("Function requires an attribute " + stm.Parameter)
			}
			scope.Set(stm.Parameter, args[0])
		}

		return stm.Expression.GetValue(scope)
	})
	return 0, nil
}

func (stm *AssignmentStatement) Eval(interpreter *Interpreter) (int, error) {
	scope := interpreter.Scope
	stm.Symbol.SetValue(scope, stm.Expression.GetValue(scope))
	return 0, nil
}

func (stm *IfStatement) Eval(interpreter *Interpreter) (int, error) {
	if stm.RelationExpression.GetValue(interpreter.Scope) {
		val := stm.LineNumber.GetValue(interpreter.Scope)

		lineNumberInt, err := getIntValue(val)
		if err != nil {
			return 0, err
		}

		interpreter.Goto(lineNumberInt)
	}
	return 0, nil
}
