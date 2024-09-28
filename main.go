package main

import (
	"fmt"
	"test/pkg"

	"github.com/nbcx/flag"
	"github.com/nbcx/go-config/fig"
	"github.com/nbcx/log"
)

func main() {
	// 指定要解析的配置文件
	// viper.SetConfigFile("test.ini")
	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatal(err)
	// }
	// // 默认section为default
	// fmt.Printf("app_mode:%v\n", viper.Get(`default.app_mode`))
	// fmt.Printf("server.protocol:%v\n", viper.GetString(`server.protocol`))

	fig.SetConfigFile("app.ini")
	if err := fig.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	// 默认section为default
	log.Info("appname::::::", fig.Sub("default").Get("appname"))
	fmt.Printf("app_mode2:%v\n", fig.Get(`default.appname`))

	// log.Info(fig.AllSettings())
	// fmt.Printf("server.protocol2:%v\n", fig.Get(`server.protocol`))

	// fig.SetConfigType("ini") // or viper.SetConfigType("YAML")
	// fig.AddConfigPath(".")
	// fig.SetConfigName("app")

	// err := fig.ReadInConfig()

	// log.Error("err", err)

	// fmt.Println(">>>>>>>>:", fig.Get("appname"))

}

func main2() {
	fig.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	fig.AddConfigPath(".")
	fig.SetConfigName("conf.yaml")

	_ = fig.ReadInConfig()

	// config.Get("name") // this would be "steve"

	flag.Int("name", 1234, "help message for flagname")

	flag.Parse()
	fig.BindPFlags(flag.CommandLine)

	i := fig.Get("name") // retrieve values from viper instead of pflag

	fmt.Println(">>>>>>>>", i)

	co := &pkg.Config{}
	_ = fig.Unmarshal(co)

	fmt.Println("co", co)

	log.Info("hello retrieve values from viper instead of pflag", 0000, "hell")
}
