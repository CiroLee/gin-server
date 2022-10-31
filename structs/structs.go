package structs

type NumberReq struct {
	Type    string `form:"type" binding:"oneof=int float"`
	Min     int    `form:"min" binding:"required,gte=-999999"`
	Max     int    `form:"max" binding:"required,lte=999999,gtfield=Min"`
	Demical int    `form:"demical" binding:"min=0"`
	Num     int    `form:"num" binding:"omitempty,gte=1"`
}
type TextsReq struct {
	Type     string `form:"type" binding:"oneof=word sentence paragraph name string"`
	Min      int    `form:"min" binding:"omitempty,gte=1"`
	Max      int    `form:"max" binding:"omitempty,lte=100,gtfield=Min"`
	Language string `form:"lang" binding:"omitempty,oneof=cn en"`
	Len      int    `form:"len" binding:"omitempty,gte=0"`
	Upper    bool   `form:"upper" binding:"omitempty"`
	Num      int    `form:"num" binding:"omitempty,gte=1"`
}

type DateReq struct {
	Type string `form:"type" binding:"oneof=date timestamp weekday"`
	From int    `form:"from" binding:"omitempty,gte=0"`
	To   int    `form:"to" binding:"omitempty,gtefield=From"`
	Num  int    `form:"num" binding:"omitempty,gte=1"`
}
