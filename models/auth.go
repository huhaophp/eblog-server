package models

// Auth Struct
type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth 检查授权
func CheckAuth(username string) (auth Auth) {
	db.Select("id,username,password").Where(Auth{Username: username}).First(&auth)
	return
}

// GetAuthInfoById 根据 id 获取用户信息
func GetAuthInfoById(id int) (auth Auth) {
	db.Select("id,username").Where(Auth{ID: id}).First(&auth)
	return
}
