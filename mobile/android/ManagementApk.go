/*
software:GoLand
PROJECT_NAME:goclass
DIR_PATH:mobile/android
file_name:ManagementApk.go
date:2022/11/26 16:39:19
*/

package android

import (
	"fmt"
)

func InstallApk(driver AndroidDriver, apkPath string, reinstall ...bool) error {
	return driver.AppInstall(apkPath, reinstall...)
}

func RemoteApp(driver AndroidDriver, pkgName string, keepDataAndCache ...bool) error {
	return driver.AppUninstall(pkgName, keepDataAndCache...)
}

func FindApps(driver AndroidDriver) {
	name, err := driver.ActiveAppPackageName()
	if err != nil {
		panic(err)
	}
	fmt.Println(name)

}
