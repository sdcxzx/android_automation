/*
software:GoLand
PROJECT_NAME:goclass
DIR_PATH:mobile
file_name:find_case.go
date:2022/12/5 18:15:14
*/

package mobile

import (
	"android_automation/mobile/global"
	"go.uber.org/zap"
	"io/ioutil"
)

type Cases struct {
	Name []string
}

func FindCases() *Cases {
	//遍历path下所以文件，并存储到struct中
	var caseName Cases
	cases, err := ioutil.ReadDir(global.CasesPath())
	if err != nil {
		global.Logger.Error(err.Error())
	}
	for _, v := range cases {
		if !v.IsDir() {
			//存入struct
			caseName.Name = append(caseName.Name, v.Name())
			global.Logger.Info("caseName:", zap.Any("Name", v.Name()))
		}
	}
	return &caseName
}
