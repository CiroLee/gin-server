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
	Type     string `form:"type" binding:"oneof=date timestamp weekday month"`
	From     int    `form:"from" binding:"omitempty,gte=0"`
	To       int    `form:"to" binding:"omitempty,gtefield=From"`
	Abbr     bool   `form:"abbr" binding:"omitempty"`
	Language string `form:"lang" binding:"omitempty,oneof=cn en"`
	Num      int    `form:"num" binding:"omitempty,gte=1"`
}

type ImageReq struct {
	Type      string `form:"type" binding:"oneof=placeholder picsum"`
	Width     int    `form:"width" binding:"omitempty,gt=0"`
	Height    int    `form:"width" binding:"omitempty,gt=0"`
	BgColor   string `form:"bgcolor" binding:"omitempty"`
	FontColor string `form:"fontcolor" binding:"omitempty"`
	Text      string `form:"text" binding:"omitempty"`
	Grayscale bool   `form:"grayscale" binding:"omitempty"`
	Blur      int    `form:"blur" binding:"omitempty, gte=1,lte=10"`
	Num       int    `form:"num" binding:"omitempty,gte=1"`
}

type ColorReq struct {
	Type string `form:"type" binding:"oneof=hex rgb rgba hsl hsla"`
	Num  int    `form:"num" binding:"omitempty,gte=1"`
}
