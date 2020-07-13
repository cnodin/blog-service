package errcode

var (
	Success						=	NewError(0, "Success")
	ServerError					=	NewError(10000000, "Internal server error")
	InvalidParams				=	NewError(10000001,"The param is mistake")
	NotFound					=	NewError(10000002, "Not found")
	UnauthorizedAuthNotExist	=	NewError(10000003, "Unauthorized, can't find appkey and appsecret")
	UnauthorizedTokenError		=	NewError(10000004, "Unauthorized, token is mistake")
	UnauthorizedTokenTimeout	=	NewError(10000005, "Unauthorized, token is timeout")
	UnauthorizedTokenGenerate	=	NewError(10000006, "Unauthorized, token is generated")
	TooManyRequests				=	NewError(10000007, "Too many requests")
)
