package commons

import "encoding/json"

type PageParam struct {
	PageNo   int    `json:"pageNo" form:"pageNo" binding:"required,gt=0"`
	PageSize int    `json:"pageSize" form:"pageSize" binding:"required,gt=0"`
	Sort     string `json:"sort" form:"sort"`
	Asc      bool   `json:"asc" form:"asc"`
}

func (pp *PageParam) GetFirstResult() int {
	return (pp.PageNo - 1) * pp.PageSize
}

type PageList[T any] struct {
	PageNo     int            `json:"pageNo"`
	PageSize   int            `json:"pageSize"`
	TotalCount int            `json:"totalCount"`
	TotalPages int            `json:"totalPages"`
	List       []*T           `json:"list"`
	Extra      map[string]any `json:"extra,omitempty"`
}

func (p *PageList[T]) String() string {
	j, err := json.Marshal(p)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}

func (p *PageList[T]) SetExtraKeyValue(key string, val any) {
	if p.Extra == nil {
		p.Extra = make(map[string]any)
	}

	p.Extra[key] = val
}

func (p *PageList[T]) GetExtraValue(key string) any {
	if p.Extra == nil {
		return nil
	}

	return p.Extra[key]
}

func BuildPageList[T any](pp *PageParam, totalCount int, list []*T) *PageList[T] {
	return &PageList[T]{
		PageNo:     pp.PageNo,
		PageSize:   pp.PageSize,
		TotalCount: totalCount,
		TotalPages: calcTotalPages(totalCount, pp.PageSize),
		List:       list,
	}
}

func ListOf[T any](pageNo int, pageSize int, totalCount int, list []*T) *PageList[T] {
	return &PageList[T]{
		PageNo:     pageNo,
		PageSize:   pageSize,
		TotalCount: totalCount,
		TotalPages: calcTotalPages(totalCount, pageSize),
		List:       list,
	}
}

func EmptyPageList[T any](pageNo int, pageSize int) *PageList[T] {
	return ListOf[T](pageNo, pageSize, 0, []*T{})
}

func calcTotalPages(totalCount int, pageSize int) int {
	if totalCount == 0 {
		return 0
	}

	totalPages := totalCount / pageSize

	if totalCount%pageSize > 0 {
		totalPages++
	}

	return totalPages
}
