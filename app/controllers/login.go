package controllers

import (
	"Course-Management/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginCtl struct{}

// 系统登录
func (c *loginCtl) Login(ctx *gin.Context) {
	// 从请求中获取用户名和密码
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	totpCode := ctx.PostForm("totp_code")

	// 验证用户名和密码是否为空
	if username == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username and password cannot be empty"})
		return
	}

	// 调用服务层的登录处理函数
	user, err := services.LoginHandler(username, password)
	if err != nil {
		// 登录失败，返回错误信息
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	// 验证 TOTP 代码
	if !services.VerifyTOTP(user, totpCode) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid TOTP code"})
		return
	}
	// 登录成功，生成token
	token := "1"

	// 返回成功的响应和token
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user":    user,
	})
}

var Login = loginCtl{}

// 验证码,依赖包版本要降
// func (c *loginCtl) Captcha(ctx *gin.Context) {
// 	// 验证码参数配置：字符,公式,验证码配置
// 	var configC = base64Captcha.ConfigCharacter{
// 		Height: 60,
// 		Width:  240,
// 		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
// 		Mode:               base64Captcha.CaptchaModeAlphabet,
// 		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
// 		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
// 		IsShowHollowLine:   false,
// 		IsShowNoiseDot:     false,
// 		IsShowNoiseText:    false,
// 		IsShowSlimeLine:    false,
// 		IsShowSineLine:     false,
// 		CaptchaLen:         6,
// 	}
// 	///create a characters captcha.
// 	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
// 	//以base64编码
// 	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)

// 	// 返回结果集
// 	ctx.JSON(http.StatusOK, common.CaptchaRes{
// 		Code:  0,
// 		IdKey: idKeyC,
// 		Data:  base64stringC,
// 		Msg:   "操作成功",
// 	})
// }
