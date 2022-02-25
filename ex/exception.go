package ex

import "encoding/json"

func TryThrow(params ...interface{}) {
	if e, success := params[len(params)-1].(*Exception); success {
		for _, param := range params[:len(params)-1] {
			if param == nil {
				continue
			}
			if err, ok := param.(error); ok {
				e.Detail = err.Error()
				panic(e)
			}
		}
	} else {
		for _, param := range params {
			if param == nil {
				continue
			}
			if err, ok := param.(error); ok {
				panic(err)
			}
		}
	}
}

type Exception struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Cause  string `json:"cause"`
	Detail string `json:"detail"`
}

func (e *Exception) Error() string {
	return e.Msg
}

func (e *Exception) DetailInfo() string {
	if bytes, err := json.Marshal(e); err != nil {
		panic(err)
	} else {
		return string(bytes)
	}
}

var (
	InternalException      = &Exception{50000, "Server Internal Error!", "", ""}
	RequestParamsException = &Exception{50001, "Request Params Error!", "", ""}
)
