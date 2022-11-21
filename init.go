package main

import (
	"log"
	"open_api_setting/fonts"
	"os"
	"path/filepath"
	"strings"

	"gitee.com/lyhuilin/util"

	"github.com/goki/freetype/truetype"

	"github.com/flopp/go-findfont"
)

// Fyne默认是不支持中文的。因此需要引入中文字体
// 设置环境变量
func init() {
	if err := initEx(); err != nil {
		log.Println("initEx", err)
	} else {
		return
	}

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

	}
	log.Println("=============", os.Getenv("FYNE_FONT"))
}

func initEx() error {
	fileName := "Alibaba-PuHuiTi-Medium.ttf"
	fileBody, err := fonts.FontList.ReadFile(fileName)
	if err != nil {
		log.Println(err)
		return err
	}
	filePath := filepath.Join(".", "FontLibrary")
	if !util.IsExist(filePath) {
		if err := util.MkDir(filePath); err != nil {
			log.Println(err)
			return err
		}
	}
	filePath = filepath.Join(".", "FontLibrary", fileName)
	if ok, err := util.IsFileExist(filePath); err != nil || !ok {
		if err := os.WriteFile(filePath, fileBody, 0644); err != nil {
			log.Println(err)
			return err
		}
	}
	return initEx2(filePath)
}
func initEx2(filePath string) error {
	fontPath, err := findfont.Find(filePath)
	if err != nil {
		log.Println(err)
		return err
	}

	// // load the font with the freetype library
	// // 原作者使用的ioutil.ReadFile已经弃用
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = truetype.Parse(fontData)
	if err != nil {
		log.Println(err)
		return err
	}
	return os.Setenv("FYNE_FONT", fontPath)
}

// 取消环境变量
func unSetenv() {
	os.Unsetenv("FYNE_FONT")
}
