/*
software:GoLand
PROJECT_NAME:goclass
DIR_PATH:mobile/android
file_name:driver_conn.go
date:2022/11/22 15:31:51
*/

package android

import (
	"android_automation/mobile/global"
	"github.com/electricbubble/guia2"
	"go.uber.org/zap"
	"os/exec"
)

func DriverConn() AndroidDriver {
	// 启动appium服务
	cmd := exec.Command("adb", "shell", "nohup", "am", "instrument", "-w", "io.appium.uiautomator2.server.test/androidx.test.runner.AndroidJUnitRunner", ">/sdcard/uia2server.log", "2>&1", "&")
	global.Logger.Info("cmd:", zap.String("cmd", cmd.String()))
	_, err2 := cmd.Output()
	if err2 != nil {
		global.Logger.Error("启动appium服务失败", zap.String("err", err2.Error()))
	}
	driver, err := guia2.NewUSBDriver()
	if err != nil {
		global.Logger.Error("连接手机失败", zap.String("err", err.Error()))
	}
	return driver
}
