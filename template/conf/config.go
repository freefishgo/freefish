package conf

import (
	"encoding/json"
	"github.com/freefishgo/freeFishGo"
	"github.com/freefishgo/freeFishGo/middlewares/mvc"
	"os"
	"{{[.ProjectName]}}/fishgo"
)
// 管道
var Build *freeFishGo.ApplicationBuilder
// 应用配置
type config struct {
	// 管道配置
	*freeFishGo.Config
	// mvc中间件的配置
	WebConfig *mvc.MvcWebConfig
}

func init() {
	Build = freeFishGo.NewFreeFishApplicationBuilder()
	conf := new(config)
	os.Chdir(`{{[.Chdir]}}`)
	f, err := os.Open("conf/app.conf")
	if err!=nil{
		panic(err.Error())
	}
	json.NewDecoder(f).Decode(conf)
	Build.Config = conf.Config
	fishgo.Mvc.Config = conf.WebConfig

}
