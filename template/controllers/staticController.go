package controllers

import (
	"github.com/freefishgo/freefishgo/middlewares/mvc"
	"io"
	"os"
	"path/filepath"
)

type staticController struct {
	mvc.Controller
}

// 控制器注册
func init() {
	static := staticController{}
	// 重写 指定动作的路由 该方法会在路由注册时调用
	static.ActionRouterList = append(static.ActionRouterList,
		&mvc.ActionRouter{RouterPattern: "static/{path:allString}",
			ControllerActionFuncName: "StaticFile"})
	mvc.AddHandlers(&staticController{})
}

type data struct {
	Path string `json:"path"`
}

// 提供静态资源服务
//
// mvc框架会分析请求数据      并把数据注入到d 中
//
// 由于重写了该动作的路由为：static/{path:allString}  所以d.Path 即为请求url中出 static/ 后面的部分
func (static *staticController) StaticFile(d *data) {
	if f, err := os.Open(filepath.Join("static", d.Path)); err == nil {
		//static.Response.Header().Set("Cache-Control","max-age=3600")
		switch filepath.Ext(d.Path) {
		case ".css":
			static.Response.Header().Set("Content-Type", "text/css")
			break
		case ".js":
			static.Response.Header().Set("Content-Type", "application/javascript")
			break
		}
		io.Copy(static.Response, f)
	} else {
		static.Response.WriteHeader(404)
		static.Response.Write([]byte(err.Error()))
	}
}
