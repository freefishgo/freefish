package controllers

import (
	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

// 实现mvc控制器的处理Main为控制器 {Controller}的值
type homeController struct {
	mvc.Controller
}

// 注册控制器
func init() {
	mvc.AddHandlers(&homeController{})
}

// Index为{Action}的值 该方法的默认路由为/home/Index 最后的单词为请求方式  默认为Get  如需设置为post 例：IndexPost
func (home *homeController) Index() {
	home.Data["Website"] = "freefishgo.com"
	home.Data["Email"] = "a1085052074@qq.com"
	// 调用模板引擎   默认模板地址为{ Controller}/{Action}.fish    即为home/Index， home.UseTplPath()等效于home.UseTplPath("home/Index")
	home.UseTplPath()
}
