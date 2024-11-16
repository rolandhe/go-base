package commons

type IdRequest struct {
	Id int64 `json:"id"`
}

type IdListReq struct {
	IdList []int64 `json:"idList" form:"idList" binding:"required,gt=0" remark:"id列表"`
}
