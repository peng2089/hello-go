package middleware

import (
	"fmt"
	"hello-go/internal/biz"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func TokenMiddleware(uu *biz.UserUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. 获取token
		tokenStr := ctx.GetHeader("Authorization")
		if len(tokenStr) == 0 || !strings.HasPrefix(tokenStr, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token不存在",
			})
			ctx.Abort()
			return
		}
		tokenStr = tokenStr[len("Bearer "):]
		// fmt.Printf("tokenStr: %s\n", tokenStr)
		// 2. 验证token
		key := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAmAoicc41x0qTYHNGl+9y
dzF6VNzU82F5FW5H1PJDGL7CjKkoSg/mjJntG3F/H4y6oec2wzBwb7HJcSFhb1n+
8PvlkwGVlUnn5hK88KHrVOfhEJsKZOKp9BaQtRyjjme03cEE8iQgZ69kF6UODWkQ
WXvEDCd9NWh4sQIqpo/nx5pljB7nuZe5GV3dSU1HTQ1sEzAkTCPu7lnrtGzSe8Ea
4BPFy62JZl4MSM9iOmHsjmDe+KIlL1q6lQISEqSSEJ/Yx8GEXDibJ7xPGB4YOvPc
Koxrmx8XF9lSj8DFbfTQEE4svYTinW+SyeuFZQMP05g9hOWwqq/31rzEADrNzOgQ
BQIDAQAB
-----END PUBLIC KEY-----`
		pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(key))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "token格式不正确",
			})
			ctx.Abort()
			return
		}
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return pubKey, nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Printf("claims: %+v\n", claims)
			// TODO: 根据token中解析的claims信息查询数据库
			if ok, _ := uu.ValidToken(ctx, &biz.User{ID: uint(claims["uid"].(float64))}); !ok {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "token验证失败",
				})
				ctx.Abort()
				return
			}
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "token验证失败",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
