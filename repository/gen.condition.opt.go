package repository

import (
	"fmt"
	"reflect"
	"strings"
)

func (r *Repository) intCondition(fieldName string, fieldType string, rt reflect.Type, abbr string) []Condition {
	var conditions []Condition

	condition := fmt.Sprintf(`
func Condition%s(v ...%s) ConditionOption {
	return func(%s *%s) gen.Condition {
        if len(v) == 1 {
            return %s.q.%s.%s.Eq(v[0])
        }
        return %s.q.%s.%s.In(v...)
    }
}
`, fieldName, strings.Trim(fieldType, "*"), abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sNot(v ...%s) ConditionOption {
	return func(%s *%s) gen.Condition {
        if len(v) == 1 {
            return %s.q.%s.%s.Neq(v[0])
        }
        return %s.q.%s.%s.NotIn(v...)
    }
}
`, fieldName, strings.Trim(fieldType, "*"), abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sGt(v ...%s) ConditionOption {
	return func(%s *%s) gen.Condition {
        if len(v) == 0 {
            return %s.q.%s.%s.Gt(0)
        }
        return %s.q.%s.%s.Gt(v[0])
    }
}
`, fieldName, strings.Trim(fieldType, "*"), abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sGte(v ...%s) ConditionOption {
	return func(%s *%s) gen.Condition {
        if len(v) == 0 {
            return %s.q.%s.%s.Gte(0)
        }
        return %s.q.%s.%s.Gte(v[0])
    }
}
`, fieldName, strings.Trim(fieldType, "*"), abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sLt(v ...%s) ConditionOption {
	return func(%s *%s) gen.Condition {
        if len(v) == 0 {
            return %s.q.%s.%s.Lt(0)
        }
        return %s.q.%s.%s.Lt(v[0])
    }
}
`, fieldName, strings.Trim(fieldType, "*"), abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sLte(v ...%s) ConditionOption {
	return func(%s *%s) gen.Condition {
        if len(v) == 0 {
            return %s.q.%s.%s.Lte(0)
        }
        return %s.q.%s.%s.Lte(v[0])
    }
}
`, fieldName, strings.Trim(fieldType, "*"), abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sBetween(left, right %s) ConditionOption {
	return func(%s *%s) gen.Condition {
        return %s.q.%s.%s.Between(left, right)
    }
}
`, fieldName, strings.Trim(fieldType, "*"), abbr, rt.Name(), abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sNotBetween(left, right %s) ConditionOption {
	return func(%s *%s) gen.Condition {
        return %s.q.%s.%s.NotBetween(left, right)
    }
}
`, fieldName, strings.Trim(fieldType, "*"), abbr, rt.Name(), abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	if !strings.Contains(fieldType, "*") {
		return conditions
	}

	condition = fmt.Sprintf(`
func Condition%sIsNull() ConditionOption {
	return func(%s *%s) gen.Condition {
        return %s.q.%s.%s.IsNull()
    }
}
`, fieldName, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sIsNotNull() ConditionOption {
	return func(%s *%s) gen.Condition {
        return %s.q.%s.%s.IsNotNull()
    }
}
`, fieldName, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	return conditions
}

func (r *Repository) isInt(fieldType string) bool {
	if fieldType == "uint8" {
		return true
	}
	if fieldType == "uint" {
		return true
	}
	if fieldType == "uint16" {
		return true
	}
	if fieldType == "uint32" {
		return true
	}
	if fieldType == "uint64" {
		return true
	}
	if fieldType == "int8" {
		return true
	}
	if fieldType == "int" {
		return true
	}
	if fieldType == "int16" {
		return true
	}
	if fieldType == "int32" {
		return true
	}
	if fieldType == "int64" {
		return true
	}
	if fieldType == "*uint8" {
		return true
	}
	if fieldType == "*uint" {
		return true
	}
	if fieldType == "*uint16" {
		return true
	}
	if fieldType == "*uint32" {
		return true
	}
	if fieldType == "*uint64" {
		return true
	}
	if fieldType == "*int8" {
		return true
	}
	if fieldType == "*int" {
		return true
	}
	if fieldType == "*int16" {
		return true
	}
	if fieldType == "*int32" {
		return true
	}
	if fieldType == "*int64" {
		return true
	}
	return false
}

func (r *Repository) genConditionOpt(rt reflect.Type, abbr string) []Condition {
	var conditions []Condition
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		typ := field.Type.String()
		if r.isInt(typ) {
			conditions = append(conditions, r.intCondition(field.Name, typ, rt, abbr)...)
		}
	}
	return conditions
}
