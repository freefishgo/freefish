package middlewares

//
//import (
//	"github.com/freefishgo/freefishgo"
//)
//
//// 例子： 组装一个Middleware服务，实现允许跨域请求
//type Mid struct {
//}
//
//// 中间件实现允许跨域请求
//func (m *Mid) Middleware(ctx *freefishgo.HttpContext, next freefishgo.Next) *freefishgo.HttpContext {
//	ctx.Response.Header().Add("Access-Control-Allow-Origin","*")
//  ctx.Response.Header().Add("Access-Control-Allow-Methods", "GET, POST, DELETE,PUT")
//	if http.MethodOptions == ctx.Request.Method {
//	return ctx
//	}
//	ctxtmp := next(ctx)
//	return ctxtmp
//}
//
//// 中间件注册时调用函数进行该中间件最后的设置
//func (*Mid) LastInit(config *freefishgo.Config) {
//	//panic("implement me")
//}
