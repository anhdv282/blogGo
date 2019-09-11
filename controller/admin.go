package controller

import (
	"github.com/kataras/iris"
	"sample-project/model"
)

func (c *Controller) AdminGetPosts(ctx iris.Context) {
	//var posts []model.Post
	//model.GetFirstPost(&post,c.DB)
	posts, err := model.GetPosts(c.DB)
	if err == nil {
		ctx.ViewData("posts", posts)
	}
	ctx.View("/admin/blog/index.html")
}

func (c *Controller) AdminCreatePost(ctx iris.Context) {
	post := model.Post{}
	err := ctx.ReadForm(&post)
	if err != nil && !iris.IsErrPath(err) /* see: https://github.com/kataras/iris/issues/1157 */ {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
	ctx.Redirect("/admin/blog/index.html")
}

func (c *Controller) AdminGetCreatePostPage(ctx iris.Context) {
	ctx.View("/admin/blog/create.html")
}
