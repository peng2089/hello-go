package data

type User struct {
	ID       uint `gorm:"primaryKey"`
	Nickname string
}
type UserLocalAuth struct {
	ID        uint `gorm:"primaryKey"`
	UserId    uint
	Username  string
	Password  string
	CreatedAt int64 `gorm:"autoCreateTime"`
}
type UserOauth struct {
	ID               uint `gorm:"primaryKey"`
	UserId           uint
	OauthName        string
	OauthId          string
	OauthAccessToken string
	OauthExpires     int64
	CreatedAt        int64 `gorm:"autoCreateTime"`
}
