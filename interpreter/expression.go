package interpreter

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

type (
	AdditiveExpression struct {
		Minus bool
		Left  Expression
		Right Expression
	}
	MultiplicationExpression struct {
		Left  Expression
		Right Expression
	}
	DivisionExpression struct {
		Left  Expression
		Right Expression
	}
	PowerExpression struct {
		BaseExpression     Expression
		ExponentExpression Expression
	}
	RelationExpression struct {
		Comparator Comparator
		Left       Expression
		Right      Expression
	}
	Comparator interface {
		Compare(*Scope, Expression, Expression) bool
	}
	EqualsComparator struct {
		Negation bool
	}
	GreaterThanComparator struct {
		OrEqual bool
	}
	LessThanComparator struct {
		OrEqual bool
	}
	Expression interface {
		GetValue(*Scope) interface{}
	}
	MapPropertyExpression struct {
		Map      Expression
		Property Expression
	}
	FunctionCall struct {
		Expression Expression
		Args       Args
	}
	ArrayDefinition struct {
		Symbol     *IdentifierSymbol
		Dimensions Args
	}
	Array struct {
		Items Args
	}
	MapKind interface {
		GetProperty(key string) interface{}
		SetProperty(key string, val interface{})
	}
	ArrayKind interface {
		GetIndex(dims []int64) interface{}
		SetIndex(dims []int64, val interface{})
	}
	Map struct {
		Properties []*MapProperty
	}
	MapProperty struct {
		Name  string
		Value Expression
	}
	StringLiteral struct {
		Value string
	}
	Integer struct {
		Value int64
	}
	Float struct {
		Value float64
	}
	Boolean struct {
		Value bool
	}
	UnaryMinusExpression struct {
		Expression Expression
	}
	Null struct {
	}
	ArrayDef struct {
		data   []interface{}
		dims   []int64
		length int64 // Should be uint64
	}
)

func (unary *UnaryMinusExpression) GetValue(scope *Scope) interface{} {
	val := unary.Expression.GetValue(scope)

	switch val := val.(type) {
	case int64:
		return -val
	case float64:
		return -val
	case string:
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			panic(err)
		}
		return -f
	default:
		panic(fmt.Sprintf("-(Expression) cant be used with %+v", val))
	}
}

func (mapsym *MapPropertyExpression) GetValue(scope *Scope) interface{} {
	mapVal := mapsym.Map.GetValue(scope)
	propertyVal := mapsym.Property.GetValue(scope)
	strVal, err := getStringValue(propertyVal)
	if err != nil {
		panic(err)
	}
	switch mapInstance := mapVal.(type) {
	case map[string]interface{}:
		return mapInstance[strVal]
	case MapKind:
		return mapInstance.GetProperty(strVal)
	}
	// TODO proper error handling
	panic("Not a map")
}

func (mapsym *MapPropertyExpression) SetValue(scope *Scope, val interface{}) error {
	mapVal := mapsym.Map.GetValue(scope)
	propertyVal := mapsym.Property.GetValue(scope)
	strVal, err := getStringValue(propertyVal)
	if err != nil {
		return err
	}

	switch mapInstance := mapVal.(type) {
	case map[string]interface{}:
		mapInstance[strVal] = val
		return nil
	case MapKind:
		mapInstance.SetProperty(strVal, val)
		return nil
	}
	// TODO proper error handling
	panic("Not a map")
}

func (sym *FunctionCall) GetValue(scope *Scope) interface{} {
	// TODO figure out how to separe function calls from array operations

	args := make([]interface{}, len(sym.Args))
	for indx, exp := range sym.Args {
		args[indx] = exp.GetValue(scope)
	}

	val := sym.Expression.GetValue(scope)

	switch castVal := val.(type) {
	case *ArrayDef:
		pos, err := getArrayPostion(castVal.dims, args)
		if err != nil {
			panic(err)
		}
		// TODO check out-of-bounds error
		return castVal.data[pos]
	case func(args []interface{}) interface{}:
		return castVal(args)
	}

	panic("No such function or array")
}

