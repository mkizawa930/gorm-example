package app

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewTestDB() (*gorm.DB, error) {
	driver := postgres.Open("host=localhost user=postgres password=postgres dbname=test port=65433 sslmode=disable TimeZone=Asia/Tokyo")
	return gorm.Open(driver, nil)
}

func MustCreate(t *testing.T, db *gorm.DB, model any) {
	if err := db.Create(model).Error; err != nil {
		t.Fatal(err)
	}
}

func Test_CreateUser(t *testing.T) {
	db, err := NewTestDB()
	if err != nil {
		t.Fatal(err)
	}
	tx := db.Begin()
	defer tx.Rollback()

	user := &User{
		Nickname: "test user",
		Email:    "test@example.com",
	}

	if err := tx.Create(user).Error; err != nil {
		t.Fatal(err)
	}

}

func Test_CreatePost(t *testing.T) {
	db, err := NewTestDB()
	if err != nil {
		t.Fatal(err)
	}
	tx := db.Begin()
	defer tx.Rollback()

	user := lo.Must(UserFactory(tx, nil))

	post, err := PostFactory(nil, func(p *Post) {
		p.UserID = user.ID
	})
	assert.Nil(t, err)

	if err := tx.Create(post).Error; err != nil {
		t.Fatal(err)
	}

}

func Test_CreateComment(t *testing.T) {
	db, err := NewTestDB()
	if err != nil {
		t.Fatal(err)
	}
	tx := db.Begin()
	defer tx.Rollback()

	user := lo.Must(UserFactory(tx, nil))

	post := lo.Must(PostFactory(tx, func(p *Post) {
		p.UserID = user.ID
	}))

	comment := lo.Must(CommentFactory(nil, func(c *Comment) {
		c.UserID = user.ID
		c.PostID = post.ID
	}))

	if err := tx.Create(comment).Error; err != nil {
		t.Fatal(err)
	}
}

func Test_LikeUserPost(t *testing.T) {
	db, err := NewTestDB()
	if err != nil {
		t.Fatal(err)
	}
	tx := db.Begin()
	defer tx.Rollback()

	user1 := lo.Must(UserFactory(tx, func(u *User) {
		u.Email = "user1@example.com"
	}))

	user2 := lo.Must(UserFactory(tx, func(u *User) {
		u.Email = "user2@example.com"
	}))

	post := lo.Must(PostFactory(tx, func(p *Post) {
		p.UserID = user1.ID
	}))

	// 投稿をlikeする
	err = tx.Model(post).Association("Likes").Append([]Like{{UserID: user2.ID}})
	if err != nil {
		t.Fatal(err)
	}

	var got Post
	if err = tx.Preload("Likes").First(&got, post.ID).Error; err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(got.Likes))
}

func Test_LikeUserComment(t *testing.T) {
	db, err := NewTestDB()
	if err != nil {
		t.Fatal(err)
	}
	tx := db.Begin()
	defer tx.Rollback()

	user1 := lo.Must(UserFactory(tx, func(u *User) {
		u.Email = "user1@example.com"
	}))

	user2 := lo.Must(UserFactory(tx, func(u *User) {
		u.Email = "user2@example.com"
	}))

	post := lo.Must(PostFactory(tx, func(p *Post) {
		p.UserID = user1.ID
	}))

	comment := lo.Must(CommentFactory(tx, func(c *Comment) {
		c.UserID = user2.ID
		c.PostID = post.ID
	}))

	if err := tx.Model(comment).Association("Likes").Append([]Like{{UserID: user1.ID}}); err != nil {
		t.Fatal(err)
	}

	var got Comment
	if err := tx.Preload("Likes").First(&got, comment.ID).Error; err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(got.Likes))
}

func Test_GetPostByID(t *testing.T) {
	db, err := NewTestDB()
	if err != nil {
		t.Fatal(err)
	}
	tx := db.Begin()
	defer tx.Rollback()

	user1 := lo.Must(UserFactory(tx, func(u *User) {
		u.Email = "user1@example.com"
	}))
	user2 := lo.Must(UserFactory(tx, func(u *User) {
		u.Email = "user2@example.com"
	}))

	post := lo.Must(PostFactory(tx, func(p *Post) {
		p.UserID = user1.ID
		p.Comments = []Comment{
			{
				UserID:  user2.ID,
				Content: "hoge",
			},
		}
	}))

	var got Post
	if err := tx.First(&got, post.ID).Error; err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, got.ID > 0)
	assert.Equal(t, 1, len(post.Comments))
}

func Test_GetCommentByID(t *testing.T) {
	db, err := NewTestDB()
	if err != nil {
		t.Fatal(err)
	}
	tx := db.Begin()
	defer tx.Rollback()

	user := lo.Must(UserFactory(tx, nil))

	post := lo.Must(PostFactory(tx, func(p *Post) {
		p.UserID = user.ID
	}))

	comment := lo.Must(CommentFactory(tx, func(c *Comment) {
		c.UserID = user.ID
		c.PostID = post.ID
	}))

	var got Comment
	if err := tx.First(&got, comment.ID).Error; err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, post.ID, got.PostID)
	assert.Equal(t, comment.ID, got.ID)
}
