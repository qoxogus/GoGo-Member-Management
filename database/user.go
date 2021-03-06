package database

// User - 유저 관리 테이블
type User struct {
	Idx       uint   `gorm:"primary_key; auto_increment:true" json:"idx"`
	UserID    string `gorm:"type:varchar(255);unique;not null" json:"user_id"`
	Pw        string `gorm:"type:varchar(255);not null" json:"pw"`
	Name      string `gorm:"type:varchar(255);not null" json:"name"`
	IsManager bool   `gorm:"not null" sql:"DEFAULT:false" json:"is_manager"`
	// JoinedAt    time.Time `gorm:"nor null" sql:"DEFAULT:current_timestamp" json:"joined_at"`
}
