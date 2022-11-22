package main

import (
	_ "open_api_setting/fonts"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {

	// 新建App
	myApp := app.New()
	myApp.SetIcon(resourceFavPng)
	// 新建窗口
	window := myApp.NewWindow("openAPI配置程序")
	// 主界面框架布局
	// mainShow(window)
	appTabs(window)
	// 尺寸
	window.Resize(fyne.Size{Width: 500, Height: 80})
	// 居中显示
	window.CenterOnScreen()
	// 循环运行
	window.ShowAndRun()
	// unSetenv()
}
