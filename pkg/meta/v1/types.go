package v1

type ListOptions struct {
	Page     *int64 `json:"page,omitempty" form:"page"`
	PageSize *int64 `json:"page_size,omitempty" form:"page_size"`
}

// ListMeta describes metadata that synthetic resources must have, including lists and
// various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
type ListMeta struct {
	TotalCount int64 `json:"total,omitempty"`
}
