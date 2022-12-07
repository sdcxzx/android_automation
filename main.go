/*
software:GoLand
PROJECT_NAME:goclass
DIR_PATH:
file_name:main.go
date:2022/10/1 15:23:06
*/

package main

import (
	"android_automation/mobile/core"
	"android_automation/mobile/global"
	"go.uber.org/zap"
)

func main() {

	global.Logger = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.Logger)

	runCases()

}
