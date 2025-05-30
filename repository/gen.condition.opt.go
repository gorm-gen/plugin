package repository

import (
	"fmt"
	"reflect"
	"strings"
)

func (r *Repository) intCondition(fieldName, fieldType string, rt reflect.Type, abbr string) []Condition {
	var conditions []Condition

	condition := fmt.Sprintf(`
func Condition%[1]s(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        length := len(v)
        if %[3]s.newTableName != nil {
            if length == 0 {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Eq(0)
            }
            if length == 1 {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Eq(v[0])
            }
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.In(v...)
        }
        if length == 0 {
            return %[3]s.q.%[4]s.%[1]s.Eq(0)
        }
        if length == 1 {
            return %[3]s.q.%[4]s.%[1]s.Eq(v[0])
        }
        return %[3]s.q.%[4]s.%[1]s.In(v...)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sNot(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        length := len(v)
        if %[3]s.newTableName != nil {
            if length == 0 {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Neq(0)
            }
            if length == 1 {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Neq(v[0])
            }
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.NotIn(v...)
        }
        if length == 0 {
            return %[3]s.q.%[4]s.%[1]s.Neq(0)
        }
        if length == 1 {
            return %[3]s.q.%[4]s.%[1]s.Neq(v[0])
        }
        return %[3]s.q.%[4]s.%[1]s.NotIn(v...)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sGt(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            if len(v) == 0 {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Gt(0)
            }
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Gt(v[0])
        }
        if len(v) == 0 {
            return %[3]s.q.%[4]s.%[1]s.Gt(0)
        }
        return %[3]s.q.%[4]s.%[1]s.Gt(v[0])
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sGte(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            if len(v) == 0 {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Gte(0)
            }
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Gte(v[0])
        }
        if len(v) == 0 {
            return %[3]s.q.%[4]s.%[1]s.Gte(0)
        }
        return %[3]s.q.%[4]s.%[1]s.Gte(v[0])
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sLt(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            if len(v) == 0 {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Lt(0)
            }
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Lt(v[0])
        }
        if len(v) == 0 {
            return %[3]s.q.%[4]s.%[1]s.Lt(0)
        }
        return %[3]s.q.%[4]s.%[1]s.Lt(v[0])
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sLte(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            if len(v) == 0 {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Lte(0)
            }
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Lte(v[0])
        }
        if len(v) == 0 {
            return %[3]s.q.%[4]s.%[1]s.Lte(0)
        }
        return %[3]s.q.%[4]s.%[1]s.Lte(v[0])
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sBetween(left, right %[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Between(left, right)
        }
        return %[3]s.q.%[4]s.%[1]s.Between(left, right)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sNotBetween(left, right %[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.NotBetween(left, right)
        }
        return %[3]s.q.%[4]s.%[1]s.NotBetween(left, right)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
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

func (r *Repository) stringCondition(fieldName, fieldType string, rt reflect.Type, abbr string) []Condition {
	var conditions []Condition

	condition := fmt.Sprintf(`
func Condition%[1]s(v %[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Eq(v)
        }
        return %[3]s.q.%[4]s.%[1]s.Eq(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sNeq(v %[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Neq(v)
        }
        return %[3]s.q.%[4]s.%[1]s.Neq(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sLike(v %[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Like(v)
        }
        return %[3]s.q.%[4]s.%[1]s.Like(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sNotLike(v %[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.NotLike(v)
        }
        return %[3]s.q.%[4]s.%[1]s.NotLike(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
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

func (r *Repository) timeCondition(fieldName, fieldType string, rt reflect.Type, abbr string) []Condition {
	var conditions []Condition

	condition := fmt.Sprintf(`
func Condition%[1]s(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            if len(v) > 0 && !v[0].IsZero() {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Eq(v[0])
            }
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Eq(time.Now())
        }
        if len(v) > 0 && !v[0].IsZero() {
            return %[3]s.q.%[4]s.%[1]s.Eq(v[0])
        }
        return %[3]s.q.%[4]s.%[1]s.Eq(time.Now())
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sNeq(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            if len(v) > 0 && !v[0].IsZero() {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Neq(v[0])
            }
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Neq(time.Now())
        }
        if len(v) > 0 && !v[0].IsZero() {
            return %[3]s.q.%[4]s.%[1]s.Neq(v[0])
        }
        return %[3]s.q.%[4]s.%[1]s.Neq(time.Now())
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sGt(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            if len(v) > 0 && !v[0].IsZero() {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Gt(v[0])
            }
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Gt(time.Now())
        }
        if len(v) > 0 && !v[0].IsZero() {
            return %[3]s.q.%[4]s.%[1]s.Gt(v[0])
        }
        return %[3]s.q.%[4]s.%[1]s.Gt(time.Now())
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sGte(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            if len(v) > 0 && !v[0].IsZero() {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Gte(v[0])
            }
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Gte(time.Now())
        }
        if len(v) > 0 && !v[0].IsZero() {
            return %[3]s.q.%[4]s.%[1]s.Gte(v[0])
        }
        return %[3]s.q.%[4]s.%[1]s.Gte(time.Now())
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sLt(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            if len(v) > 0 && !v[0].IsZero() {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Lt(v[0])
            }
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Lt(time.Now())
        }
        if len(v) > 0 && !v[0].IsZero() {
            return %[3]s.q.%[4]s.%[1]s.Lt(v[0])
        }
        return %[3]s.q.%[4]s.%[1]s.Lt(time.Now())
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sLte(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            if len(v) > 0 && !v[0].IsZero() {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Lte(v[0])
            }
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Lte(time.Now())
        }
        if len(v) > 0 && !v[0].IsZero() {
            return %[3]s.q.%[4]s.%[1]s.Lte(v[0])
        }
        return %[3]s.q.%[4]s.%[1]s.Lte(time.Now())
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sBetween(left, right %[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Between(left, right)
        }
        return %[3]s.q.%[4]s.%[1]s.Between(left, right)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sNotBetween(left, right %[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.NotBetween(left, right)
        }
        return %[3]s.q.%[4]s.%[1]s.NotBetween(left, right)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
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

func (r *Repository) boolCondition(fieldName, fieldType string, rt reflect.Type, abbr string) []Condition {
	var conditions []Condition

	condition := fmt.Sprintf(`
func Condition%[1]sIs(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        _v := true
        if len(v) > 0 {
            _v = v[0]
        }
        if %[3]s.newTableName != nil {
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Is(_v)
        }
        return %[3]s.q.%[4]s.%[1]s.Is(_v)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))
	return conditions
}

func (r *Repository) isBool(fieldType string) bool {
	if fieldType == "bool" {
		return true
	}
	if fieldType == "*bool" {
		return true
	}
	return false
}

func (r *Repository) decimalCondition(fieldName, fieldType string, rt reflect.Type, abbr string) []Condition {
	var conditions []Condition

	condition := fmt.Sprintf(`
func Condition%[1]s(v %[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Eq(value.NewDecimal(v))
        }
        return %[3]s.q.%[4]s.%[1]s.Eq(value.NewDecimal(v))
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sNeq(v %[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Neq(value.NewDecimal(v))
        }
        return %[3]s.q.%[4]s.%[1]s.Neq(value.NewDecimal(v))
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sGt(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            if len(v) == 0 {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Gt(value.NewDecimal(decimal.Zero))
            }
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Gt(value.NewDecimal(v[0]))
        }
        if len(v) == 0 {
            return %[3]s.q.%[4]s.%[1]s.Gt(value.NewDecimal(decimal.Zero))
        }
        return %[3]s.q.%[4]s.%[1]s.Gt(value.NewDecimal(v[0]))
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sGte(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            if len(v) == 0 {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Gte(value.NewDecimal(decimal.Zero))
            }
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Gte(value.NewDecimal(v[0]))
        }
        if len(v) == 0 {
            return %[3]s.q.%[4]s.%[1]s.Gte(value.NewDecimal(decimal.Zero))
        }
        return %[3]s.q.%[4]s.%[1]s.Gte(value.NewDecimal(v[0]))
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sLt(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            if len(v) == 0 {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Lt(value.NewDecimal(decimal.Zero))
            }
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Lt(value.NewDecimal(v[0]))
        }
        if len(v) == 0 {
            return %[3]s.q.%[4]s.%[1]s.Lt(value.NewDecimal(decimal.Zero))
        }
        return %[3]s.q.%[4]s.%[1]s.Lt(value.NewDecimal(v[0]))
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sLte(v ...%[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            if len(v) == 0 {
                return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Lte(value.NewDecimal(decimal.Zero))
            }
            return %[3]s.q.%[4]s.Table(*%[3]s.newTableName).%[1]s.Lte(value.NewDecimal(v[0]))
        }
        if len(v) == 0 {
            return %[3]s.q.%[4]s.%[1]s.Lte(value.NewDecimal(decimal.Zero))
        }
        return %[3]s.q.%[4]s.%[1]s.Lte(value.NewDecimal(v[0]))
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sBetween(left, right %[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            return f.NewDecimal(%[3]s.q.%[4]s.%[1]s, f.WithTableName(*%[3]s.newTableName)).Between(left, right)
        }
        return f.NewDecimal(%[3]s.q.%[4]s.%[1]s).Between(left, right)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sNotBetween(left, right %[2]s) ConditionOption {
	return func(%[3]s *%[4]s) gen.Condition {
        if %[3]s.newTableName != nil {
            return f.NewDecimal(%[3]s.q.%[4]s.%[1]s, f.WithTableName(*%[3]s.newTableName)).NotBetween(left, right)
        }
        return f.NewDecimal(%[3]s.q.%[4]s.%[1]s).NotBetween(left, right)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
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

func (r *Repository) deletedCondition(fieldName string, rt reflect.Type, abbr string) []Condition {
	var conditions []Condition

	condition := fmt.Sprintf(`
func Condition%[1]sIsZero() ConditionOption {
	return func(%[2]s *%[3]s) gen.Condition {
        if %[2]s.newTableName != nil {
            return f.NewDecimal(%[2]s.q.%[3]s.%[1]s, f.WithTableName(*%[2]s.newTableName)).Eq(decimal.Zero)
        }
        return f.NewDecimal(%[2]s.q.%[3]s.%[1]s).Eq(decimal.Zero)
    }
}
`, fieldName, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	condition = fmt.Sprintf(`
func Condition%[1]sGtZero() ConditionOption {
	return func(%[2]s *%[3]s) gen.Condition {
        if %[2]s.newTableName != nil {
            return f.NewDecimal(%[2]s.q.%[3]s.%[1]s, f.WithTableName(*%[2]s.newTableName)).Gt(decimal.Zero)
        }
        return f.NewDecimal(%[2]s.q.%[3]s.%[1]s).Gt(decimal.Zero)
    }
}
`, fieldName, abbr, rt.Name())
	conditions = append(conditions, Condition(condition))

	return conditions
}

func (r *Repository) isDeleted(fieldType string) bool {
	if fieldType == "soft_delete.DeletedAt" {
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
	if r.isBool(fieldType) {
		return true
	}
	return false
}

func (r *Repository) allowConditionType(fieldType string) bool {
	if fieldType == "soft_delete.DeletedAt" {
		return true
	}
	if fieldType == "gorm.DeletedAt" {
		return true
	}
	return r.allowType(fieldType)
}

func (r *Repository) genConditionOpt(rt reflect.Type, abbr string) (conditions []Condition, timePkg, decimalPkg, pfPkg, pfvPkg, reflectPkg bool) {
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		typ := field.Type.String()
		if !r.allowConditionType(typ) {
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
			pfPkg = true
			pfvPkg = true
			conditions = append(conditions, r.decimalCondition(field.Name, fieldType, rt, abbr)...)
		}
		if r.isBool(typ) {
			conditions = append(conditions, r.boolCondition(field.Name, fieldType, rt, abbr)...)
		}
		if r.isDeleted(typ) {
			decimalPkg = true
			pfPkg = true
			conditions = append(conditions, r.deletedCondition(field.Name, rt, abbr)...)
		}

		if !strings.Contains(typ, "*") && fieldType != "gorm.DeletedAt" {
			continue
		}

		condition := fmt.Sprintf(`
func Condition%[1]sIsNull() ConditionOption {
	return func(%[2]s *%[3]s) gen.Condition {
        if %[2]s.newTableName != nil {
            return %[2]s.q.%[3]s.Table(*%[2]s.newTableName).%[1]s.IsNull()
        }
        return %[2]s.q.%[3]s.%[1]s.IsNull()
    }
}
`, field.Name, abbr, rt.Name())
		conditions = append(conditions, Condition(condition))

		condition = fmt.Sprintf(`
func Condition%[1]sIsNotNull() ConditionOption {
	return func(%[2]s *%[3]s) gen.Condition {
        if %[2]s.newTableName != nil {
            return %[2]s.q.%[3]s.Table(*%[2]s.newTableName).%[1]s.IsNotNull()
        }
        return %[2]s.q.%[3]s.%[1]s.IsNotNull()
    }
}
`, field.Name, abbr, rt.Name())
		conditions = append(conditions, Condition(condition))
	}

	return
}
