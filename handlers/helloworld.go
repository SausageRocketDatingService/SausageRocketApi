package handlers

import (
	"github.com/kataras/iris"
)

// HelloWorld return "Hello World!" html string
func HelloWorld(ctx iris.Context) {
	ctx.HTML("hello world!!")
}
