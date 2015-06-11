package interpreter

type (
	Symbol interface {
		Expression
		SetValue(*Scope, interface{}) error
	}
	IdentifierSymbol struct {
		Name string
	}
)

func (sym *IdentifierSymbol) GetValue(scope *Scope) interface{} {
	return scope.Get(sym.Name)
}

func (sym *IdentifierSymbol) SetValue(scope *Scope, val interface{}) error {
	scope.Set(sym.Name, val)
	return nil
}
