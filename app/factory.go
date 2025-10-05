package app

import "gorm.io/gorm"

func UserFactory(db *gorm.DB, setter func(*User)) (*User, error) {
	u := new(User)

	u.Nickname = "test user"
	u.Email = "test@example.com"

	if setter != nil {
		setter(u)
	}

	if db != nil {
		if err := db.Create(u).Error; err != nil {
			return nil, err
		}
	}

	return u, nil
}

func PostFactory(db *gorm.DB, setter func(*Post)) (*Post, error) {
	p := new(Post)

	p.Content = "post content"

	if setter != nil {
		setter(p)
	}
	if db != nil {
		if err := db.Create(p).Error; err != nil {
			return nil, err
		}
	}

	return p, nil
}

func CommentFactory(db *gorm.DB, setter func(c *Comment)) (*Comment, error) {
	c := new(Comment)

	c.Content = "comment content"

	if setter != nil {
		setter(c)
	}

	if db != nil {
		if err := db.Create(c).Error; err != nil {
			return nil, err
		}
	}

	return c, nil
}
