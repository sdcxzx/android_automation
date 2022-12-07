/*
software:GoLand
PROJECT_NAME:goclass
DIR_PATH:mobile/global
file_name:get_path.go
date:2022/12/6 17:33:23
*/

package global

import "os"

func PwdPath() string {

	dir, err := os.Getwd()
	if err != nil {
		Logger.Error(err.Error())
	}
	return dir

}

func CasesPath() string {

	return PwdPath() + "/mobile/cases"

}
