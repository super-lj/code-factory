package domain

type Commit struct {
	Id             int64  `json:"id"`
	Msg            string `json:"msg"`
	Author         string `json:"author"`
	LastUpdateTime int64  `json:"last_update_time"`
}
