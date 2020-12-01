package captcha

import (
	"ginn/package/snowflake"
	"github.com/afocus/captcha"
	"image/color"

	"github.com/mojocn/base64Captcha"
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

var store = base64Captcha.DefaultMemStore

//configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

func DriverStringFunc() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = snowflake.GenIDToString()
	e.DriverString = base64Captcha.NewDriverString(46, 140, 2, 2, 4, "234567890abcdefghjkmnpqrstuvwxyz", &color.RGBA{240, 240, 246, 246}, []string{"wqy-microhei.ttc"})
	driver := e.DriverString.ConvertFonts()
	cap := base64Captcha.NewCaptcha(driver, store)
	return cap.Generate()
}

func DriverDigitFunc() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = snowflake.GenIDToString()
	e.DriverDigit = base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	driver := e.DriverDigit
	cap := base64Captcha.NewCaptcha(driver, store)
	return cap.Generate()
}
