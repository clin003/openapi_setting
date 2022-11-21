package sdk_jingdong

import (
	"open_api_setting/app_conf"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var (
	appkey    binding.String
	appsecret binding.String
	adzoneid  binding.String
	siteid    binding.String
	pid       binding.String
)

func init() {
	appkey, appsecret, adzoneid, siteid, pid =
		binding.NewString(), binding.NewString(), binding.NewString(), binding.NewString(), binding.NewString()
}
func title() string {
	return "京东联盟相关配置"
}

func Save() error {
	if v, err := appkey.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.Sdk.Jingdong.Appkey = v
	}

	if v, err := appsecret.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.Sdk.Jingdong.Appsecret = v
	}

	if v, err := adzoneid.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.Sdk.Jingdong.Adzoneid = v
	}
	if v, err := siteid.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.Sdk.Jingdong.Siteid = v
	}

	if v, err := pid.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.Sdk.Jingdong.UnionId = v
		app_conf.AppConf.App.Jd.UnionId = v
	}
	return nil
}

// MainShow 主界面函数
func Show(parent fyne.Window) fyne.CanvasObject {
	title := widget.NewLabel(title())
	labelAppkey := widget.NewLabel("appkey:")
	inputAppkey := widget.NewEntryWithData(appkey)
	inputAppkey.SetPlaceHolder("请输入appkey")

	labelAppsert := widget.NewLabel("appsert:")
	inputAppseret := widget.NewEntryWithData(appsecret)
	inputAppseret.SetPlaceHolder("请输入appsert")

	labelAdzoneid := widget.NewLabel("adzoneid:")
	inputAdzoneid := widget.NewEntryWithData(adzoneid)
	inputAdzoneid.SetPlaceHolder("请输入adzoneid: PID第三段数字")

	labelSiteid := widget.NewLabel("siteid:")
	inputSiteid := widget.NewEntryWithData(siteid)
	inputSiteid.SetPlaceHolder("请输入adzoneid: PID第二段数字")

	labelPid := widget.NewLabel("unionid:")
	inputPid := widget.NewEntryWithData(pid)
	inputPid.SetPlaceHolder("请输入unionid: 京东联盟Id")

	head := container.NewCenter(title)
	contentAppkey := container.NewVBox(labelAppkey, inputAppkey)
	contentAppseret := container.NewVBox(labelAppsert, inputAppseret)
	contentAdzoneid := container.NewVBox(labelAdzoneid, inputAdzoneid)
	contentSiteid := container.NewVBox(labelSiteid, inputSiteid)
	contentPid := container.NewVBox(labelPid, inputPid)
	content := container.NewVBox(
		head, contentAppkey, contentAppseret, contentAdzoneid, contentSiteid, contentPid,
	)
	return content
}
