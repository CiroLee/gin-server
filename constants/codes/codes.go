package codes

type ResResponse struct {
	Code int
	Msg  string
	Data any
}

var ErrorCode = map[string]ResResponse{
	"missParams": {
		Code: -1000,
		Msg:  "missing params",
	},
	"invalidParams": {
		Code: -1001,
		Msg:  "invalid params",
	},
}
