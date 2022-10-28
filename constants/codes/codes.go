package codes

type ResResponse struct {
	Code int
	Msg  string
	Data any
}

var ErrorCode = map[string]ResResponse{
	"invalidParams": {
		Code: -1000,
		Msg:  "invalid params",
	},
}
