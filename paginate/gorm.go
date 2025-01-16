package paginate

import (
	"gorm.io/gorm"
)

// Gorm 分页
func Gorm[T Int](page, pageSize T) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		return db.Offset((int(page) - 1) * int(pageSize)).Limit(int(pageSize))
	}
}
