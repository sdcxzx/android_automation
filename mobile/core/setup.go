/*
software:GoLand
PROJECT_NAME:goclass
DIR_PATH:mobile/core
file_name:setup.go
date:2022/12/6 18:43:33
*/

package core

import (
	"android_automation/mobile"
	"android_automation/mobile/global"
)

func caseSetup() bool {
	cases := mobile.FindCases()
	//判断是否在cases内
	for _, v := range cases.Name {
		if v == "setup.yaml" {
			global.Logger.Info("找到setup，执行setup...........")
			return true
		}
	}
	global.Logger.Info("未找到setup...........")
	return false
}
