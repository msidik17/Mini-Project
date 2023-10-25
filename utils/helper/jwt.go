package helper

import (
	modelsresponse "Mini-Project/models/models-response"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateAdminToken(adminLoginResponse *modelsresponse.AdminLoginResponse, id uint) (string, error) {
	expireTime := time.Now().Add(time.Hour * 24 * 7).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["name"] = adminLoginResponse.Name
	claims["email"] = adminLoginResponse.Email
	claims["exp"] = expireTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}
