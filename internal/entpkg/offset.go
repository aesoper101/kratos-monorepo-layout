package entpkg

import "context"

func Offset(_ context.Context, pageNo, pageSize int) int {
	if pageNo < 0 {
		pageNo = 1
	}
	return (pageNo - 1) * pageSize
}
