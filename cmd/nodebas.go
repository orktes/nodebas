package main

import (
	"io/ioutil"
	"os"

	"github.com/orktes/nodebas/interpreter"
	_ "github.com/orktes/nodebas/runtime" // Automaticly adds default runtime func and operands to scope
)

func main() {
	argsWithoutProg := os.Args[1:]
	data, err := ioutil.ReadFile(argsWithoutProg[0])
	if err != nil {
		panic(err)
	}

	interpreter.EvalBasicCode(string(data))
}
