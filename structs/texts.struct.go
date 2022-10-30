package structs

type TextsReq struct {
	Type     string `form:"type" binding:"oneof=word sentence paragraph name string"`
	Min      int    `form:"min" binding:"omitempty,gte=1"`
	Max      int    `form:"max" binding:"omitempty,lte=100,gtfield=Min"`
	Language string `form:"lang" binding:"omitempty,oneof=cn en"`
	Len      int    `form:"len" binding:"omitempty,gte=0"`
	Upper    bool   `form:"upper" binding:"omitempty"`
	Num      int    `form:"num" binding:"omitempty,gte=1"`
}
