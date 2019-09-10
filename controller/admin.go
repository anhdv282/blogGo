package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"sample-project/model"
)

func (c *Controller) AdminGetPosts(ctx iris.Context) {
	var post model.Post
	post = model.GetFirstPost(post,c.DB)
	fmt.Println(post.Content, post.Title)
	//ctx.ViewData("Title", post.Title)
	ctx.ViewData("post", post)
	ctx.View("/admin/blog/index.html")

}

func (c *Controller) AdminCreatePost(ctx iris.Context) {
	fmt.Println("AAAAAAAAAA")
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
