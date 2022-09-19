package enum

import (
	"errors"
	"fmt"
)

type Enum struct {
	enum map[string]int
}

func NewEnum(args ...string) Enum {
	enum := make(map[string]int, 0)
	for index, val := range args {
		enum[val] = index
	}
	return Enum{
		enum: enum,
	}
}

func (e Enum) IsEnum(sy string) bool {
	_, ok := e.enum[sy]
	return ok
}

func (e Enum) IsEnums(sys ...string) bool {
	for i := 0; i < len(sys); i++ {
		if !e.IsEnum(sys[i]) {
			return false
		}
	}
	return true
}

func (e Enum) AppendEnum(sy string) error {
	if !e.IsEnum(sy) {
		e.enum[sy] = len(e.enum)
		return nil
	} else {
		return fmt.Errorf("duplicate data")
	}
}

func (e Enum) RemoveEnum(sy string) error {
	if e.IsEnum(sy) {
		delete(e.enum, sy)
		return nil
	} else {
		return fmt.Errorf("does not exist in the enum")
	}
}

func (e Enum) SetEnum(sy string, syv int) error {
	if !e.IsEnum(sy) {
		return errors.New("does not exist in the enum")
	}
	e.enum[sy] = syv
	return nil
}

func (e Enum) Enum(s string) (int, error) {
	return e.strQueryVal(s)
}

func (e Enum) Enums(s ...string) ([]int, error) {
	result := make([]int, 0)
	for i := 0; i < len(s); i++ {
		enumValue, err := e.strQueryVal(s[i])
		if err != nil {
			return nil, err
		}
		result = append(result, enumValue)
	}
	return result, nil
}

func (e Enum) strQueryVal(s string) (int, error) {
	val, ok := e.enum[s]
	if !ok {
		return 0, errors.New("does not exist in the enum")
	}
	return val, nil
}
