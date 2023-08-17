package db

import (
	"fmt"
	"strconv"
	"strings"
)

type DataType interface {
	String() string
	Size() int
}

type VarCharType struct {
	n int
}

func (t VarCharType) String() string {
	return fmt.Sprintf("VARCHAR_%d_TYPE", t.n)
}

func (t VarCharType) Size() int {
	// store size integer + n bytes for string
	return t.n + 4
}

// getVarCharTypeSize returns the associated length with the given VarCharType string.
// if the string is not a valid VarCharType string, -1 is returned.
func getVarCharTypeLength(name string) int {
	parts := strings.Split(name, "_")
	if len(parts) != 3 {
		return -1
	}
	if parts[0] != "VARCHAR" || parts[2] != "TYPE" {
		return -1
	}
	n, err := strconv.Atoi(parts[1])
	if err != nil || n <= 0 {
		return -1
	}
	return n
}

type BoolType struct{}

var BOOL_TYPE BoolType // singleton

func (t BoolType) String() string {
	return "BOOL_TYPE"
}

func (t BoolType) Size() int {
	return 1
}

type IntType struct{}

var INT_TYPE IntType // singleton

func (t IntType) String() string {
	return "INT_TYPE"
}

func (t IntType) Size() int {
	return 4
}

func ToDataType(typeName string) (DataType, error) {
	switch typeName {
	case BOOL_TYPE.String():
		return BOOL_TYPE, nil
	case INT_TYPE.String():
		return INT_TYPE, nil
	default:
		n := getVarCharTypeLength(typeName)
		if n == -1 {
			return nil, fmt.Errorf("unknown DataType")
		}
		return VarCharType{n: n}, nil
	}
}