func (sym *FunctionCall) SetValue(scope *Scope, val interface{}) error {
	// TODO figure out how to separe function calls from array operations

	args := make([]interface{}, len(sym.Args))
	for indx, exp := range sym.Args {
		args[indx] = exp.GetValue(scope)
	}

	if arr, ok := sym.Expression.GetValue(scope).(*ArrayDef); ok {
		pos, err := getArrayPostion(arr.dims, args)
		if err != nil {
			return err
		}
		arr.data[pos] = val
		return nil
	}

	panic("No such function or array")
}

func (stringLiteral *StringLiteral) GetValue(scope *Scope) interface{} {
	return stringLiteral.Value
}

func (integer *Integer) GetValue(scope *Scope) interface{} {
	return integer.Value
}

func (float *Float) GetValue(scope *Scope) interface{} {
	return float.Value
}

func (boolean *Boolean) GetValue(scope *Scope) interface{} {
	return boolean.Value
}

func (arr *Array) GetValue(scope *Scope) interface{} {
	length := int64(len(arr.Items))
	data := make([]interface{}, length)
	dims := []int64{length}

	for index, expr := range arr.Items {
		data[index] = expr.GetValue(scope)
	}

	return &ArrayDef{length: length, data: data, dims: dims}
}

func (m *Map) GetValue(scope *Scope) interface{} {
	minstance := map[string]interface{}{}

	for _, property := range m.Properties {
		minstance[property.Name] = property.Value.GetValue(scope)
	}

	return minstance
}

func (relation *RelationExpression) GetValue(scope *Scope) bool {
	return relation.Comparator.Compare(scope, relation.Left, relation.Right)
}

func (comparator *EqualsComparator) Compare(scope *Scope, left Expression, right Expression) bool {
	return comparator.Negation != (left.GetValue(scope) == right.GetValue(scope))
}
func (comparator *GreaterThanComparator) Compare(scope *Scope, left Expression, right Expression) bool {
	leftValue := left.GetValue(scope)
	rightValue := right.GetValue(scope)

	switch leftValue := leftValue.(type) {
	case int64:
		switch rightValue := rightValue.(type) {
		case int64:
			if comparator.OrEqual {
				return leftValue >= rightValue
			}
			return leftValue > rightValue
		case float64:
			if comparator.OrEqual {
				return float64(leftValue) >= rightValue
			}
			return float64(leftValue) > rightValue
		}
	case float64:
		switch rightValue := rightValue.(type) {
		case int64:
			if comparator.OrEqual {
				return leftValue >= float64(rightValue)
			}
			return leftValue > float64(rightValue)
		case float64:
			if comparator.OrEqual {
				return leftValue >= rightValue
			}
			return leftValue > rightValue
		}
	}
	return false
}
func (comparator *LessThanComparator) Compare(scope *Scope, left Expression, right Expression) bool {
	leftValue := left.GetValue(scope)
	rightValue := right.GetValue(scope)

	switch leftValue := leftValue.(type) {
	case int64:
		switch rightValue := rightValue.(type) {
		case int64:
			if comparator.OrEqual {
				return leftValue <= rightValue
			}
			return leftValue < rightValue
		case float64:
			if comparator.OrEqual {
				return float64(leftValue) <= rightValue
			}
			return float64(leftValue) < rightValue
		}
	case float64:
		switch rightValue := rightValue.(type) {
		case int64:
			if comparator.OrEqual {
				return leftValue <= float64(rightValue)
			}
			return leftValue < float64(rightValue)
		case float64:
			if comparator.OrEqual {
				return leftValue <= rightValue
			}
			return leftValue < rightValue
		}
	}

	return false
}

