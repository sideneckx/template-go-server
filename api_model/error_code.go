package api_model

type ErrorCode int

const (
	Success       ErrorCode = 0
	MissingParam  ErrorCode = 2
	RequestFail   ErrorCode = 3
	NeedAuthorize ErrorCode = 4
)
