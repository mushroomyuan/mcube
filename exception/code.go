package exception

import "net/http"

const (
	// TokenExpired token过期
	TokenExpired = 1000

	// 1xx - 5xx copy from http status code
	Unauthorized        = http.StatusUnauthorized
	BadRequest          = http.StatusBadRequest
	InternalServerError = http.StatusInternalServerError
	Forbidden           = http.StatusForbidden
	NotFound            = http.StatusNotFound

	UnKnownException = 9999
)

var (
	reasonMap = map[int]string{
		Unauthorized:        "认证失败",
		NotFound:            "资源为找到",
		BadRequest:          "请求不合法",
		InternalServerError: "系统内部错误",
		Forbidden:           "访问未授权",
		UnKnownException:    "未知异常",
		TokenExpired:        "访问过期, 请重新访问",
	}
)

func codeReason(code int) string {
	v, ok := reasonMap[code]
	if !ok {
		v = reasonMap[UnKnownException]
	}

	return v
}
