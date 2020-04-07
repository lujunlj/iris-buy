package main

import (
	"github.com/kataras/iris"
)

func main() {
	//1.创建iris实例
	app := iris.New()
	//2.设置代码错误等级
	app.Logger().SetLevel("debug")
	//3.注册模板
	template := iris.HTML("./webapp/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	//4.设置静态目录且设置出现异常跳转
	app.StaticWeb("/static", "./webapp/views/static")
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("./webapp/views/shared/error.html")
	})
	//5.注册控制器

	//6.启动服务（设置地址）
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
