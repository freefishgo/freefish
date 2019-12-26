package controllers

import (
	"github.com/freefishgo/freeFishGo/middlewares/mvc"
	"{{[.ProjectName]}}/fishgo"
)

// 实现mvc控制器的处理Main为控制器 {Controller}的值
type HomeController struct {
	mvc.Controller
}

// 把控制器注入到mvc框架中去
func init() {
	fishgo.Mvc.AddHandlers(&HomeController{})
}
// MyControllerActionStrut为{Action}的值 该方法的默认路由为/Home/MyControllerActionStrut 最后的单词为请求方式  该例子为Post请求
func (c *HomeController) Index() {
	c.Data["Website"] = "freefishgo.com"
	c.Data["Email"] = "a1085052074@qq.com"
	// 调用模板引擎   默认模板地址为{ Controller}/{Action}.fish    即为Main/MyControllerActionStrut， c.UseTplPath()等效于c.UseTplPath("Home/MyControllerActionStrut")
	c.UseTplPath()
}