/**
 * 验证码 服务层
 */
package services

import (
	"gin-api/common/utils"
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// 使用redis作为store
var sotre = utils.RedisStore{}

func CaptMake() (id, b64s string) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString
	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          6,
		Source:          "abcdefghijklmnopqrstuvwxyz0123456789",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 120,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, sotre)
	lid, lb64s, _ := captcha.Generate()
	return lid, lb64s
}

// 验证captcha是否正确
func CaptVerify(id string, capt string) bool {
	if sotre.Verify(id, capt, false) {
		return true
	} else {
		return false
	}
}
