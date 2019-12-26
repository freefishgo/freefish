package fishgo

import (
	"github.com/freefishgo/freeFishGo/middlewares/mvc"
)

var Mvc *mvc.MvcApp

// 实例化一个mvc中间件服务
func init() {
	Mvc = mvc.NewFreeFishMvcApp()
}
