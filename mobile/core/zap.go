/*
software:GoLand
PROJECT_NAME:goclass
DIR_PATH:mobile/core
file_name:zap.go
date:2022/12/5 16:06:52
*/

package core

import (
	"android_automation/mobile/core/internal"
	"android_automation/mobile/zapConfig"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() (logger *zap.Logger) {
	if ok, _ := pathExists(zapConfig.LoggerFile.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", zapConfig.LoggerFile.Director)
		_ = os.Mkdir(zapConfig.LoggerFile.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if zapConfig.LoggerShowLine.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
