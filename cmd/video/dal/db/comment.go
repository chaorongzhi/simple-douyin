package db

type Comment struct {
	CommentID uint `gorm:"primarykey"`
	UserID    uint
	VideoID   uint
	Content   string `gorm:"not null"`
	CreatedAt int
}





