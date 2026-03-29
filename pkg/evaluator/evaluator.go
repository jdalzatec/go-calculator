package evaluator

import "errors"

func Evaluate(lhs int64, rhs int64, operator string) (*int64, error) {
	switch operator {
	case "+":
		return new(lhs + rhs), nil
	case "-":
		return new(lhs - rhs), nil
	case "*":
		return new(lhs * rhs), nil
	case "/":
		if rhs == 0 {
			return nil, errors.New("rhs must be different than 0")
		}
		return new(lhs / rhs), nil
	default:
		return nil, errors.New("unknown operator")
	}
}
