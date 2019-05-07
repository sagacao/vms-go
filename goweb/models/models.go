package models

type LogStats struct {
	Channel   string `json:"channel" form:"channel" binding:"required"`
	Game      string `json:"game" form:"game" binding:"required"`
	Newly     string `json:"newly" form:"newly" binding:"required"`
	TowPr     string `json:"tow_pr" form:"tow_pr" binding:"required"`
	ThreePr   string `json:"three_pr" form:"three_pr" binding:"required"`
	SevenPr   string `json:"seven_pr" form:"seven_pr" binding:"required"`
	Retention string `json:"retention" form:"retention" binding:"required"`
	LogDate   string `json:"logdate" form:"logdate" binding:"required"`
}
