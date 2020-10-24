package highjacker

import (
	"github.com/kataras/iris/v12"
)

func jd(ctx iris.Context) {
	ctx.Text("jd highjack")
}
