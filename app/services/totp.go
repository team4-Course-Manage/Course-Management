package services

import (
	"Course-Management/app/models"
	"github.com/pquerna/otp/totp"
)

// 生成TOTP密钥并返回密钥和二维码URL
func GenerateTOTP(user *models.User) (string, string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Course-Management",
		AccountName: user.Email, 
	})
	if err != nil {
		return "", "", err
	}

	// 更新用户的TOTP密钥
	user.TOTPSecret = key.Secret()
	if err := db.Save(user).Error; err != nil {
		return "", "", err
	}

	return key.Secret(), key.URL(), nil
}

// 验证TOTP代码
func VerifyTOTP(user *models.User, code string) bool {
	return totp.Validate(code, user.TOTPSecret)
}

// 获取用户的TOTP密钥
func GetTOTPSecret(userID uint) (string, error) {
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return "", err
	}
	return user.TOTPSecret, nil
}
