package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"

	"yooplus_indication/lib/base64Captcha"
)

// 验证码接口

type CaptchaController struct {
	BaseController
}

// [options == 0  生成验证码]
// [options == 1  验证验证码]
func (this *CaptchaController) Post() {

	options, err := strconv.Atoi(this.Input().Get("options"))

	if err != nil {
		beego.Info(err.Error())
	}

	//生成验证码
	// [captcha_type == 0  数字验证码]
	// [captcha_type == 1  声音验证码]
	// [captcha_type == 2  公式验证码]
	if options == 0 {

		captcha_type, err := strconv.Atoi(this.Input().Get("captcha_type"))
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "验证码类型错误"}
			this.ServeJSON()
			return
		}
		key, base64string := CaptchaCreate(captcha_type)
		beego.Info(key)
		this.Data["json"] = map[string]interface{}{"status": 200, "data": base64string, "key": key, "message": "success"}
		this.ServeJSON()
		return

	}

	//验证验证码
	// [key --- 传给客户端的uid]
	// [value --- 用户输入的验证码]
	if options == 1 {

		key := this.Input().Get("key")
		value := this.Input().Get("value")
		beego.Info(key)
		beego.Info(value)
		flag := verfiyCaptcha(key, value)
		if flag {
			this.Data["json"] = map[string]interface{}{"status": 200, "message": "success"}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 400, "message": "验证码错误"}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"code": 400, "message": "options error"}
	this.ServeJSON()
	return
}

func (this *CaptchaController) Get() {
	this.TplName = "captcha_test.html"
}

//生成验证码
// [captcha_type == 0  数字验证码]
// [captcha_type == 1  声音验证码]
// [captcha_type == 2  公式验证码]
func CaptchaCreate(captcha_type int) (string, string) {
	var idKey string
	var base64string string
	if captcha_type == 0 {
		//config struct for digits
		//数字验证码配置
		var configD = base64Captcha.ConfigDigit{
			Height:     80,
			Width:      240,
			MaxSkew:    0.7,
			DotCount:   80,
			CaptchaLen: 5,
		}
		//创建数字验证码.
		//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
		idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
		//以base64编码
		base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)

		fmt.Println(idKeyD, base64stringD, "\n")
		// beego.Info(idKeyD)
		// beego.Info(base64stringD)
		idKey = idKeyD
		base64string = base64stringD
	}
	if captcha_type == 1 {
		//config struct for audio
		//声音验证码配置
		var configA = base64Captcha.ConfigAudio{
			CaptchaLen: 6,
			Language:   "zh",
		}
		//创建声音验证码
		//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
		idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
		//以base64编码
		base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
		fmt.Println(idKeyA, base64stringA, "\n")
		idKey = idKeyA
		base64string = base64stringA
	}
	if captcha_type == 2 {
		//config struct for Character
		//字符,公式,验证码配置
		var configC = base64Captcha.ConfigCharacter{
			Height: 60,
			Width:  240,
			//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
			Mode:               base64Captcha.CaptchaModeNumber,
			ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
			ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
			IsShowHollowLine:   false,
			IsShowNoiseDot:     false,
			IsShowNoiseText:    false,
			IsShowSlimeLine:    false,
			IsShowSineLine:     false,
			CaptchaLen:         6,
		}

		//创建字符公式验证码.
		//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
		idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
		//以base64编码
		base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)

		fmt.Println(idKeyC, base64stringC, "\n")

		idKey = idKeyC
		base64string = base64stringC
	}
	return idKey, base64string
}

//验证验证码
func verfiyCaptcha(idkey, verifyValue string) bool {
	verifyResult := base64Captcha.VerifyCaptcha(idkey, verifyValue)
	if verifyResult {
		beego.Info("验证码验证成功...")
		return true
	} else {
		beego.Info("验证码验证失败...")
		return false
	}
}
