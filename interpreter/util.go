package interpreter

import (
	"errors"
	"fmt"
	"strconv"
)

func getStringValue(val interface{}) (str string, err error) {
	switch strval := val.(type) {
	case string:
		str = strval
	case float64:
		str = fmt.Sprintf("%f", strval)
	case int64:
		str = fmt.Sprintf("%d", strval)
	default:
		return "", errors.New("Type can't be converted to string")
	}

	return
}

func getFloatValue(val interface{}) (f float64, err error) {
	switch fval := val.(type) {
	case int64:
		f = float64(fval)
	case float64:
		f = fval
	case string:
		f, err = strconv.ParseFloat(fval, 64)
	default:
		return 0.0, errors.New("Type can't be converted to float")
	}

	return
}

func getIntValue(val interface{}) (f int64, err error) {
	switch fval := val.(type) {
	case int64:
		f = fval
	case float64:
		f = int64(fval)
	case string:
		f, err = strconv.ParseInt(fval, 10, 64)
	default:
		return 0.0, errors.New("Type can't be converted to float")
	}

	return
}

func getArrayPostion(dimensions []int64, args []interface{}) (int64, error) {
	pos := int64(0)
	argLength := len(args)
	for i := 0; i < argLength-1; i++ {
		val := args[i]
		intval, err := getIntValue(val)
		if err != nil {
			return 0, err
		}

		pos += intval * dimensions[i]
	}

	val := args[argLength-1]
	intval, err := getIntValue(val)
	if err != nil {
		return 0, err
	}
	pos += int64(intval)

	return pos, nil
}
