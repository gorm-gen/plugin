package list_test

import (
	"encoding/json"
	"testing"

	"github.com/gorm-gen/plugin/sharding/query/list"
)

func TestAnalysis(t *testing.T) {
	var l []*list.St
	l = append(l, &list.St{
		ShardingValue: "202501",
		Total:         7,
	})
	l = append(l, &list.St{
		ShardingValue: "202502",
		Total:         10,
	})
	l = append(l, &list.St{
		ShardingValue: "202503",
		Total:         20,
	})
	l = append(l, &list.St{
		ShardingValue: "202504",
		Total:         5,
	})
	l = append(l, &list.St{
		ShardingValue: "202505",
		Total:         6,
	})
	l = append(l, &list.St{
		ShardingValue: "202506",
		Total:         30,
	})
	res := list.New(l, list.WithAsc(), list.WithPage(2), list.WithPageSize(10), list.WithOffset(3)).Analysis()
	jb, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(jb))
}
