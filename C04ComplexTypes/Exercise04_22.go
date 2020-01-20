package main

import (
	"errors"
	"fmt"
)

func main() {
	res, _ := doubler04_22(-5)
	fmt.Println("-5  :", res)
	res, _ = doubler04_22(5)
	fmt.Println("5   :", res)
	res, _ = doubler04_22("yum")
	fmt.Println("yum :", res)
	res, _ = doubler04_22(true)
	fmt.Println("true:", res)
	res, _ = doubler04_22(float32(3.14))
	fmt.Println("3.14:", res)
}

func doubler04_22(v interface{}) (string, error) {

	switch t := v.(type) {
	case string:
		return t + t, nil
	case bool:
		if t {
			return "truetrue", nil
		}
		return "falsefalse", nil
	case float32, float64:
		if f, ok := t.(float64); ok {
			return fmt.Sprint(f * 2), nil
		}
		return fmt.Sprint(t.(float32) * 2), nil
	case int:
		return fmt.Sprint(t * 2), nil
	case int8:
		return fmt.Sprint(t * 2), nil
	case int16:
		return fmt.Sprint(t * 2), nil
	case int32:
		return fmt.Sprint(t * 2), nil
	case int64:
		return fmt.Sprint(t * 2), nil
	case uint:
		return fmt.Sprint(t * 2), nil
	case uint8:
		return fmt.Sprint(t * 2), nil
	case uint16:
		return fmt.Sprint(t * 2), nil
	case uint32:
		return fmt.Sprint(t * 2), nil
	case uint64:
		return fmt.Sprint(t * 2), nil
	default:
		return "", errors.New("unsupported type passed")
	}
}
