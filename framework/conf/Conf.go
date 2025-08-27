package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	Http struct {
		Port int
	}
}

var Config *Conf // 全局公共变量

func InitConfig() {
	//var Config = &Conf{}
	file, err := os.ReadFile("/home/chenwenxi/npc/go_demos/gin-demo/framework/conf.yml")
	if err != nil {
		panic(err) // 中断程序
	}
	var errYaml = yaml.Unmarshal(file, &Config)
	if errYaml != nil {
		panic(errYaml) // 中断程序
	}
}
