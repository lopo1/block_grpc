package core

import (
	"block_tool/user/server/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

const (
	ConfigEnv         = "GVA_CONFIG"
	ConfigDefaultFile = "./user/config.yaml"
	ConfigTestFile    = "./user/config-test.yaml"
	ConfigDevFile     = "./user/config-dev.yaml"
	ConfigReleaseFile = "./user/config-release.yaml"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
// Author [SliverHorn](https://github.com/SliverHorn)
func Viper(path ...string) *viper.Viper {
	var config string
	var mode string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.StringVar(&mode, "m", "dev", "choose config file.")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv("dev"); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
				switch mode {
				case "main":
					config = ConfigDefaultFile
					fmt.Printf("您正在使用模式的%s环境名称,config的路径为%s\n", mode, ConfigDefaultFile)
				case "dev":
					config = ConfigDevFile
					fmt.Printf("您正在使用模式的%s环境名称,config的路径为%s\n", mode, ConfigDevFile)
				case "release":
					config = ConfigReleaseFile
					fmt.Printf("您正在使用模式的%s环境名称,config的路径为%s\n", mode, ConfigReleaseFile)
				case "test":
					config = ConfigTestFile
					fmt.Printf("您正在使用模式的%s环境名称,config的路径为%s\n", mode, ConfigTestFile)
				}
			} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				config = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", ConfigEnv, config)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	//tt, _ := filepath.Abs(".")
	//fmt.Println("filepath.Abs(\"..\") = ", tt)
	return v
}
