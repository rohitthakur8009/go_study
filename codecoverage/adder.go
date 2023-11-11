package codecoverage

import (
	"errors"
	"fmt"
)

func add(val1, val2 interface{}) (interface{}, error){
	switch val1.(type) {
	case int:
		return val1.(int) + val2.(int), nil
	case string:
		return fmt.Sprintf("%s %s", val1, val2), nil
	default:
		return nil, errors.New("Unsupported Type")
	}
}


