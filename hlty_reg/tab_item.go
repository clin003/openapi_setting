package hlty_reg

import (
	"errors"
	"fmt"
	"open_api_setting/app_conf"
	"open_api_setting/http_info"
	"open_api_setting/utils"

	"fyne.io/fyne/v2/dialog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var (
	username binding.String
	password binding.String
	email    binding.String
)

func init() {
	username, password, email =
		binding.NewString(), binding.NewString(), binding.NewString()
}
func title() string {
	return "授权注册"
}

func Save() error {
	if v, err := username.Get(); err != nil {
		return err
	} else {
		app_conf.UserInfo.Username = v
	}

	if v, err := password.Get(); err != nil {
		return err
	} else {
		app_conf.UserInfo.Password = v
	}

	if v, err := email.Get(); err != nil {
		return err
	} else {
		app_conf.UserInfo.Email = v
	}

	return nil
}

// MainShow 主界面函数
func Show(parent fyne.Window) fyne.CanvasObject {
	title := widget.NewLabel(title())
	label1 := widget.NewLabel("账号:")
	input1 := widget.NewEntryWithData(username)
	input1.SetPlaceHolder("请输入账号（手机号）")

	label2 := widget.NewLabel("密码:")
	input2 := widget.NewPasswordEntry()
	input2.Bind(password)
	input2.SetPlaceHolder("请输入密码")

	label3 := widget.NewLabel("邮箱:")
	input3 := widget.NewEntryWithData(email)
	input3.SetPlaceHolder("请输入邮箱")

	text := widget.NewMultiLineEntry()
	text.Disable()
	bt1 := widget.NewButton("注册信息", func() {
		if (input1.Text != "") && (input2.Text != "") {
			Save()
			app_conf.UserInfo.Action = "adduserinfo"
			if respText, err := utils.RegUserInfoEx(app_conf.UserInfo); err != nil {
				dialog.ShowError(err, parent)
				fmt.Println(err)
			} else {
				text.SetText("授权UUID:" + respText)
				text.Refresh()
				if len(respText) > 0 {
					app_conf.AppConf.HltyUuid = respText
					http_info.UpdateUUID(respText)
				}
			}
		} else {
			dialog.ShowError(errors.New("请补全注册信息"), parent)
		}
	})
	bt2 := widget.NewButton("登录信息", func() {
		if (input1.Text != "") && (input2.Text != "") {
			app_conf.UserInfo.Action = "userlogin"
			if respText, err := utils.LoginUserInfoEx(app_conf.UserInfo); err != nil {
				dialog.ShowError(err, parent)
				// fmt.Println(err)
			} else {
				text.SetText("授权UUID:" + respText)
				text.Refresh()
				if len(respText) > 0 {
					app_conf.AppConf.HltyUuid = respText
					http_info.UpdateUUID(respText)
				}
			}

		} else {
			dialog.ShowError(errors.New("请补全登录信息"), parent)
		}
	})

	head := container.NewCenter(title)
	content1 := container.NewVBox(label1, input1)
	content2 := container.NewVBox(label2, input2)
	content3 := container.NewVBox(label3, input3)
	v3 := container.NewHBox(bt1, bt2)
	content4 := container.NewCenter(v3)
	content := container.NewVBox(head, content1, content2, content3, content4, text)
	return content
}
