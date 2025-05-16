package repository

import (
	"fmt"
	"reflect"
	"strings"
)

func (r *Repository) allowOrderField(fieldName string) bool {
	fn := strings.ToLower(fieldName)
	if fn == "id" {
		return true
	}
	if fn == "created_at" {
		return true
	}
	if fn == "updated_at" {
		return true
	}
	if fieldName == "CreatedAt" {
		return true
	}
	if fieldName == "UpdatedAt" {
		return true
	}
	return false
}

func (r *Repository) genOrderOpt(rt reflect.Type, abbr string) []Order {
	orders := make([]Order, 0)
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)

		if !r.allowOrderField(field.Name) {
			continue
		}

		order := fmt.Sprintf(`
func Order%[1]sAsc() OrderOption {
	return func(%[2]s *%[3]s) field.Expr {
        if %[2]s.newTableName != nil {
            return %[2]s.q.%[3]s.Table(*%[2]s.newTableName).%[1]s.Asc()
        }
        return %[2]s.q.%[3]s.%[1]s.Asc()
    }
}
`, field.Name, abbr, rt.Name())
		orders = append(orders, Order(order))

		order = fmt.Sprintf(`
func Order%[1]sDesc() OrderOption {
	return func(%[2]s *%[3]s) field.Expr {
        if %[2]s.newTableName != nil {
            return %[2]s.q.%[3]s.Table(*%[2]s.newTableName).%[1]s.Desc()
        }
        return %[2]s.q.%[3]s.%[1]s.Desc()
    }
}
`, field.Name, abbr, rt.Name())
		orders = append(orders, Order(order))
	}

	return orders
}
