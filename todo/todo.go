package todo

type List struct {
	Todo string `json:"todo"`
	Done bool   `json:"done"`
	Id   int    `json:"id"`
}
