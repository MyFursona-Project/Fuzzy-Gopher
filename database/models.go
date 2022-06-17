package database

import "time"

type User struct {
	UserID    string `gorm:"primaryKey"`
	EMail     string
	Name      string
	Password  string
	CreatedAt time.Time
	locked    bool
	// foreign keys
	Verify           []Verify  `gorm:"foreignKey:UserID"`
	ProfileUserID    []Profile `gorm:"foreignKey:UserID"`
	ProfileName      []Profile `gorm:"foreignKey:Name"`
	ProfileCreatedAt []Profile `gorm:"foreignKey:CreatedAt"`
	Fursona          []Fursona `gorm:"foreignKey:UserID;foreignKey:OwnerID"`
	Like             []Like    `gorm:"foreignKey:UserID"`
}
type Verify struct {
	ID        string `gorm:"primaryKey"`
	UserID    string // User.UserID
	EMailCode string
}
type Profile struct {
	UserID    string    `gorm:"primaryKey"` // -> User.UserID
	Name      string    // -> User.Name
	CreatedAt time.Time // -> User.CreatedAt
	Location  string
	State     string
	Gender    string
	Sexuality string
}
type Fursona struct {
	FursonaID string `gorm:"primaryKey"`
	OwnerID   string // -> User.UserID
	Name      string
	Gender    string
	Sexuality string
	// foreign keys
	FursonaMedia []FursonaMedia `gorm:"foreignKey:FursonaID"`
	Like         []Like         `gorm:"foreignKey:FursonaID"`
}

type FursonaMedia struct {
	MediaID   string `gorm:"primaryKey"`
	FursonaID string // -> Fursona.FursonaID
	URLHigh   string // original / PNG lossless
	URLMedium string // webp lossless
	URLLow    string // webp resized / compressed
	Sensitive bool   // nsfw
}

type Like struct {
	LikeID    string `gorm:"primaryKey"`
	UserID    string // -> User.UserID
	FursonaID string // -> Fursona.FursonaID
}
