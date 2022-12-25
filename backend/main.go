package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	template := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	app.HandleDir("/assets", "./backend/web/assets")
	app.OnAnyErrorCode(func(context iris.Context) {
		context.ViewData("message", context.Values().GetStringDefault("message", "Page Error"))
		context.ViewLayout("")
		context.View("shared/error.html")
	})

	app.Run(
		iris.Addr("localhost:9001"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
