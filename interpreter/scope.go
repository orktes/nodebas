package interpreter

type Scope struct {
	variables   map[string]interface{}
	operators   map[string]func(interpreter *Interpreter, args Args) (int, error)
	parentScope *Scope
}

func (scope *Scope) Get(key string) interface{} {
	if val, ok := scope.variables[key]; ok {
		return val
	}

	if scope.parentScope != nil {
		return scope.parentScope.Get(key)
	}

	return 0
}

func (scope *Scope) Set(key string, val interface{}) {
	scope.variables[key] = val
}

func (scope *Scope) SetString(key string, val string) {
	scope.Set(key, val)
}

func (scope *Scope) SetInteger(key string, val int64) {
	scope.Set(key, val)
}

func (scope *Scope) SetFloat(key string, val float64) {
	scope.Set(key, val)
}

func (scope *Scope) SetFunc(key string, val func(args []interface{}) interface{}) {
	scope.Set(key, val)
}

func (scope *Scope) SetArray(key string, val *ArrayDef) {
	scope.Set(key, val)
}

func (scope *Scope) SetOperator(key string, operator func(interpreter *Interpreter, args Args) (int, error)) {
	scope.operators[key] = operator
}

func (scope *Scope) SubScope() (subScope *Scope) {
	subScope = NewScope()
	subScope.parentScope = scope
	return
}

func (scope *Scope) Copy() (copiedScope *Scope) {
	copiedScope = &Scope{}
	variables := map[string]interface{}{}
	for k, v := range copiedScope.variables {
		variables[k] = v
	}
	copiedScope.variables = variables

	operators := map[string]func(interpreter *Interpreter, args Args) (int, error){}
	for k, v := range copiedScope.operators {
		operators[k] = v
	}
	copiedScope.operators = operators
	return
}

func NewScope() (scope *Scope) {
	scope = &Scope{
		variables: map[string]interface{}{},
		operators: map[string]func(interpreter *Interpreter, args Args) (int, error){},
	}
	return
}
