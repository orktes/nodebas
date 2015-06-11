package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/orktes/nodebas/interpreter"
)

func main() {
	js.Module.Get("exports").Set("Eval", Eval)
}

func Eval(code string) *js.Object {
	interpreter.EvalBasicCode(code)
	return nil
}
