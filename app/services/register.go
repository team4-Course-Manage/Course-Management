package services

import (
	"regexp"
)

func IsValidEmail(email string) bool {
	re := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(re).MatchString(email)
}

func VerifyEmailCode(email, code string) bool {
	// 邮箱验证码验证逻辑
	return true
}

func IsValidPhone(phone string) bool {
	re := `^\+?[1-9]\d{1,14}$`
	return regexp.MustCompile(re).MatchString(phone)
}

func VerifyPhoneCode(phone, code string) bool {
	// 手机验证码验证逻辑
	return true
}
