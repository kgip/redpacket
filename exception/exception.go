package exception

import "fmt"

func TryThrow(params ...interface{}) {
	for _, param := range params {
		if param == nil {
			continue
		}
		if e, ok := param.(error); ok {
			panic(e)
		}
	}
}

type Exception struct {
	code   int    `json:"code"`
	msg    string `json:"msg"`
	cause  string `json:"cause"`
	detail string `json:"detail"`
}

func (e *Exception) Error() string {
	msg := e.msg
	if e.cause != "" {
		msg = fmt.Sprintf("%s Cause:%s", msg, e.cause)
	}
	if e.detail != "" {
		msg = fmt.Sprintf("%s Detail:%s", msg, e.detail)
	}
	return msg
}

func (e *Exception) GetCode() int {
	return e.code
}

func (e *Exception) Code(code int) *Exception {
	e.code = code
	return e
}

func (e *Exception) Msg(msg string) *Exception {
	e.msg = msg
	return e
}

func (e *Exception) Cause(cause string) *Exception {
	e.cause = cause
	return e
}

func (e *Exception) Detail(detail string) *Exception {
	e.detail = detail
	return e
}

var (
	InternalException      = Exception{50000, "Server Internal Error!", "", ""}
	RequestParamsException = Exception{50001, "Request Params Error!", "", ""}
)
