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
func Order%sAsc() OrderOption {
	return func(%s *%s) field.Expr {
        return %s.q.%s.%s.Asc()
    }
}
`, field.Name, abbr, rt.Name(), abbr, rt.Name(), field.Name)
		orders = append(orders, Order(order))

		order = fmt.Sprintf(`
func Order%sDesc() OrderOption {
	return func(%s *%s) field.Expr {
        return %s.q.%s.%s.Desc()
    }
}
`, field.Name, abbr, rt.Name(), abbr, rt.Name(), field.Name)
		orders = append(orders, Order(order))
	}

	return orders
}
