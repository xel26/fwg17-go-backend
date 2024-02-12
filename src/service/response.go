package service

type PageInfo struct {
	CurrentPage int `json:"currentPage"`
	TotalPage   int `json:"totalPage"`
	NextPage    int `json:"nextPage"`
	PrevPage    int `json:"prevPage"`
	Limit       int `json:"limit"`
	TotalData   int `json:"totalData"`
}

type ResponseList struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	PageInfo PageInfo    `json:"pageInfo"`
	Results  interface{} `json:"results"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

type ResponseOnly struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}