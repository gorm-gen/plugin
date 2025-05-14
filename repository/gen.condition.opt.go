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
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
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
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
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
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
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
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
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
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
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
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sBetween(left, right %s) ConditionOption {
	return func(%s *%s) gen.Condition {
        return %s.q.%s.%s.Between(left, right)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sNotBetween(left, right %s) ConditionOption {
	return func(%s *%s) gen.Condition {
        return %s.q.%s.%s.NotBetween(left, right)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
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

func (r *Repository) stringCondition(fieldName string, fieldType string, rt reflect.Type, abbr string) []Condition {
	var conditions []Condition

	condition := fmt.Sprintf(`
func Condition%sEq(v %s) ConditionOption {
	return func(%s *%s) gen.Condition {
        return %s.q.%s.%s.Eq(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sNeq(v %s) ConditionOption {
	return func(%s *%s) gen.Condition {
        return %s.q.%s.%s.Neq(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sLike(v %s) ConditionOption {
	return func(%s *%s) gen.Condition {
        return %s.q.%s.%s.Like(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sNotLike(v %s) ConditionOption {
	return func(%s *%s) gen.Condition {
        return %s.q.%s.%s.NotLike(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	return conditions
}

func (r *Repository) isString(fieldType string) bool {
	if fieldType == "string" {
		return true
	}
	if fieldType == "*string" {
		return true
	}
	return false
}

func (r *Repository) timeCondition(fieldName string, fieldType string, rt reflect.Type, abbr string) []Condition {
	var conditions []Condition

	condition := fmt.Sprintf(`
func Condition%sEq(v ...%s) ConditionOption {
	return func(%s *%s) gen.Condition {
        if len(v) > 0 && !time.IsZero(v[0]) {
            return %s.q.%s.%s.Eq(v[0])
        }
        return %s.q.%s.%s.Eq(time.Now())
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sNeq(v ...%s) ConditionOption {
	return func(%s *%s) gen.Condition {
        if len(v) > 0 && !time.IsZero(v[0]) {
            return %s.q.%s.%s.Neq(v[0])
        }
        return %s.q.%s.%s.Neq(time.Now())
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sGt(v ...%s) ConditionOption {
	return func(%s *%s) gen.Condition {
        if len(v) > 0 && !time.IsZero(v[0]) {
            return %s.q.%s.%s.Gt(v[0])
        }
        return %s.q.%s.%s.Gt(time.Now())
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sGte(v ...%s) ConditionOption {
	return func(%s *%s) gen.Condition {
        if len(v) > 0 && !time.IsZero(v[0]) {
            return %s.q.%s.%s.Gte(v[0])
        }
        return %s.q.%s.%s.Gte(time.Now())
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sLt(v ...%s) ConditionOption {
	return func(%s *%s) gen.Condition {
        if len(v) > 0 && !time.IsZero(v[0]) {
            return %s.q.%s.%s.Lt(v[0])
        }
        return %s.q.%s.%s.Lt(time.Now())
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sLte(v ...%s) ConditionOption {
	return func(%s *%s) gen.Condition {
        if len(v) > 0 && !time.IsZero(v[0]) {
            return %s.q.%s.%s.Lte(v[0])
        }
        return %s.q.%s.%s.Lte(time.Now())
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sBetween(left, right %s) ConditionOption {
	return func(%s *%s) gen.Condition {
        return %s.q.%s.%s.Between(left, right)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sNotBetween(left, right %s) ConditionOption {
	return func(%s *%s) gen.Condition {
        return %s.q.%s.%s.NotBetween(left, right)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))
	return conditions
}

func (r *Repository) isTime(fieldType string) bool {
	if fieldType == "time.Time" {
		return true
	}
	if fieldType == "*time.Time" {
		return true
	}
	return false
}

func (r *Repository) decimalCondition(fieldName string, fieldType string, rt reflect.Type, abbr string) []Condition {
	var conditions []Condition

	condition := fmt.Sprintf(`
func Condition%sEq(v %s) ConditionOption {
	return func(%s *%s) gen.Condition {
        return %s.q.%s.%s.Eq(value.NewDecimal(v))
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sNeq(v %s) ConditionOption {
	return func(%s *%s) gen.Condition {
        return %s.q.%s.%s.Neq(value.NewDecimal(v))
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sGt(v ...%s) ConditionOption {
	return func(%s *%s) gen.Condition {
        if len(v) == 0 {
            return %s.q.%s.%s.Gt(value.NewDecimal(decimal.Zero))
        }
        return %s.q.%s.%s.Gt(value.NewDecimal(v[0]))
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sGte(v ...%s) ConditionOption {
	return func(%s *%s) gen.Condition {
        if len(v) == 0 {
            return %s.q.%s.%s.Gte(value.NewDecimal(decimal.Zero))
        }
        return %s.q.%s.%s.Gte(value.NewDecimal(v[0]))
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sLt(v ...%s) ConditionOption {
	return func(%s *%s) gen.Condition {
        if len(v) == 0 {
            return %s.q.%s.%s.Lt(value.NewDecimal(decimal.Zero))
        }
        return %s.q.%s.%s.Lt(value.NewDecimal(v[0]))
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sLte(v ...%s) ConditionOption {
	return func(%s *%s) gen.Condition {
        if len(v) == 0 {
            return %s.q.%s.%s.Lte(value.NewDecimal(decimal.Zero))
        }
        return %s.q.%s.%s.Lte(value.NewDecimal(v[0]))
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName, abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sBetween(left, right %s) ConditionOption {
	return func(%s *%s) gen.Condition {
        return f.NewDecimal(%s.q.%s.%s).Between(left, right)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%sNotBetween(left, right %s) ConditionOption {
	return func(%s *%s) gen.Condition {
        return f.NewDecimal(%s.q.%s.%s).NotBetween(left, right)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	conditions = append(conditions, Condition(condition))
	return conditions
}

func (r *Repository) isDecimal(fieldType string) bool {
	if fieldType == "decimal.Decimal" {
		return true
	}
	if fieldType == "*decimal.Decimal" {
		return true
	}
	return false
}

func (r *Repository) allowType(fieldType string) bool {
	if r.isInt(fieldType) {
		return true
	}
	if r.isDecimal(fieldType) {
		return true
	}
	if r.isString(fieldType) {
		return true
	}
	if r.isTime(fieldType) {
		return true
	}
	return false
}

func (r *Repository) genConditionOpt(rt reflect.Type, abbr string) (conditions []Condition, timePkg, decimalPkg, numberDecimalPkg bool) {
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		typ := field.Type.String()
		if !r.allowType(typ) {
			continue
		}
		fieldType := strings.Trim(field.Type.String(), "*")
		if r.isInt(typ) {
			conditions = append(conditions, r.intCondition(field.Name, fieldType, rt, abbr)...)
		}
		if r.isString(typ) {
			conditions = append(conditions, r.stringCondition(field.Name, fieldType, rt, abbr)...)
		}
		if r.isTime(typ) {
			timePkg = true
			conditions = append(conditions, r.timeCondition(field.Name, fieldType, rt, abbr)...)
		}
		if r.isDecimal(typ) {
			decimalPkg = true
			numberDecimalPkg = true
			conditions = append(conditions, r.decimalCondition(field.Name, fieldType, rt, abbr)...)
		}

		if !strings.Contains(typ, "*") {
			continue
		}

		condition := fmt.Sprintf(`
func Condition%sIsNull() ConditionOption {
	return func(%s *%s) gen.Condition {
        return %s.q.%s.%s.IsNull()
    }
}
`, field.Name, abbr, rt.Name(), abbr, rt.Name(), field.Name)
		conditions = append(conditions, Condition(condition))

		condition = fmt.Sprintf(`
func Condition%sIsNotNull() ConditionOption {
	return func(%s *%s) gen.Condition {
        return %s.q.%s.%s.IsNotNull()
    }
}
`, field.Name, abbr, rt.Name(), abbr, rt.Name(), field.Name)
		conditions = append(conditions, Condition(condition))
	}

	return
}
