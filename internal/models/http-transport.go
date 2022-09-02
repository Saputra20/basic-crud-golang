package models

type HTTPResponse struct {
	Success bool        `json:"sucess"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

type HTTPResponsePaginate struct {
	Page        int `json:"page"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
}

type HTTPAPIParamReq struct {
	ID int `uri:"id" binding:"required"`
}

type HTTPAPIBodyReq struct {
	Name string `json:"name" binding:"required"`
}

type HTTPAPIQueryReq struct {
	Limit int    `form:"limit" binding:"required"`
	Page  int    `form:"page" binding:"required"`
	Sort  string `form:"sort" binding:"required"`
}
