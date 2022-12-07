/*
software:GoLand
PROJECT_NAME:goclass
DIR_PATH:mobile/zapConfig
file_name:zapConfig.go
date:2022/12/5 16:17:43
*/

package zapConfig

type Config struct {
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
}
