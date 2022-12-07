/*
software:GoLand
PROJECT_NAME:goclass
DIR_PATH:mobile
file_name:find_element.go
date:2022/11/26 16:48:05
*/

package mobile

import (
	"android_automation/mobile/android"
	"fmt"
	"github.com/electricbubble/guia2"
)

func waitForElement(driver android.AndroidDriver, bySelector guia2.BySelector) (element android.AndroidElement, err error) {
	var ce error
	exists := func(d android.AndroidDriver) (bool, error) {
		element, ce = d.FindElement(bySelector)
		if ce == nil {
			return true, nil
		}
		// 如果直接返回 error 将直接终止 `driver.Wait`
		return false, nil
	}
	if err = driver.Wait(exists); err != nil {
		return nil, fmt.Errorf("%s: %w", err.Error(), ce)
	}
	return
}
