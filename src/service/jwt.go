package service

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(id int, email string, role string) (string, error){
	claims := jwt.MapClaims{
		"id": id,
		"email": email,
		"role": role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))
	if err != nil{
		return tokenStr, err
	}

	return tokenStr, err
}



func ValidateToken(tokenstring string)(*jwt.Token, error){
	return  jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error){
		if _, isvalid:= token.Method.(*jwt.SigningMethodHMAC); !isvalid{
			return nil, fmt.Errorf("Invalid token. %v", token.Header["alg"])
		}

	return []byte(os.Getenv(os.Getenv("APP_SECRET"))), nil
})
}