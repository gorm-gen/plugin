package paginate

import (
	"gorm.io/gen"
)

// Gen 分页
func Gen[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](page, pageSize T) func(gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		if page == 0 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		return dao.Offset((int(page) - 1) * int(pageSize)).Limit(int(pageSize))
	}
}
