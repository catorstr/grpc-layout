package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// jwt 自定义结构体
type MyClaims struct {
	UserId     int64  `json:"user_id"`
	Account    string `json:"account"`
	UserName   string `json:"username"`
	UserAvatar string `json:"user_avatar"`
	Operation  int32  `json:"operation"`
	ThreeToken string `json:"three_token"`
	jwt.StandardClaims
}

// 创建token
func GeToken(userId int64, account, username, user_avatar string, operation int32, three_token string, secretKey []byte) (string, error) {
	//创建声明
	c := MyClaims{
		userId,
		account,
		username,
		user_avatar,
		operation,
		three_token,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), //有效期2小时
			Issuer:    "helloword-app",                            //签发人 写项目名称吧
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c) //哈希256签名算法
	//使用指定的secret秘钥签发并且获得完整的编码后的字符串token
	//参数为一个切片才行，其他的报错，不知道为啥
	return token.SignedString(secretKey) //secret按实际填写

}

func ParseToken(tokenString string, secretKey []byte) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { //校验token
		return claims, nil
	}
	return nil, errors.New("invaild token")
}
