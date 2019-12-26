package routers

import (
	"github.com/freefishgo/freeFishGo/middlewares/mvc"
	"{{[.ProjectName]}}/fishgo"
)

func init() {
	// 注册主路由  可多主路由格式      但 主页面 设置只有第一个有效
	fishgo.Mvc.AddMainRouter(&mvc.MainRouter{
		RouterPattern: "/{ Controller}/{Action}",
		HomeController:"Home",
		IndexAction:"Index",})
}
