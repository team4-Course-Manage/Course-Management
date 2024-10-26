package dto

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Captcha  string `form:"captcha" binding:"required,min=4,max=6"`
}

type LoginResp struct {
	Token string `json:"token"`
}
