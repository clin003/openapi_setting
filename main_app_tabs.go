package main

import (
	"errors"
	"log"
	"open_api_setting/http_info"
	"open_api_setting/sdk_jingdong"
	"open_api_setting/sdk_pinduoduo"
	"open_api_setting/sdk_taobao"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func appTabs(parent fyne.Window) {
	// title := widget.NewLabel("openAPI配置生成程序")
	tabs := container.NewAppTabs(
		// container.NewTabItem("Tab1", widget.NewLabel("tab1")),
		container.NewTabItem("openAPI", http_info.Show(parent)),
		container.NewTabItem("淘宝", sdk_taobao.Show(parent)),
		container.NewTabItem("京东", sdk_jingdong.Show(parent)),
		container.NewTabItem("拼多多", sdk_pinduoduo.Show(parent)),
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	bt := widget.NewButton("保存配置", func() {

		if err := save(); err != nil {
			showError(parent, err, "错误")
		} else {
			dialog.ShowConfirm("保存", "确认保存？", func(isYes bool) {
				log.Println(isYes)
				if isYes {
					saveToYaml()
				}
			}, parent)
		}
	})
	// labelLast := widget.NewLabel("LYHUILIN Team.    ALL Right Reserved")

	// head := container.NewCenter(title)
	contentTab := container.New(layout.NewMaxLayout(), tabs)
	contentButton := container.New(layout.NewVBoxLayout(), bt)
	parent.SetContent(
		container.NewVBox(
			// head,
			contentTab, contentButton,
			// labelLast,
		),
	)
}

func showError(parent fyne.Window, err error, errString string) {
	dialog.ShowError(errors.New(errString), parent)
}
