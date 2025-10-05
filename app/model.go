package app

import "gorm.io/gorm"

type UserID = uint

// ユーザーモデル
type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex"`
	Nickname string
	Auth     *Auth
}

// 認証モデル
type Auth struct {
	gorm.Model
	UserID   UserID
	Provider string
	User     User
}

type PostID = uint

// 投稿モデル
type Post struct {
	gorm.Model
	UserID   UserID
	Content  string
	Comments []Comment
	Likes    []Like `gorm:"many2many:posts_likes;"`
}

// コメントモデル
type Comment struct {
	gorm.Model
	PostID  PostID
	UserID  UserID
	Content string
	User    User
	Post    Post
	Likes   []Like `gorm:"many2many:comments_likes;"`
}

// お気に入りモデル
type Like struct {
	gorm.Model
	UserID UserID
	User   User
}
