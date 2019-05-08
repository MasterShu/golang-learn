package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:root@/goweb?charset=utf8mb4,utf8&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{})
}

func retrieve(id int) (post Post, err error) {
	post = Post{}
	err = Db.Debug().Raw("select * from posts where id = ?", id).Scan(&post).Error
	return post, err
}

func (post *Post) create() (err error) {
	err = Db.Create(&post).Error
	return err
}

func (post *Post) update() (err error) {
	err = Db.Save(&post).Error
	return err
}

func (post *Post) delete() (err error) {
	err = Db.Delete(&post).Error
	return err
}
