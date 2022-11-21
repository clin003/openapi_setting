package main

import (
	"errors"
	"fmt"
	"log"

	"fyne.io/fyne/v2/storage"

	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// MainShow 主界面函数
func mainShow(w fyne.Window) {
	title := widget.NewLabel("openAPI配置生成程序")
	hello := widget.NewLabel("文件夹路径:")
	entry1 := widget.NewEntry()
	dia1 := widget.NewButton("打开", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}
			entry1.SetText(reader.URI().Path())
		}, w)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".yaml"}))
		fd.Show()
	})
	label2 := widget.NewLabel("切面方式:")
	text := widget.NewMultiLineEntry()
	text.Disable()
	//labelLast := widget.NewLabel("摩比天线技术（深圳）有限公司    ALL Right Reserved")
	labelLast := widget.NewLabel("LYHUILIN Team.    ALL Right Reserved")
	combox1 := widget.NewSelect([]string{"最大值切面", "固定倾角切面"}, func(s string) { fmt.Println("selected", s) })

	label4 := widget.NewLabel("结果文件夹:")
	entry2 := widget.NewEntry()
	dia2 := widget.NewButton("打开", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if list == nil {
				log.Println("Cancelled")
				return
			}
			//out := fmt.Sprintf(list.String())
			entry2.SetText(list.Path())
		}, w)
	})
	combox1.SetSelectedIndex(0)

	bt3 := widget.NewButton("生成脚本", func() {
		if (entry1.Text != "") && (entry2.Text != "") {
			text.SetText("")
			text.Refresh()
			txtInfo := fmt.Sprintf("%s %s %s", entry1.Text, entry2.Text, combox1.Selected, w)
			text.SetText("TXT脚本生成成功。请复制下面的路径信息：\n" + txtInfo)
			text.Refresh()
		} else {
			dialog.ShowError(errors.New("读取文件错误"), w)
		}
	})
	bt4 := widget.NewButton("汇总结果", func() {
		fmt.Println(entry2.Text)
		if entry2.Text != "" {

		} else {
			dialog.ShowError(errors.New("文件夹路径错误"), w)
		}
	})
	head := container.NewCenter(title)
	v1 := container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), hello, dia1, entry1)
	v2 := container.NewHBox(label2, combox1)
	v4 := container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), label4, dia2, entry2)
	v5 := container.NewHBox(bt3, bt4)
	v5Center := container.NewCenter(v5)
	ctnt := container.NewVBox(head, v1, v2, v4, v5Center, text, labelLast)
	w.SetContent(ctnt)
}
