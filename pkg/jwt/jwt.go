package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var mySecret = []byte("夏天夏天悄悄过去")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var keyFunc = func(token *jwt.Token) (interface{}, error) {
	return mySecret, nil
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明的数据
	c := MyClaims{
		userID,
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(8760) * time.Hour).Unix(), // 过期时间
			Issuer: "bluebell", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象 必须是SigningMethodHS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	// jwt 由header payload sign（签名）组成，前面生成的token是包含了header ,payload,要用指定的secret签名之后获得完整的token
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT parse 解析
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("非法token")
}

// GenTokenAccessAndRefresh ⽣成access token 和 refresh token
func GenTokenAccessAndRefresh(userID int64, username string) (aToken, rToken string, err error) {
	// 创建⼀个我们⾃⼰的声明
	c := MyClaims{
		userID, // ⾃定义字段
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(8760) * time.Hour).Unix(), // 过期时间
			Issuer: "bluebell", // 签发⼈
		},
	}
	// 加密并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256,
		c).SignedString(mySecret)
	// refresh token 不需要存任何⾃定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(), // 过期时间
		Issuer:    "bluebell",                              // 签发⼈
	}).SignedString(mySecret)
	// 使⽤指定的secret签名并获得完整的编码后的字符串token
	return
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// refresh token⽆效直接返回
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}
	// 从旧access token中解析出claims数据
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)
	// 当access token是过期错误 并且 refresh token没有过期时就创建⼀个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenTokenAccessAndRefresh(claims.UserID, claims.Username)
	}
	return
}
