package util

type OpResult struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

func NewOpResult(code int64, msg string) OpResult {
	return OpResult{
		Code: code,
		Msg:  msg,
	}
}

func NewSucOpResult() OpResult {
	return OpResult{
		Code: Success,
		Msg:  "",
	}
}

func (op *OpResult) Success() bool {
	return op.Code == Success
}

const (
	Success = 0

	ErrBizIllegalParam = 1001
	ErrNoSupportAction = 1002

	ErrSystemInternal = 3000
	ErrDBConnect      = 3001
	ErrDBRead         = 3002
	ErrDBWrite        = 3003
	ErrDBUpdate       = 3004

	ErrRedis          = 3011
)
