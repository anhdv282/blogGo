package model

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"time"
)

type Post struct {
	// Tao shema blog
	TableName []byte `sql:"blog.post"`
	// ID bài viết
	Id string `sql:",pk"`
	// Tiêu đề
	Title string 
	// Thumbnail
	Thumbnail string 
	// Nội dung
	Content string 
	// Tác giả
	CreatedBy string 
	// Thời gian tạo 
	CreatedAt time.Time 
	// Ngày xuất bản
	PublishedAt time.Time 
	// Trạng thái: 0 - nháp, 1 - xuất bản 
	Status int32 `sql:",notnull"`
}

func GetFirstPost(model *Post, db *pg.DB){
	err := db.Model(model).Column("post.*").First()
	if err != nil {
		panic(err)
	}
	fmt.Println(model.Content)
}

func GetPosts(db *pg.DB) ([]Post, error) {
	var posts []Post
	_, err := db.Query(&posts, `SELECT * FROM blog.post`)
	return posts, err
}
