package models

type Auth struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	count := 0
	db.Model(Auth{}).Where(Auth{Username: username, Password: password}).Count(&count)
	
	//auth := Auth{}
	//db.Model(Auth{}).Select("id").Where("username = ? and password = ?", username, password).First(&auth)
	//return auth.Id > 0
	return count > 0
}
