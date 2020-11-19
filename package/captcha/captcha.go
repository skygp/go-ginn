package captcha

import (
	"github.com/afocus/captcha"
	"image/color"
)

func NewCaptcha(fontPath string, number, width, height int, disturbance captcha.DisturLevel) (img *captcha.Image, captchaRet string) {
	Cap := captcha.New()
	if err := Cap.SetFont(fontPath); err != nil {
		panic(err.Error())
	}
	Cap.SetSize(width, height)
	Cap.SetDisturbance(disturbance)
	Cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	Cap.SetBkgColor(color.RGBA{255, 125, 125, 255}, color.RGBA{125, 125, 255, 255}, color.RGBA{0, 153, 0, 255})
	img, captchaRet = Cap.Create(number, captcha.ALL)
	return
}
