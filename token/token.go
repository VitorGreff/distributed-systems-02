package tokenPkg

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("chaveAbsurdamenteSecreta")

func GenerateToken(userID uint64) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	// claim for correct auth
	claims["userID"] = userID
	// token expires in 6 hours
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("erro ao assinar o token")
	}

	return signedToken, nil
}

func ValidateToken(tokenString string, userID uint64) error {
	token, err := jwt.Parse(tokenString, returnKey)
	if err != nil {
		fmt.Println(err)
		return errors.New("erro ao dar parse no token")
	}

	tokenID, err := ExtrairUsuarioID(tokenString)
	if err != nil {
		return err
	}

	if tokenID != userID {
		return errors.New("token não pertence ao usuário da requisição")
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

func ExtrairUsuarioID(tokenString string) (uint64, error) {
	token, erro := jwt.Parse(tokenString, returnKey)
	if erro != nil {
		return 0, erro
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", claims["userID"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return usuarioID, nil
	}

	return 0, errors.New("token inválido")
}

func returnKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("metodo de assinatura inesperado! %v", token.Header["alg"])
	}
	return jwtSecret, nil
}
