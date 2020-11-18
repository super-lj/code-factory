package domain

type Commit struct {
	Hash    string  `json:"hash"`
	Msg     string  `json:"msg"`
	Author  string  `json:"author"`
	RunNums []int32 `json:"run_nums"`
}
