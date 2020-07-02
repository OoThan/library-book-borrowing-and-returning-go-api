package auth

import (
	"encoding/json"
	"fmt"
	"github.com/OoThan/library-book-borrowing-and-returning-go-api/api/models"
	"github.com/OoThan/library-book-borrowing-and-returning-go-api/api/responses"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func GenerateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["password"] = user.Password
	claims["exp"] =time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func TokenValid(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected Signing Method: %v ", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractUserTokenID(r *http.Request) (uint64, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected Signed Method: %v ", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 64)
		if err != nil {
			return 0, nil
		}
		return uint64(uid), nil
	}
	return 0, nil
}

func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(b))
}

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr := r.Header.Get("authorization")
			if tokenStr != "" {
				responses.ResponseWithJSON(w, http.StatusUnauthorized, responses.Response{Code: http.StatusUnauthorized, Message: "Not Authorized"})
			} else {
				token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						responses.ResponseWithJSON(w, http.StatusUnauthorized, responses.Response{Code: http.StatusUnauthorized, Message: "Not Authorized"})
						return nil, fmt.Errorf("Not Authorized ")
					}
					return []byte(os.Getenv("API_SECRET")), nil
				})
				if !token.Valid {
					responses.ResponseWithJSON(w, http.StatusUnauthorized, responses.Response{Code: http.StatusUnauthorized, Message: "Not Authorized"})
				} else {
					next.ServeHTTP(w, r)
				}
			}
	})
}
