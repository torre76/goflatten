// Package flatten constains ,ethods that are used to flatten nested slices
package flatten

import "errors"

// Int flattens a slice that contains integers
func Int(target interface{}) ([]int, error) {
	return recursiveFlattenInt([]int{}, target)
}

// recursiveFlattenInt is a recursive method to flatten a slice of integers
func recursiveFlattenInt(buffer []int, target interface{}) ([]int, error) {
	var err error

	switch value := target.(type) {
	case int:
		// if type is an integer this is the recursion leaf, add to buffer
		buffer = append(buffer, value)
	case []int:
		// if type is a slice of integers, it is also a leaf, we could copy the values into buffer
		buffer = append(buffer, value...)
	case []interface{}:
		// This is not a leaf, apply recursion
		for i := range value {
			buffer, err = recursiveFlattenInt(buffer, value[i])
			if err != nil {
				return nil, errors.New("target type mismatch - int or []int allowed")
			}
		}
	default:
		return nil, errors.New("target type mismatch - int or []int allowed")
	}

	return buffer, nil
}
