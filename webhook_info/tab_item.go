package webhook_info

import (
	"open_api_setting/app_conf"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var (
	weixin_work_webhook_url      binding.String
	weixin_work_webhook_enable   binding.Bool
	dingtalk_webhook_url         binding.String
	dingtalk_webhook_secret      binding.String
	dingtalk_webhook_enable      binding.Bool
	dingtalk_webhook_rich_enable binding.Bool
)

func init() {
	weixin_work_webhook_url, weixin_work_webhook_enable,
		dingtalk_webhook_url, dingtalk_webhook_secret,
		dingtalk_webhook_enable,
		dingtalk_webhook_rich_enable =
		binding.NewString(), binding.NewBool(),
		binding.NewString(), binding.NewString(),
		binding.NewBool(), binding.NewBool()
}
func title() string {
	return "Webhook配置"
}

func Save() error {
	if v, err := weixin_work_webhook_url.Get(); err != nil {
		return err
	} else {
		if len(v) > 0 {
			app_conf.AppConf.WeixinWorkWebhookUrl = v
		}
	}

	if v, err := weixin_work_webhook_enable.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.WeixinWorkWebhookEnable = v
	}

	if v, err := dingtalk_webhook_url.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.DingtalkWebhookUrl = v
	}

	if v, err := dingtalk_webhook_secret.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.DingtalkWebhookSecret = v
	}
	if v, err := dingtalk_webhook_enable.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.DingtalkWebhookEnable = v
	}
	if v, err := dingtalk_webhook_rich_enable.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.DingtalkWebhookRichEnable = v
	}

	return nil
}

// MainShow 主界面函数
func Show(parent fyne.Window) fyne.CanvasObject {
	title := widget.NewLabel(title())
	label1 := widget.NewLabel("企业微信Webhook地址:")
	input1 := widget.NewEntryWithData(weixin_work_webhook_url)
	input1.SetPlaceHolder("企业微信Webhook地址 (可选配置)")

	// label2 := widget.NewLabel("企业微信webhook发布开启:")
	input2 := widget.NewCheckWithData("企业微信webhook发布开启", weixin_work_webhook_enable)

	label3 := widget.NewLabel("钉钉群bot Webhook地址:")
	input3 := widget.NewEntryWithData(dingtalk_webhook_url)
	input3.SetPlaceHolder("钉钉群bot Webhook地址 (可选配置)")

	label4 := widget.NewLabel(" 钉钉群bot Secret:")
	input4 := widget.NewEntryWithData(dingtalk_webhook_secret)
	input4.SetPlaceHolder("钉钉群bot Webhook地址 Secret (可选配置)")

	input5 := widget.NewCheckWithData("钉钉群bot webhook发布开启", dingtalk_webhook_enable)
	input6 := widget.NewCheckWithData("钉钉群bot webhook发布图文开启", dingtalk_webhook_rich_enable)

	head := container.NewCenter(title)
	// content := container.NewVBox(label, input)
	// content1 := container.NewGridWithColumns(label1, input1)
	// content1 := container.NewBorder(label1, input1, label2, input2)
	// content2 := container.NewVBox(label2, input2)
	// content1 := container.NewVBox(input1, input2)
	content1 := container.NewVBox(label1, input1)
	content2 := container.NewVBox(input2)
	content3 := container.NewVBox(label3, input3)
	content4 := container.NewVBox(label4, input4)
	content5 := container.NewHBox(input5, input6)

	content := container.NewVBox(head, content1,
		content2,
		content3, content4, content5)
	return content
}
