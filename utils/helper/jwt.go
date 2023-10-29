package helper

import (
	modelsresponse "Mini-Project/models/models-response"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAdminToken(adminLoginResponse *modelsresponse.AdminLoginResponse, id uint) (string, error) {
	expireTime := time.Now().Add(time.Hour * 24 * 7).Unix()
	claims := jwt.MapClaims{
		"id":       id,
		"email":    adminLoginResponse.Email,
		"exp":      expireTime,
		"role":     "admin",
	}
	// claims["authorized"] = true
	// claims["id"] = id
	// claims["email"] = adminLoginResponse.Email
	// claims["password"] = adminLoginResponse.Password
	// claims["exp"] = expireTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}

func GenerateUserToken(UserLoginResponse *modelsresponse.UserLoginResponse, id uint) (string, error) {
	expireTime := time.Now().Add(time.Hour * 24 * 7).Unix()
	claims := jwt.MapClaims{
		"id":       id,
		"email":    UserLoginResponse.Email,
		"exp":      expireTime,
		"role":     "user",
	}
	// claims["authorized"] = true
	// claims["id"] = id
	// claims["name"] = UserLoginResponse.Name
	// claims["email"] = UserLoginResponse.Email
	// claims["exp"] = expireTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}

// func ExtractTokenAdminId(e echo.Context) float64 {
// 	user := e.Get("admin").(*jwt.Token)
// 	if user.Valid {
// 		claims := user.Claims.(jwt.MapClaims)
// 		UserId := claims["id"].(float64)
// 		return UserId
// 	}
// 	return 0
// }

// func ExtractTokenUserId(e echo.Context) float64 {
// 	user := e.Get("user").(*jwt.Token)
// 	if user.Valid {
// 		claims := user.Claims.(jwt.MapClaims)
// 		UserId := claims["id"].(float64)
// 		return UserId
// 	}
// 	return 0
// }
