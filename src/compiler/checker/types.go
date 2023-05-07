package checker

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type BFType interface {
	CheckAssign(value string) error
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
	FLOAT    BFPrimitive = "float"
	TEXT     BFPrimitive = "string"
)

func (bfType BFPrimitive) CheckAssign(value string) error {
	switch bfType {
	case BOOLEAN:
		if !(value == "true" || value == "false") {
			return errors.New("teste")
		}
	case BYTE:
		var _, err = strconv.ParseUint(value, 10, 8)
		return err
	case UNSIGNED:
		var _, err = strconv.ParseUint(value, 10, 0)
		return err
	case INTEGER:
		var _, err = strconv.ParseInt(value, 10, 0)
		return err
	case FLOAT:
		var _, err = strconv.ParseFloat(value, 64)
		return err
	case TEXT:
		if strings.Count(value, "\"") != 2 {
			return errors.New("teste")
		}
	}
	return nil
}

func (bfType BFPrimitive) String() string {
	return string(bfType)
}

type MapType struct {
	keyType   BFPrimitive
	valueType BFType
}

func (m MapType) CheckAssign(value string) error {
	if !regexp.MustCompile("map\\[.+].+\\{").MatchString(value) {
		return errors.New("out of range")
	}

	return nil
}

type ArrayType struct {
	capacity  int
	valueType BFType
}

func (m ArrayType) CheckAssign(value string) error {
	if !regexp.MustCompile("map\\[.+].+\\{").MatchString(value) {
		return errors.New("out of range")
	}
	return nil
}

func DefaultType(value string) (BFType, error) {
	if value == "true" || value == "false" {
		return BOOLEAN, nil
	}
	if maxValue, err := strconv.ParseInt(value, 10, 0); err != nil {
		return INTEGER, nil
	} else if maxValue > 0 {
		return UNKNOWN, errors.New("out of range")
	}
	if maxValue, err := strconv.ParseFloat(value, 64); err != nil {
		return FLOAT, nil
	} else if maxValue > 0 {
		return UNKNOWN, errors.New("out of range")
	}
	return MODULE, nil
}

func ParseType(value string) BFType {
	if strings.HasPrefix(value, "map") {
		return parseMapType(value)
	}
	if strings.HasPrefix(value, "[") {
		return parseArray(value)
	}
	return BFPrimitive(value)
}

func parseArray(value string) ArrayType {
	var capacity, _ = strconv.Atoi(string(value[1]))
	return ArrayType{
		capacity:  capacity,
		valueType: ParseType(value[3:]),
	}
}

func parseMapType(value string) MapType {
	var subMatches = regexp.MustCompile("(\\[.+])(.+)").FindStringSubmatch(value)
	var keyType, valueType = subMatches[1], subMatches[2]
	return MapType{
		keyType:   BFPrimitive(keyType),
		valueType: ParseType(valueType),
	}
}
