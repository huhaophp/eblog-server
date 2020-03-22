package models

// Auth Struct
type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth 检查授权
func GetAuthByUsername(username string) (auth Auth) {
	db.Select("id,username,password").Where(Auth{Username: username}).First(&auth)
	return
}
