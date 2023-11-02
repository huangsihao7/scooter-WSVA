package xcode

var (
	OK                 = add(0, "OK")
	NoLogin            = add(101, "用户未登录")
	RequestErr         = add(400, "INVALID_ARGUMENT")
	Unauthorized       = add(401, "UNAUTHENTICATED")
	AccessDenied       = add(403, "PERMISSION_DENIED")
	NotFound           = add(404, "NOT_FOUND")
	MethodNotAllowed   = add(405, "METHOD_NOT_ALLOWED")
	Canceled           = add(498, "CANCELED")
	ServerErr          = add(500, "服务内部错误")
	ServiceUnavailable = add(503, "UNAVAILABLE")
	Deadline           = add(504, "服务响应超时")
	LimitExceed        = add(509, "RESOURCE_EXHAUSTED")
)
