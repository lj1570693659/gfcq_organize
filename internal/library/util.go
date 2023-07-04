package library

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func GetListWithPage(query *gdb.Model, page, size int32) (*gdb.Model, int32, error) {
	totalSize, err := query.Count()
	if err != nil {
		return query, gconv.Int32(totalSize), err
	}

	query = query.Limit(gconv.Int((page-1)*size), gconv.Int(size))
	return query, gconv.Int32(totalSize), nil
}

func FindIntSlice(slice []int32, val int32) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func DeleteIntSlice(a []int32) []int32 {
	ret := make([]int32, 0, len(a))
	for _, val := range a {
		if !g.IsEmpty(val) {
			ret = append(ret, val)
		}
	}
	return ret
}
