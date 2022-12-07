/*
software:GoLand
PROJECT_NAME:goclass
DIR_PATH:mobile
file_name:setup.go
date:2022/11/28 09:19:04
*/

package mobile

type Setup struct {
	Setup []setup `yaml:"setup"`
}
type CasesStep struct {
	Step []step `yaml:"step"`
}
type setup struct {
	StepName     string `yaml:"stepName"`
	AppInstall   bool   `yaml:"appInstall,omitempty"`
	AppUninstall bool   `yaml:"appUninstall,omitempty"`
	Type         string `yaml:"type,omitempty"`
	Element      string `yaml:"element,omitempty"`
	Action       string `yaml:"events,omitempty"`
	SendText     string `yaml:"sendText,omitempty"`
	apkPath      string `yaml:"apkPath,omitempty"`
	appName      string `yaml:"appName,omitempty"`
}

type step struct {
	StepName string `yaml:"stepName"`
	Type     string `yaml:"type,omitempty"`
	Element  string `yaml:"element,omitempty"`
	Action   string `yaml:"events,omitempty"`
	SendText string `yaml:"sendText,omitempty"`
	appName  string `yaml:"appName,omitempty"`
}
