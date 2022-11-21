package http_info

import (
	"fmt"
	"os/exec"
	"runtime"
)

var commands = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func openUrl(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("%s 系统不知道怎么打开,请反馈！", runtime.GOOS)
	}
	cmd := exec.Command(run, uri)
	return cmd.Start()
}

func lanchToTaobaoOauth(uuidStr string) error {
	uri := fmt.Sprintf(
		"https://oauth.m.taobao.com/authorize?response_type=code&client_id=23530087&redirect_uri=https://vip.lyhuilin.com/sdk/taobao/AuthGet&state=%s&view=wap",
		uuidStr)
	return openUrl(uri)
}
func lanchToWeipinhuiOauth(uuidStr string) error {
	uri := fmt.Sprintf(
		"https://auth.vip.com/oauth2/authorize?response_type=code&client_id=a762893d&redirect_uri=https://vip.lyhuilin.com/sdk/weipinhui/AuthGet&state=%s&view=wap",
		uuidStr,
	)
	return openUrl(uri)
}
