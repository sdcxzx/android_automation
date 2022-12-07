/*
software:GoLand
PROJECT_NAME:goclass
DIR_PATH:mobile/core
file_name:read_cases_yaml.go
date:2022/12/6 17:30:16
*/

package core

import (
	"android_automation/mobile"
	"android_automation/mobile/global"
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var casesFile []string

func caseFilePath() []string {
	// 读取yaml文件
	cases := mobile.FindCases()
	casesPath := global.CasesPath()
	for _, v := range cases.Name {
		casesFile = append(casesFile, casesPath+"/"+v)
	}
	fmt.Println(casesFile)
	return casesFile
}

func ReadCases() {
	if caseSetup() {
		var yamlMain = mobile.Setup{}
		// 读取setup.yaml
		content, err := ioutil.ReadFile(global.CasesPath() + "/setup.yaml")
		checkError(err)
		err = yaml.Unmarshal(content, &yamlMain)
		checkError(err)
		global.Logger.Info("yamlMain", zap.Any("yamlMain", yamlMain))
		for _, setup := range yamlMain.Setup {
			mobile.Step("", setup.Element, setup.Action, setup.Type, "", "")
		}
		for _, caseFile := range caseFilePath() {
			// 读取yaml文件
			global.Logger.Info("setup.yaml-1")
			content, err := ioutil.ReadFile(caseFile)
			checkError(err)
			err = yaml.Unmarshal(content, &yamlMain)
			checkError(err)
			//fmt.Println(err, &yamlMain)
			for _, setup := range yamlMain.Setup {
				mobile.Step("", setup.Element, setup.Action, setup.Type, "", "")
			}
		}
	} else {
		//没有setup则走
		var caseStep = mobile.CasesStep{}
		//清除struct中的值
		for _, caseFile := range caseFilePath() {
			// 读取yaml文件
			content, err := ioutil.ReadFile(caseFile)
			checkError(err)
			err = yaml.Unmarshal(content, &caseStep)
			checkError(err)
			for _, setup := range caseStep.Step {
				mobile.Step("", setup.Element, setup.Action, setup.Type, "", "")
			}
		}
	}

}
func checkError(err error) {
	if err != nil {
		global.Logger.Error("yamlFile.Get err", zap.Error(err))
	}
}
