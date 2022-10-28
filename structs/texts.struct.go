package structs

type TextsReq struct {
	Type     string `form:"type" binding:"oneof=word sentence"`
	Min      int    `form:"min" binding:"omitempty,gte=1"`
	Max      int    `form:"max" binding:"omitempty,lte=100"`
	Language string `form:"lang" binding:"omitempty,oneof=cn en"`
	Len      int    `form:"len" binding:"gte=0"`
}
