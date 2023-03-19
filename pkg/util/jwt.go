package util

import (
	"time"
	
	"github.com/dgrijalva/jwt-go"
	"github.com/siddontang/go/num"
	
	"jaingke2023.com/BlogService/pkg/settings"
)

//1.获取配置文件中的 jwt秘钥
var jwtSecret []byte = num.Int64ToBytes(settings.JwtSecret)

type Claims struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	jwt.StandardClaims
}

// GenerateToken  给定一个用户名和密码，生成并返回一个token，以及可能遇到的error
func GenerateToken(username, password string) (token string, err error) {
	//1.根据当前系统时间，设置生成 token 的到期时间
	now := time.Now()
	expireAt := now.Add(30 * time.Minute)
	
	claims := Claims{
		UserName: username,
		PassWord: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "blogService",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	
	token, err = tokenClaims.SignedString(jwtSecret)
	return
}

// ParseToken  解析token ，判断token是否合法
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(*jwt.Token) (interface{}, error) {
		
		return jwtSecret, nil
	})
	
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	
	return nil, err
	
}
