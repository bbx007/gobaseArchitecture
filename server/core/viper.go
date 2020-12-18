package core

import (
	"flag"
	"fmt"
	"gin-vue-admin/global"
	_ "gin-vue-admin/packfile"
	"gin-vue-admin/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				config = utils.ConfigFile
				fmt.Printf("/**\n " +
					"*                      江城子 . 程序员之歌\n " +
					"*\n " +
					"*                  十年生死两茫茫，写程序，到天亮。\n " +
					"*                      千行代码，Bug何处藏。\n " +
					"*                  纵使上线又怎样，朝令改，夕断肠。\n " +
					"*\n " +
					"*                  领导每天新想法，天天改，日日忙。\n " +
					"*                      相顾无言，惟有泪千行。\n " +
					"*                  每晚灯火阑珊处，夜难寐，加班狂。\n" +
					"*/您正在使用config的默认值,config的路径为%v\n", utils.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
