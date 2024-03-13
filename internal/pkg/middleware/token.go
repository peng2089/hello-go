package middleware

import (
	"fmt"
	"hello-go/internal/biz"
	"io"
	"net/http"
	"os"
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
		pubKeyFile, err := os.Open("./config/public.key")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "缺失公钥",
			})
			ctx.Abort()
			return
		}
		pubByte, err := io.ReadAll(pubKeyFile)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "读取公钥失败",
			})
			ctx.Abort()
			return
		}
		pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubByte)
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
			// fmt.Printf("claims: %+v\n", claims)
			// TODO: 根据token中解析的claims信息查询数据库
			ok, u, err := uu.ValidToken(ctx, &biz.User{ID: uint(claims["uid"].(float64))})
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "token验证失败",
				})
				ctx.Abort()
				return
			}
			if !ok {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "token验证失败",
				})
				ctx.Abort()
				return
			}
			ctx.Set("user", &biz.User{
				ID:       u.ID,
				Username: u.Username,
				Password: u.Password,
			})
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
