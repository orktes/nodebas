package interpreter

import "fmt"

var DefaultScope *Scope = NewScope()

type ForLoop struct {
	BodyLineNumber  LineNumber
	ControlVariable *IdentifierSymbol
	Increment       float64
	Limit           float64
	Reverse         bool
}

type Interpreter struct {
	Scope       *Scope
	Lines       Lines
	lineMap     map[int64]int
	forLoops    map[string]*ForLoop
	returnLines []int
	gotoLine    int
	lineCount   int
	stopped     bool
}

// Util methods
func (interpreter *Interpreter) GetIntValue(expression Expression) (int64, error) {
	return getIntValue(expression.GetValue(interpreter.Scope))
}

func (interpreter *Interpreter) GetFloatValue(expression Expression) (float64, error) {
	return getFloatValue(expression.GetValue(interpreter.Scope))
}

func (interpreter *Interpreter) GetStringValue(expression Expression) (string, error) {
	return getStringValue(expression.GetValue(interpreter.Scope))
}

func (interpreter *Interpreter) Step() error {
	gotoLine := interpreter.gotoLine

	if gotoLine < 0 || gotoLine > interpreter.lineCount-1 {
		interpreter.stopped = true
		return nil
	}

	line := interpreter.Lines[gotoLine]

	interpreter.gotoLine++

	for _, statement := range line.Statements {
		_, err := statement.Eval(interpreter)
		if err != nil {
			return err
		}
	}

	return nil
}

func (interpreter *Interpreter) Goto(lineNumber int64) (err error) {
	if physicalLine, ok := interpreter.lineMap[lineNumber]; ok {
		interpreter.gotoLine = physicalLine
	} else {
		// TODO proper error
		err = fmt.Errorf("No such line %d", lineNumber)
	}
	return
}

func (interpreter *Interpreter) Run() (err error) {
	for !interpreter.stopped {
		err = interpreter.Step()
		if err != nil {
			return
		}
	}

	return
}

func NewIntepreterWithScope(lines Lines, scope *Scope) *Interpreter {
	lineMap := map[int64]int{}
	forLoops := map[string]*ForLoop{}

	for indx, line := range lines {
		lineMap[int64(line.Number)] = indx
	}

	return &Interpreter{
		Lines:     lines,
		Scope:     scope,
		gotoLine:  0,
		lineCount: len(lines),
		lineMap:   lineMap,
		forLoops:  forLoops,
	}
}

func NewIntepreter(lines Lines) *Interpreter {
	return NewIntepreterWithScope(lines, DefaultScope)
}
