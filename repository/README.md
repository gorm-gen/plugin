### 使用教程
```go
package main

import (
	"fmt"

	"github.com/gorm-gen/plugin/repository"
	
	"demo/internal/models"
)

func main() {
	r := repository.New(
		repository.WithModule("demo"),
		repository.WithRepositoryPath("cmd/internal/repositories"),
		repository.WithGenQueryPkg("demo/internal/query"),
		repository.WithGormDBVar("global.DB"),
		repository.WithGormDBVarPkg("demo/internal/global"),
		repository.WithZapVar("global.Logger"),
		repository.WithZapVarPkg("demo/internal/global"),
	)
	err := r.Generate(
		models.User{},
		models.Admin{},
	)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("success")
}

```