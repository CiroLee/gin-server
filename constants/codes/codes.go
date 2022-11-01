package codes

type ResResponse struct {
	Code int
	Msg  string
	Data any
}

var ErrorCode = map[string]ResResponse{
	"RequestError": {
		Code: -1000,
		Msg:  "request error",
	},
}
