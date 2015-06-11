package runtime

import (
	"errors"

	"github.com/orktes/nodebas/interpreter"
)

type httpResponse struct {
}

func (response *httpResponse) GetProperty(key string) interface{} {
	return nil
}
func (response *httpResponse) SetProperty(key string, val interface{}) {
	return
}

type httpRequest struct {
}

func (request *httpRequest) GetProperty(key string) interface{} {
	return nil
}
func (request *httpRequest) SetProperty(key string, val interface{}) {
	return
}

type httpInfo struct {
	response *httpResponse
	request  *httpRequest
}

func init() {
	scope := interpreter.DefaultScope
	scope.SetOperator("HTTPR", httpRead)
	scope.SetOperator("HTTPW", httpWrite)
	scope.SetOperator("HTTPWH", httpWriteHeader)
}

func readFromHTTP(port string) *httpInfo {
	return nil
}

func httpRead(intpr *interpreter.Interpreter, args interpreter.Args) (int, error) {
	if len(args) < 3 {
		return 0, errors.New("HTTPR: Invalid number of arguments")
	}

	port, err := intpr.GetStringValue(args[0])
	if err != nil {
		return 0, err
	}

	httpInf := readFromHTTP(port)

	if sym, ok := args[1].(interpreter.Symbol); ok {
		sym.SetValue(intpr.Scope, httpInf.request)
	} else {
		return 0, errors.New("HTTPR: Second argument should be a symbol")
	}

	if sym, ok := args[2].(interpreter.Symbol); ok {
		sym.SetValue(intpr.Scope, httpInf.response)
	} else {
		return 0, errors.New("HTTPR: Second argument should be a symbol")
	}

	return 0, nil
}

func httpWrite(interpreter *interpreter.Interpreter, args interpreter.Args) (int, error) {
	return 0, nil
}

func httpWriteHeader(interpreter *interpreter.Interpreter, args interpreter.Args) (int, error) {
	return 0, nil
}
