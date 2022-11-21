package orther_info

import (
	"open_api_setting/app_conf"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var (
	tpwd_symbol_left            binding.String
	tpwd_symbol_right           binding.String
	content_keyword_replace_all binding.String
	content_keyword_filter      binding.String
	content_taolijin_model      binding.Int
	content_keyword_ex_filter   binding.String
	content_keyword_ex_model    binding.Int
)

func init() {
	tpwd_symbol_left, tpwd_symbol_right,
		content_keyword_replace_all, content_keyword_filter,
		content_taolijin_model, content_keyword_ex_filter,
		content_keyword_ex_model =
		binding.NewString(), binding.NewString(),
		binding.NewString(), binding.NewString(),
		binding.NewInt(), binding.NewString(),
		binding.NewInt()
}
func title() string {
	return "附加配置"
}

func Save() error {
	if v, err := tpwd_symbol_left.Get(); err != nil {
		return err
	} else {
		if len(v) > 0 {
			app_conf.AppConf.TpwdSymbolLeft = v
		}
	}

	if v, err := tpwd_symbol_right.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.TpwdSymbolRight = v
	}

	if v, err := content_keyword_replace_all.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.ContentKeywordReplaceAll = v
	}

	if v, err := content_keyword_filter.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.ContentKeywordFilter = v
	}
	if v, err := content_taolijin_model.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.ContentTaolijinModel = v
	}
	if v, err := content_keyword_ex_filter.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.ContentKeywordExFilter = v
	}
	if v, err := content_keyword_ex_model.Get(); err != nil {
		return err
	} else {
		app_conf.AppConf.ContentKeywordExModel = v
	}

	return nil
}

// MainShow 主界面函数
func Show(parent fyne.Window) fyne.CanvasObject {
	title := widget.NewLabel(title())
	// label1 := widget.NewLabel("淘口令自定义符号左侧部分:")
	input1 := widget.NewEntryWithData(tpwd_symbol_left)
	input1.SetPlaceHolder("淘口令自定义符号左侧部分 (可选配置)")

	// label2 := widget.NewLabel("淘口令自定义符号右侧部分:")
	input2 := widget.NewEntryWithData(tpwd_symbol_right)
	input2.SetPlaceHolder("淘口令自定义符号右侧部分 (可选配置)")

	label3 := widget.NewLabel("文案级关键词替换:")
	input3 := widget.NewEntryWithData(content_keyword_replace_all)
	input3.SetPlaceHolder("格式 key>>key2|key3 (可选配置)")

	label4 := widget.NewLabel("文案级关键词过滤:")
	input4 := widget.NewEntryWithData(content_keyword_filter)
	input4.SetPlaceHolder("格式 key|key2 (可选配置)")

	label5 := widget.NewLabel("文案中淘礼金口令处理模式:")
	// input5 := widget.NewEntryWithData(content_taolijin_model)
	// input5.SetPlaceHolder("0、转链，1、屏蔽，2、原链(可选配置)")
	input5 := widget.NewSelect([]string{"0、转链", "1、屏蔽", "2、原链"}, func(value string) {
		switch {
		// case strings.HasPrefix(value, "0"):
		case strings.HasPrefix(value, "1"):
			content_taolijin_model.Set(1)
		case strings.HasPrefix(value, "2"):
			content_taolijin_model.Set(2)
		default:
			content_taolijin_model.Set(0)
		}
	})
	input5.SetSelectedIndex(0)

	label6 := widget.NewLabel("文案中特定关键词列表:")
	input6 := widget.NewEntryWithData(content_keyword_ex_filter)
	input6.SetPlaceHolder("格式 key|key2 (可选配置)")

	label7 := widget.NewLabel("文案中特定关键词处理模式:")
	// input7 := widget.NewEntryWithData(content_keyword_ex_model)
	// input7.SetPlaceHolder("0、转链，1、屏蔽，2、原链 (可选配置)")
	input7 := widget.NewSelect([]string{"0、转链", "1、屏蔽", "2、原链"}, func(value string) {
		switch {
		// case strings.HasPrefix(value, "0"):
		case strings.HasPrefix(value, "1"):
			content_keyword_ex_model.Set(1)
		case strings.HasPrefix(value, "2"):
			content_keyword_ex_model.Set(2)
		default:
			content_keyword_ex_model.Set(0)
		}
	})
	input7.SetSelectedIndex(0)

	head := container.NewCenter(title)
	// content := container.NewVBox(label, input)
	// content1 := container.NewGridWithColumns(label1, input1)
	// content1 := container.NewBorder(label1, input1, label2, input2)
	// content2 := container.NewVBox(label2, input2)
	content1 := container.NewVBox(input1, input2)
	content3 := container.NewVBox(label3, input3)
	content4 := container.NewVBox(label4, input4)
	content5 := container.NewHBox(label5, input5)
	content6 := container.NewVBox(label6, input6)
	content7 := container.NewHBox(label7, input7)

	content := container.NewVBox(head, content1,
		// content2,
		content3, content4, content5, content6, content7)
	return content
}
