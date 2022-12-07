/*
software:GoLand
PROJECT_NAME:android_automation
DIR_PATH:mobile
file_name:runCases.go
date:2022/12/7 16:45:21
*/

package main

import (
	"android_automation/mobile/core"
	"android_automation/mobile/global"
	"go.uber.org/zap"
	"time"
)

func runCases() {
	global.Logger.Info("开始执行用例...........")
	startTime := time.Now().Unix()
	core.ReadCases()
	endTime := time.Now().Unix()
	global.Logger.Info("用例执行完毕...........")
	global.Logger.Info("用例执行时间为：", zap.Any("time", endTime-startTime))
}
