package main

import (
	"io/ioutil"
	"log"
	"open_api_setting/app_conf"
	"open_api_setting/http_info"
	"open_api_setting/sdk_jingdong"
	"open_api_setting/sdk_pinduoduo"
	"open_api_setting/sdk_taobao"

	yaml "gopkg.in/yaml.v3"
)

func saveToYaml() {
	b, err := yaml.Marshal(app_conf.AppConf)
	if err != nil {
		log.Fatalf("nable to marshal config to YAML: %v", err)
	}

	filename := "config.yaml"

	// if f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644); err != nil {
	// 	log.Fatalf("os.OpenFile: %v", err)
	// } else {
	// 	if _, err := f.Write(b); err != nil {
	// 		log.Fatalf("f.Write: %v", err)
	// 	} else {
	// 		return
	// 	}
	// }
	if err := ioutil.WriteFile(filename, b, 0644); err != nil {
		log.Fatalf("ioutil.WriteFile: %v", err)
		return
	}
	return
}

// 提取组件数据到struct
func save() error {
	if err := sdk_taobao.Save(); err != nil {
		return err
	}
	if err := sdk_jingdong.Save(); err != nil {
		return err
	}
	if err := sdk_pinduoduo.Save(); err != nil {
		return err
	}
	if err := http_info.Save(); err != nil {
		return err
	}
	return nil
}
