package repository

import (
	"fmt"
	"reflect"
	"strings"
)

func (r *Repository) intUpdate(fieldName string, fieldType string, rt reflect.Type, abbr string) []Update {
	var updates []Update

	update := fmt.Sprintf(`
func Update%sAdd(v %s) UpdateOption {
	return func(%s *%s) field.AssignExpr {
        return %s.q.%s.%s.Add(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	updates = append(updates, Update(update))

	update = fmt.Sprintf(`
func Update%sSub(v %s) UpdateOption {
	return func(%s *%s) field.AssignExpr {
        return %s.q.%s.%s.Sub(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	updates = append(updates, Update(update))

	update = fmt.Sprintf(`
func Update%sMul(v %s) UpdateOption {
	return func(%s *%s) field.AssignExpr {
        return %s.q.%s.%s.Mul(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	updates = append(updates, Update(update))

	update = fmt.Sprintf(`
func Update%sDiv(v %s) UpdateOption {
	return func(%s *%s) field.AssignExpr {
        return %s.q.%s.%s.Div(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	updates = append(updates, Update(update))

	return updates
}

func (r *Repository) decimalUpdate(fieldName string, fieldType string, rt reflect.Type, abbr string) []Update {
	var updates []Update

	update := fmt.Sprintf(`
func Update%s(v %s) UpdateOption {
	return func(%s *%s) field.AssignExpr {
        return f.NewDecimal(%s.q.%s.%s).Value(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	updates = append(updates, Update(update))

	update = fmt.Sprintf(`
func Update%sAdd(v %s) UpdateOption {
	return func(%s *%s) field.AssignExpr {
        return f.NewDecimal(%s.q.%s.%s).Add(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	updates = append(updates, Update(update))

	update = fmt.Sprintf(`
func Update%sSub(v %s) UpdateOption {
	return func(%s *%s) field.AssignExpr {
        return f.NewDecimal(%s.q.%s.%s).Sub(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	updates = append(updates, Update(update))

	update = fmt.Sprintf(`
func Update%sMul(v %s) UpdateOption {
	return func(%s *%s) field.AssignExpr {
        return f.NewDecimal(%s.q.%s.%s).Mul(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	updates = append(updates, Update(update))

	update = fmt.Sprintf(`
func Update%sDiv(v %s) UpdateOption {
	return func(%s *%s) field.AssignExpr {
        return f.NewDecimal(%s.q.%s.%s).Div(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name(), abbr, rt.Name(), fieldName)
	updates = append(updates, Update(update))

	return updates
}

func (r *Repository) genUpdateOpt(rt reflect.Type, abbr string) (updates []Update, timePkg, decimalPkg, numberDecimalPkg bool) {
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if strings.ToLower(field.Name) == "id" {
			continue
		}
		typ := field.Type.String()
		if !r.allowType(typ) {
			continue
		}
		fieldType := strings.Trim(field.Type.String(), "*")

		if !r.isDecimal(typ) {
			update := fmt.Sprintf(`
func Update%s(v %s) UpdateOption {
	return func(%s *%s) field.AssignExpr {
        return %s.q.%s.%s.Value(v)
    }
}
`, field.Name, fieldType, abbr, rt.Name(), abbr, rt.Name(), field.Name)
			updates = append(updates, Update(update))
		}

		if r.isInt(typ) {
			updates = append(updates, r.intUpdate(field.Name, fieldType, rt, abbr)...)
		}
		if r.isDecimal(typ) {
			decimalPkg = true
			numberDecimalPkg = true
			updates = append(updates, r.decimalUpdate(field.Name, fieldType, rt, abbr)...)
		}

		if !strings.Contains(typ, "*") {
			continue
		}

		update := fmt.Sprintf(`
func Update%sNull() UpdateOption {
	return func(%s *%s) field.AssignExpr {
        return %s.q.%s.%s.Null()
    }
}
`, field.Name, abbr, rt.Name(), abbr, rt.Name(), field.Name)
		updates = append(updates, Update(update))
	}

	return
}
