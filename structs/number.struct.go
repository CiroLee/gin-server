package structs

type NumberReq struct {
	Type    string `form:"type" binding:"oneof=int float"`
	Min     int    `form:"min" binding:"required,gte=-999999"`
	Max     int    `form:"max" binding:"required,lte=999999"`
	Demical int    `form:"demical" binding:"min=0"`
}
