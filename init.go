package main

import (
	"log"
	"os"
	"strings"

	"github.com/flopp/go-findfont"
)

// Fyne默认是不支持中文的。因此需要引入中文字体
// 设置环境变量
func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		log.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "Alibaba-PuHuiTi-Medium.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
		if strings.Contains(path, "Arial Unicode.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
		if strings.Contains(path, "JetBrains Mone.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
		if strings.Contains(path, "STHeiti Light.ttc") {
			os.Setenv("FYNE_FONT", path)
			break
		}
		if strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
		if strings.Contains(path, "simhei.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
		if strings.Contains(path, "Hack") {
			os.Setenv("FYNE_FONT", path)
			break
		}
		if strings.Contains(path, "CJK") && strings.HasSuffix(path, ".ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
		if strings.Contains(path, "Arial") {
			os.Setenv("FYNE_FONT", path)
			break
		}
		if strings.Contains(path, "Unicode") {
			os.Setenv("FYNE_FONT", path)
			break
		}
		if strings.Contains(path, "TerminusTTF") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
	log.Println("=============", os.Getenv("FYNE_FONT"))
}

// 取消环境变量
func unSetenv() {
	os.Unsetenv("FYNE_FONT")
}
