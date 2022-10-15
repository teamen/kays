package v1

import (
	"encoding/json"
	"testing"
)

type ListOption struct {
	Page     int64 `json:"page,omitempty" form:"page"`
	PageSize int64 `json:"page_size,omitempty" form:"page_size"`
}

type ListOptionUsePtr struct {
	Page     *int64 `json:"page,omitempty" form:"page"`
	PageSize *int64 `json:"page_size,omitempty" form:"page_size"`
}

func TestLi(t *testing.T) {
	// page := (int64)(1)
	// pageSize := (int64)(30)
	page := (int64)(0)
	pageSize := (int64)(0)

	opt := &ListOptionUsePtr{
		Page:     &page,
		PageSize: &pageSize,
	}
	bytes, _ := json.Marshal(opt)
	t.Log(string(bytes))

	opt2 := &ListOption{
		Page:     page,
		PageSize: pageSize,
	}
	bytes, _ = json.Marshal(opt2)
	t.Log(string(bytes))
}
