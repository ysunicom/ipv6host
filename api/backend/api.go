package backend

type CommonPaginationReq struct {
	Page int `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	Size int `json:"size" in:"query" d:"10" v:"min:1|max:50#分页数量错误" dc:"分页数量，最小1，最大50"`
}

type CommonPaginationRes struct {
	Total int `dc:"总数"`
}