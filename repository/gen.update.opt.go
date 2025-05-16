package repository

import (
	"fmt"
	"reflect"
	"strings"
)

func (r *Repository) intUpdate(fieldName string, fieldType string, rt reflect.Type, abbr string) []Update {
	var updates []Update

	update := fmt.Sprintf(`
func Update%[1]sAdd(v %[2]s) UpdateOption {
	return func(%[3]s *%[4]s) field.AssignExpr {
        return %[3]s.q.%[4]s.%[1]s.Add(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	updates = append(updates, Update(update))

	update = fmt.Sprintf(`
func Update%[1]sSub(v %[2]s) UpdateOption {
	return func(%[3]s *%[4]s) field.AssignExpr {
        return %[3]s.q.%[4]s.%[1]s.Sub(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	updates = append(updates, Update(update))

	update = fmt.Sprintf(`
func Update%[1]sMul(v %[2]s) UpdateOption {
	return func(%[3]s *%[4]s) field.AssignExpr {
        return %[3]s.q.%[4]s.%[1]s.Mul(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	updates = append(updates, Update(update))

	update = fmt.Sprintf(`
func Update%[1]sDiv(v %[2]s) UpdateOption {
	return func(%[3]s *%[4]s) field.AssignExpr {
        return %[3]s.q.%[4]s.%[1]s.Div(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	updates = append(updates, Update(update))

	return updates
}

func (r *Repository) decimalUpdate(fieldName string, fieldType string, rt reflect.Type, abbr string) []Update {
	var updates []Update

	update := fmt.Sprintf(`
func Update%[1]s(v %[2]s) UpdateOption {
	return func(%[3]s *%[4]s) field.AssignExpr {
        return f.NewDecimal(%[3]s.q.%[4]s.%[1]s).Value(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	updates = append(updates, Update(update))

	update = fmt.Sprintf(`
func Update%[1]sAdd(v %[2]s) UpdateOption {
	return func(%[3]s *%[4]s) field.AssignExpr {
        return f.NewDecimal(%[3]s.q.%[4]s.%[1]s).Add(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	updates = append(updates, Update(update))

	update = fmt.Sprintf(`
func Update%[1]sSub(v %[2]s) UpdateOption {
	return func(%[3]s *%[4]s) field.AssignExpr {
        return f.NewDecimal(%[3]s.q.%[4]s.%[1]s).Sub(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	updates = append(updates, Update(update))

	update = fmt.Sprintf(`
func Update%[1]sMul(v %[2]s) UpdateOption {
	return func(%[3]s *%[4]s) field.AssignExpr {
        return f.NewDecimal(%[3]s.q.%[4]s.%[1]s).Mul(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
	updates = append(updates, Update(update))

	update = fmt.Sprintf(`
func Update%[1]sDiv(v %[2]s) UpdateOption {
	return func(%[3]s *%[4]s) field.AssignExpr {
        return f.NewDecimal(%[3]s.q.%[4]s.%[1]s).Div(v)
    }
}
`, fieldName, fieldType, abbr, rt.Name())
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
func Update%[1]s(v %[2]s) UpdateOption {
	return func(%[3]s *%[4]s) field.AssignExpr {
        return %[3]s.q.%[4]s.%[1]s.Value(v)
    }
}
`, field.Name, fieldType, abbr, rt.Name())
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
func Update%[1]sNull() UpdateOption {
	return func(%[2]s *%[3]s) field.AssignExpr {
        return %[2]s.q.%[3]s.%[1]s.Null()
    }
}
`, field.Name, abbr, rt.Name())
		updates = append(updates, Update(update))
	}

	return
}
