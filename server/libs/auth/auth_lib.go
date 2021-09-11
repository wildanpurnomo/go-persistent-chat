package authLibs

import (
	"context"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	environmentLibs "github.com/wildanpurnomo/go-persistent-chat/server/libs/environment"
)

type ContextValueWrapper struct {
	GinContext *gin.Context
	UserID     int64
}

type AuthTokenClaims struct {
	UserID int64
	jwt.StandardClaims
}

var (
	JwtSecret                      = []byte(os.Getenv("JWT_SECRET"))
	CookieAndTokenMaxAge           = 60 * 60 * 24
	CookieMaxAgeForClearance       = 1
	ContextValueKey                = "ContextValueKey"
	InvalidInitialUserID     int64 = -1
)

func ExtractContextValueWrapper(context context.Context) *ContextValueWrapper {
	return context.Value(ContextValueKey).(*ContextValueWrapper)
}

func CheckWhetherRequestIsUsingValidToken(requestContext context.Context) (bool, int64) {
	userId := ExtractContextValueWrapper(requestContext).UserID
	return userId > InvalidInitialUserID, userId
}

func SetJwtCookie(requestContext context.Context, userId int64) {
	token, _ := GenerateJwtToken(userId)
	contextValueWrapper := ExtractContextValueWrapper(requestContext)
	contextValueWrapper.GinContext.SetCookie("jwt", token, CookieAndTokenMaxAge, "/", "", environmentLibs.IsRunningOnProductionEnvironment(), true)
}

func ClearJwtCookie(requestContext context.Context) {
	contextValueWrapper := ExtractContextValueWrapper(requestContext)
	contextValueWrapper.GinContext.SetCookie("jwt", "", CookieMaxAgeForClearance, "/", "", environmentLibs.IsRunningOnProductionEnvironment(), true)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := InvalidInitialUserID

		jwtCookie, _ := c.Request.Cookie("jwt")
		if jwtCookie != nil {
			userId = ExtractUserIdFromToken(jwtCookie.Value)
		}

		c.Request = c.Request.WithContext(
			context.WithValue(
				c.Request.Context(),
				ContextValueKey,
				&ContextValueWrapper{
					GinContext: c,
					UserID:     userId,
				},
			),
		)

		c.Next()
	}
}

func ExtractUserIdFromToken(tokenString string) int64 {
	claims := &AuthTokenClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return int64(InvalidInitialUserID)
	}

	return claims.UserID
}

func GenerateJwtToken(userId int64) (string, error) {
	claims := &AuthTokenClaims{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	signedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := signedToken.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
