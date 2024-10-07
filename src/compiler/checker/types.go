package checker

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type BFType interface {
	CanAssign(value string, scope *BFScope) bool
}

type BFPrimitive string

const (
	UNKNOWN  BFPrimitive = ""
	MODULE   BFPrimitive = "module"
	MESSAGE  BFPrimitive = "message"
	BOOLEAN  BFPrimitive = "bool"
	BYTE     BFPrimitive = "byte"
	INTEGER  BFPrimitive = "int"
	UNSIGNED BFPrimitive = "uint"
	FLOAT    BFPrimitive = "float64"
	TEXT     BFPrimitive = "string"
	ANY      BFPrimitive = "any"
)

var assignTesters = map[BFPrimitive]func(string, *BFScope) bool{
	TEXT: func(value string, scope *BFScope) bool {
		var items = strings.Split(value, "+")
		for _, item := range items {
			var trimmed = strings.TrimSpace(item)
			if !identifierHasType(trimmed, TEXT, scope) {
				return false
			}
			if regexp.MustCompile("^\"([^\"]|\\\")\"$").MatchString(trimmed) {
				return false
			}
		}
		return true
	},

	BOOLEAN: func(value string, scope *BFScope) bool {
		if value[0] == '(' && value[len(value)-1] == ')' {
			value = value[1 : len(value)-1]
		}
		if regexp.MustCompile("[><]=|>|<").MatchString(value) {
			return true
		}
		if value == "true" || value == "false" {
			return true
		}
		var items = regexp.MustCompile("&&|[|]{2}").Split(value, -1)
		fmt.Print(items)
		return false
	},
}

func (bfType BFPrimitive) CanAssign(value string, scope *BFScope) bool {
	return identifierHasType(value, bfType, scope) || assignTesters[bfType](value, scope)
	switch bfType {

	case BOOLEAN:

	case BYTE:
		var _, err = strconv.ParseUint(value, 10, 8)
		return err != nil
	case UNSIGNED:
		var _, err = strconv.ParseUint(value, 10, 0)
		return err != nil
	case INTEGER:
		var _, err = strconv.ParseInt(value, 10, 0)
		return err != nil
	case FLOAT:
		var _, err = strconv.ParseFloat(value, 64)
		return err != nil
	}
	return false
}

func (bfType BFPrimitive) String() string {
	return string(bfType)
}

type MapType struct {
	keyType   BFPrimitive
	valueType BFType
}

func (m MapType) CheckAssign(value string, scope *BFScope) error {
	if !regexp.MustCompile("map\\[.+].+\\{").MatchString(value) {
		return errors.New("out of range")
	}

	return nil
}

type ArrayType struct {
	capacity  int
	valueType BFType
}

func (m ArrayType) CanAssign(value string, scope *BFScope) bool {
	if !regexp.MustCompile("map\\[.+].+\\{").MatchString(value) {
		return false //errors.New("out of range")
	}
	return true
}

func DefaultType(value string, scope *BFScope) (BFType, error) {

	fmt.Println(value)
	if isIdentifier(value) {
		var symbol = scope.GetSymbol(value)
		if symbol != nil {
			return symbol.Type(), nil
		}
		return nil, errors.New("Identifier not found")
	}
	if value[0] == '[' && value[len(value)-1] == ']' {
		var capacity = len(strings.Split(value, ","))
		return ArrayType{
			capacity:  capacity,
			valueType: ANY,
		}, nil
	}
	if value[0] == '{' && value[len(value)-1] == '}' {
		var items = strings.Split(value, ",")
		var valueType = checkItemsType(items)
		return ArrayType{
			capacity:  len(items),
			valueType: valueType,
		}, nil
	}
	return checkPrimitiveType(value, scope)
	// if value == "true" || value == "false" {
	// 	return BOOLEAN, nil
	// }
	// if maxValue, err := strconv.ParseInt(value, 10, 0); err == nil {
	// 	return INTEGER, nil
	// } else if maxValue > 0 {
	// 	return UNKNOWN, errors.New("integer out of range")
	// }
	// if maxValue, err := strconv.ParseFloat(value, 64); err == nil {
	// 	return FLOAT, nil
	// } else if maxValue > 0 {
	// 	return UNKNOWN, errors.New("float out of range")
	// }
	// if regexp.MustCompile("^\".*\"$").MatchString(value) {
	// 	return TEXT, nil
	// }

	// return nil, errors.New("Unknown Type")
}

func checkItemsType(items []string) BFType {
	var reference = items[0]
	var bfType, _ = DefaultType(reference, nil)
	for _, item := range items[1:] {
		if bfType.CanAssign(item, nil) {
			continue
		}
		if newType, _ := DefaultType(item, nil); newType.CanAssign(reference, nil) {
			reference, bfType = item, newType
		}
		bfType = ANY
	}
	return bfType
}

func isIdentifier(value string) bool {
	return regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$").MatchString(value)
}

func identifierHasType(value string, bfType BFType, scope *BFScope) bool {
	if !isIdentifier(value) {
		return false
	}
	var symbol = scope.GetSymbol(value)
	return symbol != nil && symbol.Type() == bfType
}

func checkPrimitiveType(value string, scope *BFScope) (BFPrimitive, error) {
	if BOOLEAN.CanAssign(value, scope) {
		return BOOLEAN, nil
	}
	if INTEGER.CanAssign(value, scope) {
		return INTEGER, nil
	}
	if FLOAT.CanAssign(value, scope) {
		return FLOAT, nil
	}
	if TEXT.CanAssign(value, scope) {
		return TEXT, nil
	}
	return UNKNOWN, errors.New("")
}
