package model

import "time"

type User struct {
	ID        uint64    `gorm:"primarykey" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Email     string    `gorm:"uniqueIndex;size:100;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Posts     []Post    `gorm:"foreignKey:UserID" json:"posts,omitempty"` //如果字段值为空，则在 JSON 序列化时忽略该字段
	Comments  []Comment `gorm:"foreignKey:UserID" json:"comments,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Post struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Title     string    `gorm:"size:200;not null" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Comments  []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Comment struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	PostID    uint      `gorm:"not null;index" json:"post_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Post      Post      `gorm:"foreignKey:PostID" json:"post"`
	CreatedAt time.Time `json:"created_at"`
}
