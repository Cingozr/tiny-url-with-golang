package http

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

var (
	ContentType = []string{"Content-Type", "application/json; charset=UTF-8"}
	SECRET      = []byte("cingoz_recai")
)

func validateToken(accessToken string) bool {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there has been an error")
		}
		return SECRET, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}

func JwtAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("jwt authentication hit")
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			w.Header().Set(ContentType[0], ContentType[1])
			SendErrorResponse(w, "not authorized", errors.New("not authorized"))
			return
		}

		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			w.Header().Set(ContentType[0], ContentType[1])
			SendErrorResponse(w, "not authorized", errors.New("not authorized"))
		}

		if validateToken(authHeaderParts[1]) {
			original(w, r)
		} else {
			w.Header().Set(ContentType[0], ContentType[1])
			SendErrorResponse(w, "not authorized", errors.New("not authorized"))
		}
	}
}

func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString(SECRET)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}