func (sym *MultiplicationExpression) GetValue(scope *Scope) interface{} {
	left := sym.Left
	right := sym.Right

	leftValue := left.GetValue(scope)
	rightValue := right.GetValue(scope)

	switch leftValue := leftValue.(type) {
	case int64:
		switch rightValue := rightValue.(type) {
		case int64:
			return leftValue * rightValue
		case float64:
			return float64(leftValue) * rightValue
		case string:
			floatValue, err := strconv.ParseFloat(rightValue, 64)
			if err != nil {
				panic(err)
			}

			return float64(leftValue) * floatValue
		}
	case float64:
		switch rightValue := rightValue.(type) {
		case int64:
			return leftValue * float64(rightValue)
		case float64:
			return leftValue * rightValue
		case string:
			floatValue, err := strconv.ParseFloat(rightValue, 64)
			if err != nil {
				panic(err)
			}
			return leftValue * floatValue
		}
	case string:
		switch rightValue := rightValue.(type) {
		case int64:
		case float64:
			floatValue, err := strconv.ParseFloat(leftValue, 64)
			if err != nil {
				panic(err)
			}

			return floatValue * float64(rightValue)
		case string:
			rightFloatValue, err := strconv.ParseFloat(rightValue, 64)
			if err != nil {
				panic(err)
			}
			leftFloatValue, err := strconv.ParseFloat(leftValue, 64)
			if err != nil {
				panic(err)
			}
			return leftFloatValue * rightFloatValue
		}
	}

	panic("Invalid multiplication expression")
}

func (sym *DivisionExpression) GetValue(scope *Scope) interface{} {
	left := sym.Left
	right := sym.Right

	leftF, err := getFloatValue(left.GetValue(scope))
	if err != nil {
		panic(err)
	}
	rightF, err := getFloatValue(right.GetValue(scope))
	if err != nil {
		panic(err)
	}

	return leftF / rightF
}

func (sym *PowerExpression) GetValue(scope *Scope) interface{} {
	base := sym.BaseExpression
	exponent := sym.ExponentExpression

	baseF, err := getFloatValue(base.GetValue(scope))
	if err != nil {
		panic(err)
	}
	exponentF, err := getFloatValue(exponent.GetValue(scope))
	if err != nil {
		panic(err)
	}

	return math.Pow(baseF, exponentF)
}

func (sym *AdditiveExpression) GetValue(scope *Scope) interface{} {
	left := sym.Left
	right := sym.Right

	leftValue := left.GetValue(scope)
	rightValue := right.GetValue(scope)
	minus := sym.Minus

	switch leftValue := leftValue.(type) {
	case int64:
		switch rightValue := rightValue.(type) {
		case int64:
			if minus {
				return leftValue - rightValue
			}
			return leftValue + rightValue
		case float64:
			if minus {
				return float64(leftValue) - rightValue
			}
			return float64(leftValue) + rightValue
		case string:
			floatValue, err := strconv.ParseFloat(rightValue, 64)
			if err != nil {
				panic(err)
			}
			if minus {
				return float64(leftValue) - floatValue
			}

			return fmt.Sprintf("%d%s", leftValue, rightValue)
		}
	case float64:
		switch rightValue := rightValue.(type) {
		case int64:
			if minus {
				return leftValue - float64(rightValue)
			}
			return leftValue + float64(rightValue)
		case float64:
			if minus {
				return leftValue - rightValue
			}
			return leftValue + rightValue
		case string:
			floatValue, err := strconv.ParseFloat(rightValue, 64)
			if err != nil {
				panic(err)
			}
			if minus {
				return leftValue - floatValue
			}
			return fmt.Sprintf("%f%s", leftValue, rightValue)
		}
	case string:
		switch rightValue := rightValue.(type) {
		case int64:
			floatValue, err := strconv.ParseFloat(leftValue, 64)
			if err != nil {
				panic(err)
			}
			if minus {
				return floatValue - float64(rightValue)
			}

			return fmt.Sprintf("%s%d", leftValue, rightValue)
		case float64:
			floatValue, err := strconv.ParseFloat(leftValue, 64)
			if err != nil {
				panic(err)
			}
			if minus {
				return floatValue - rightValue
			}
			return fmt.Sprintf("%s%f", leftValue, rightValue)
		case string:
			rightFloatValue, err := strconv.ParseFloat(rightValue, 64)
			if err != nil {
				panic(err)
			}
			leftFloatValue, err := strconv.ParseFloat(leftValue, 64)
			if err != nil {
				panic(err)
			}
			if minus {
				return leftFloatValue - rightFloatValue
			}
			return fmt.Sprintf("%s%s", leftValue, rightValue)
		}
	}

	fmt.Printf("%s %s\n", reflect.TypeOf(leftValue), reflect.TypeOf(rightValue))

	panic("Invalid additive expression")
}
