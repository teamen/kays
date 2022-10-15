package gormutil

const DefaultPage = 1
const DefaultPageSize = 30

type LimitAndOffset struct {
	Offset int
	Limit  int
}

func ParseLimitAndOffset(page *int64, pageSize *int64) *LimitAndOffset {

	var offset, limit int = 0, DefaultPageSize

	if page != nil {
		offset = (int(*page) - 1) * int(*pageSize)
	}

	if pageSize != nil {
		limit = int(*pageSize)
	}

	return &LimitAndOffset{
		Offset: offset,
		Limit:  limit,
	}
}
