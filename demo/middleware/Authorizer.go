package middleware

import (
	"demo/dao"
	"demo/util"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(util.GetConfig("jwtSecret"))

type Claims struct {
	Name string `json:"name"`
	Uid  string `json:"phone"`
	jwt.StandardClaims
}

func GenerateToken(username, Uid string) (string, error) {
	nowTime := time.Now()                     //当前时间
	expireTime := nowTime.Add(24 * time.Hour) //有效时间 一天

	claims := Claims{
		username,
		Uid,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "admin-server",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg": "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		// parseToken 解析token包含的信息
		claims, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
		log.Println(claims)
		// 在redis里查找uid与token是否匹配
		var val string
		if val, err = dao.Rdb.Get("Token:" + claims.Uid).Result(); err != nil {
			log.Fatal(err)
			return
		}
		if val != token {
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"msg":    "token与用户不符",
			})
			c.Abort()
			return
		}

		// 得到claims 就是上面那个结构体
		c.Next()
	}

}
