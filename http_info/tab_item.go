package http_info

import (
	"errors"
	"fmt"
	"open_api_setting/app_conf"

	"fyne.io/fyne/v2/dialog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var (
	addr           binding.String
	autotls_enable binding.Bool
	autotls_domain binding.String
	hlty_token     binding.String
	hlty_uuid      binding.String
	FmtText        string
)

func init() {
	addr, autotls_enable, autotls_domain, hlty_token, hlty_uuid =
		binding.NewString(), binding.NewBool(), binding.NewString(), binding.NewString(), binding.NewString()

}
func title() string {
	return "openAPI服务相关配置"
}

func Save() error {
	if v, err := addr.Get(); err != nil {
		return err
	} else {
		if len(v) > 0 {
			app_conf.AppConf.Addr = fmt.Sprintf(":%s", v)
		} else {
			app_conf.AppConf.Addr = fmt.Sprintf(":%s", "8080")
		}
	}

	if v, err := autotls_enable.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.AutotlsEnable = v
	}

	if v, err := autotls_domain.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.AutotlsDomain = v
	}
	if len(app_conf.AppConf.Addr) > 0 {
		if app_conf.AppConf.Addr != ":80" {
			app_conf.AppConf.Url = fmt.Sprintf("http://127.0.0.1%s", app_conf.AppConf.Addr)
		} else {
			app_conf.AppConf.Url = "http://127.0.0.1"
		}
	}

	app_conf.AppConf.MaxPingCount = 10
	app_conf.AppConf.TlsKeyFile = "./cert/apiclient.key"
	app_conf.AppConf.TlsPemFile = "./cert/apiclient.pem"

	app_conf.AppConf.Log.Writers = "file,stdout"
	app_conf.AppConf.Log.LoggerLevel = "INFO"
	app_conf.AppConf.Log.LoggerFile = "log/openapi.log"
	app_conf.AppConf.Log.LogFormatText = false
	app_conf.AppConf.Log.RollingPolicy = "size"
	app_conf.AppConf.Log.LogRotateDate = 1
	app_conf.AppConf.Log.LogRotateSize = 1
	app_conf.AppConf.Log.LogBackupCount = 7

	app_conf.AppConf.Db.Dbtype = "sqlite"
	app_conf.AppConf.Db.Name = "hltyapi"
	app_conf.AppConf.Db.Addr = "127.0.0.1:3306"
	app_conf.AppConf.Db.Host = "127.0.0.1"
	app_conf.AppConf.Db.Port = 3306
	app_conf.AppConf.Db.Username = ""
	app_conf.AppConf.Db.Password = ""

	app_conf.AppConf.SyncEnable = false

	app_conf.AppConf.UploadFile = true
	app_conf.AppConf.WsserverEnable = true

	// app_conf.AppConf.LiveEnbale = true

	app_conf.AppConf.CheckToken = false

	if v, err := hlty_token.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.HltyToken = v
	}

	if v, err := hlty_uuid.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.HltyUuid = v
	}

	app_conf.AppConf.TbkprivilegeEnable = true
	app_conf.AppConf.TbkprivilegeDisable = false

	FmtText = ""
	if app_conf.AppConf.AutotlsEnable {
		openAPIUrl := fmt.Sprintf("https://%s", app_conf.AppConf.AutotlsDomain)
		FmtText = fmt.Sprintf("token：%s \nopenAPI地址: %s", app_conf.AppConf.HltyToken, openAPIUrl)
	}
	if len(app_conf.AppConf.AutotlsDomain) <= 0 {
		openAPIUrl := app_conf.AppConf.Url
		FmtText = fmt.Sprintf("token：%s \nopenAPI地址: %s", app_conf.AppConf.HltyToken, openAPIUrl)
	}

	return nil
}

// MainShow 主界面函数
func Show(parent fyne.Window) fyne.CanvasObject {
	title := widget.NewLabel(title())
	labelAddr := widget.NewLabel("绑定端口:")
	inputAddr := widget.NewEntryWithData(addr)
	inputAddr.SetPlaceHolder("请输入绑定端口: 8080")

	labelAutotlsDomain := widget.NewLabel("TLS域名:")
	inputAutotlsDomain := widget.NewEntryWithData(autotls_domain)
	inputAutotlsDomain.SetPlaceHolder("请输入TLS域名： api.lyhuilin.com")
	inputAutotlsDomain.Disable()
	inputAutotlsEnable := widget.NewCheck("TLS", func(value bool) {
		autotls_enable.Set(value)
		if value {
			inputAutotlsDomain.Enable()
		} else {
			inputAutotlsDomain.Disable()
		}
	})
	// inputAutotlsEnable := widget.NewCheckWithData("TLS", inputAutotlsEnable)

	labelHltyToken := widget.NewLabel("验证Token:")
	inputHltyToken := widget.NewEntryWithData(hlty_token)
	inputHltyToken.SetPlaceHolder("请输入验证Token: 可选配置")

	labelHltyUuid := widget.NewLabel("授权UUID:")
	inputHltyUuid := widget.NewEntryWithData(hlty_uuid)
	inputHltyUuid.SetPlaceHolder("请输入授权UUID: 例子 xx-xx-xx-xx-xx")

	bt1 := widget.NewButton("淘宝授权", func() {
		// log.Println(inputHltyUuid.Text)
		if len(inputHltyUuid.Text) > 0 {
			uuidStr := inputHltyUuid.Text
			if err := lanchToTaobaoOauth(uuidStr); err != nil {
				dialog.ShowError(err, parent)
			}
		} else {
			dialog.ShowError(errors.New("授权UUID,为空"), parent)
		}
	})
	bt2 := widget.NewButton("唯品会授权", func() {
		// fmt.Println(inputHltyUuid.Text)
		if len(inputHltyUuid.Text) > 0 {
			uuidStr := inputHltyUuid.Text
			if err := lanchToWeipinhuiOauth(uuidStr); err != nil {
				dialog.ShowError(err, parent)
			}
		} else {
			dialog.ShowError(errors.New("授权UUID,为空"), parent)
		}
	})

	head := container.NewCenter(title)
	content1 := container.NewVBox(labelAddr, inputAddr)
	v1 := container.NewHBox(labelAutotlsDomain, inputAutotlsEnable)
	content2 := container.NewVBox(v1, inputAutotlsDomain)
	content3 := container.NewVBox(labelHltyToken, inputHltyToken)
	content4 := container.NewVBox(labelHltyUuid, inputHltyUuid)

	v5 := container.NewHBox(bt1, bt2)
	v5Center := container.NewCenter(v5)

	content := container.NewVBox(head, content1, content2, content3, content4, v5Center)
	return content
}
func UpdateUUID(uuidStr string) {
	hlty_uuid.Set(uuidStr)
}
