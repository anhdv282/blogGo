package controller

import (
	"github.com/kataras/iris"
	"sample-project/model"
)

func (c *Controller) AdminGetPosts(ctx iris.Context) {


	ctx.View("/admin/blog/index.html")
}

func (c *Controller) AdminCreatePost(ctx iris.Context) {
	post := model.Post{}
	err := ctx.ReadForm(&post)
	if err != nil && !iris.IsErrPath(err) {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}

	ctx.Writef("Post: %#v", post)
}

func (c *Controller) AdminGetCreatePostPage(ctx iris.Context) {
	ctx.View("/admin/blog/create.html")
}
