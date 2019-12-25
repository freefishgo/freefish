package conf

import (
	"encoding/json"
	"github.com/freefishgo/freeFish"
	"github.com/freefishgo/freeFish/middlewares/mvc"
	"os"
	"{ProjectName}/fishgo"
)

var Build *freeFish.ApplicationBuilder

type config struct {
	*freeFish.Config
	WebConfig *mvc.WebConfig
}

func init() {
	Build = freeFish.NewFreeFishApplicationBuilder()
	conf := new(config)
	f, err := os.Open("conf/app.conf")
	if err!=nil{
		panic(err.Error())
	}
	json.NewDecoder(f).Decode(conf)
	Build.Config = conf.Config
	fishgo.Mvc.Config = conf.WebConfig

}
