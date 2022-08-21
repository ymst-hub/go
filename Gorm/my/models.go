package my
import (
	"gorm.io/gorm"

)
//user model
type User struct{
	gorm.Model
	Account string
	Name string
	Password string
	Message string
}

//post model
type Post struct{
	gorm.Model
	Address string
	Message string
	UserId int
	GroupId int
}

//group model
type Group struct{
	gorm.Model
	UserId int
	Name string
	Message string
}

//Comment
type Comment struct{
	gorm.Model
	UserId int
	PostId int
	Message string
}

//CommentJoin

type CommentJoin struct{
	Comment
	User
	Post
}

//User ログインするユーザー情報
//Post 投稿する動画の情報
//Group グループの情報
//Comment 投稿データに付けられるコメントの情報
