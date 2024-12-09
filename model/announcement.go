package model

type Announcement struct {
	ID      int    `json:"id"`
	Title   string `gorm:"size:255" json:"title"`
	Content string `gorm:"size:4096" json:"content"`
}
