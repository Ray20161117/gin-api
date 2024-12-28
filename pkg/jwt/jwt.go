/**
 * JWT 工具类
 */
package jwt

import (
	"errors"
	"fmt"
	"gin-api/config/constant"
	config "gin-api/config/yaml_config"
	"gin-api/models/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userStdClaims struct {
	entity.JwtAdminDto
	jwt.StandardClaims
}

// token过期时间
var TokenExpireDuration time.Duration

func init() {
	TokenExpireDuration = time.Duration(config.Cfg.App.JwtTokenExpire) * time.Hour
}

// token密钥
var Secret = []byte(config.Cfg.App.JwtSecret)
var (
	ErrAbsent  = "token absent"  // 令牌不存在
	ErrInvalid = "token invalid" //令牌无效
)

// 根据用户信息生成token
func GenerateTokenByAdmin(admin entity.SysAdmin) (string, error) {
	var jwtAdmin = entity.JwtAdminDto{
		ID:       admin.ID,
		Username: admin.Username,
		Nickname: admin.Nickname,
		Icon:     admin.Icon,
		Email:    admin.Email,
		Phone:    admin.Phone,
		Note:     admin.Note,
	}
	c := userStdClaims{
		jwtAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //过期时间
			Issuer:    "admin",                                    //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

// ValidateToken 解析JWT
func ValidateToken(tokenString string) (*entity.JwtAdminDto, error) {
	if tokenString == "" {
		return nil, errors.New(ErrAbsent)
	}

	claims := userStdClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return Secret, nil
	})
	if err != nil {
		return nil, errors.New(ErrInvalid)
	}
	if !token.Valid {
		return nil, errors.New(ErrInvalid)
	}
	return &claims.JwtAdminDto, nil
}

// 返回id
func GetAdminId(c *gin.Context) (uint, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return 0, errors.New("can't get user id")
	}
	admin, ok := u.(*entity.JwtAdminDto)
	if ok {
		return admin.ID, nil
	}
	return 0, errors.New("can't convert to id struct")
}

// 返回用户名
func GetAdminName(c *gin.Context) (string, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return string(string(0)), errors.New("can't get user name")
	}
	admin, ok := u.(*entity.JwtAdminDto)
	if ok {
		return admin.Username, nil
	}
	return string(string(0)), errors.New("can't convert to api name")
}

// 返回admin信息
func GetAdmin(c *gin.Context) (*entity.JwtAdminDto, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return nil, errors.New("can't get api")
	}
	admin, ok := u.(*entity.JwtAdminDto)
	if ok {
		return admin, nil
	}
	return nil, errors.New("can't convert to api struct")
}
