package ctl

type Response struct {
	Status int
	Data   interface{}
	Msg    string
	Error  string
}

func RespSuccess(code ...int) *Response {
	status := 200
	if code != nil {
		status = code[0]
	}
	r := &Response{
		Status: status,
		Data:   "操作成功",
	}
	return r
}
