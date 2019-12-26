package controllers

import (
	"github.com/freefishgo/freeFishGo/middlewares/mvc"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"{{[.ProjectName]}}/fishgo"
)

type staticController struct {
	mvc.Controller
}

// 控制器注册
func init() {
	fishgo.Mvc.AddHandlers(&staticController{})
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
		io.Copy(static.Response,f)
	} else {
		static.Response.WriteHeader(404)
		static.Response.Write([]byte(err.Error()))
	}
}
// 重写 指定动作的路由 该方法会在路由注册时调用
func (static *staticController) OverwriteRouter() []*mvc.ControllerActionInfo {
	tmp := make([]*mvc.ControllerActionInfo, 0)
	tmp = append(tmp, &mvc.ControllerActionInfo{RouterPattern: "static/{path:allString}", ControllerActionFuncName: "StaticFile"})
	return tmp
}
