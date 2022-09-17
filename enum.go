package enum

import (
	"errors"
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

func (e Enum) AppendEnum(sy string) {
	if !e.IsEnum(sy) {
		e.enum[sy] = len(e.enum) + 1
	}
}

func (e Enum) RemoveEnum(sy string) {
	if e.IsEnum(sy) {
		delete(e.enum, sy)
	}
}

func (e Enum) SetEnum(sy string, syv int) {
	e.enum[sy] = syv
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
