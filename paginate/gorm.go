package paginate

import (
	"gorm.io/gorm"
)

// Gorm 分页
func Gorm[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](page, pageSize T) func(*gorm.DB) *gorm.DB {
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
