package models

type LogStats struct {
	Channel   string `json:"channel" form:"channel" binding:"required"`
	Game      string `json:"game" form:"game" binding:"required"`
	LogDate   string `json:"logdate" form:"logdate" binding:"required"`
	Newly     string `json:"newly" form:"newly" `
	TowPr     string `json:"tow_pr" form:"tow_pr" `
	ThreePr   string `json:"three_pr" form:"three_pr"`
	SevenPr   string `json:"seven_pr" form:"seven_pr" `
	Retention string `json:"retention" form:"retention" `
}

type GiftCode struct {
	Game   string `json:"game" form:"game" binding:"required"`
	Code   string `json:"code" form:"code" binding:"required"`
	Atype  string `json:"atype" form:"atype" binding:"required"`
	Acount string `json:"acount" form:"acount" binding:"required"`
}
