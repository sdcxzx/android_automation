/*
software:GoLand
PROJECT_NAME:goclass
DIR_PATH:mobile
file_name:case_step.go
date:2022/11/26 16:42:54
*/

package mobile

import (
	"android_automation/mobile/android"
	"android_automation/mobile/global"
	"fmt"
	"github.com/electricbubble/guia2"
	"go.uber.org/zap"
	"strings"
)

// Step 用例步骤
/*
pkgName: 包名;

ele: 元素;

action: 操作	1: 点击 2: 输入;

eleType: 元素类型	1: id 2: xpath 3: text;

sendText: 输入的文本 可以为空;

apkPath: apk路径 可以为空

*/

func Step(pkgName string, ele string, action string, eleType string, sendText string, apkPath string) {
	var eleSelector android.AndroidElement
	var err error
	driver := android.DriverConn()
	err6 := android.RemoteApp(driver, pkgName, false)
	if err6.Error() == fmt.Sprintf("apk uninstalled: %s", "Failure [DELETE_FAILED_INTERNAL_ERROR]\n") {
		result := fmt.Sprintf("未找到%s，无需卸载", pkgName)
		global.Logger.Info("未找到应用，无需卸载", zap.String("result", result))
	}
	if apkPath != "" {
		err3 := android.InstallApk(driver, apkPath)
		if err3 != nil {
			global.Logger.Error("安装apk失败", zap.String("err", err3.Error()))
			return
		}
	}
	err2 := driver.AppLaunch("com.qq.reader")
	if err2 != nil {
		global.Logger.Error("启动应用失败", zap.String("err", err2.Error()))
	}
	// 通过元素定位器查找元素
	switch strings.ToUpper(eleType) {
	case "ID":
		// 通过id查找元素
		eleSelector, err = waitForElement(driver, guia2.BySelector{ResourceIdID: ele})
		if err != nil {
			fmt.Println(err)
			global.Logger.Error("ID查找元素失败", zap.Error(err), zap.String("ele", ele))
		}
	case "XPATH":
		// 通过xpath查找元素
		eleSelector, err = waitForElement(driver, guia2.BySelector{XPath: ele})
		if err != nil {
			fmt.Println(err)
			global.Logger.Error("XPATH查找元素失败", zap.Error(err), zap.String("元素", ele))
		}
	case "ClASSNAME":
		// 通过class查找元素
		eleSelector, err = waitForElement(driver, guia2.BySelector{ClassName: ele})
		if err != nil {
			fmt.Println(err)
			global.Logger.Error("ClassName查找元素失败", zap.Error(err), zap.String("ele", ele))
		}
	default:
		global.Logger.Info("请查看定位标签，目前只支持ID、XPATH、ClassName，定位元素是否正确")
	}
	// 通过元素执行操作
	switch strings.ToUpper(action) {
	case "CLICK":
		// 点击元素
		err = eleSelector.Click()
		if err != nil {
			global.Logger.Error("点击元素失败", zap.Error(err), zap.String("元素", ele))
		}
	case "SENDKEYS":
		// 输入内容
		err = eleSelector.SendKeys(sendText)
		if err != nil {
			global.Logger.Error("输入内容失败", zap.Error(err), zap.String("内容", sendText))
		}
	case "CLEAR":
		// 清空内容
		err = eleSelector.Clear()
		if err != nil {
			global.Logger.Error("清空内容失败", zap.Error(err))
		}
	default:
		global.Logger.Info("请查看操Click、SendKeys、Clear是否正确")
	}

}
