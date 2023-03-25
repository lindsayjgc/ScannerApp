package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateResponse(message string) map[string]string {
	res := make(map[string]string)
	res["message"] = message
	return res
}

// Return tokenString, expirationTime, error, and error statu scode
func CreateCookie(w http.ResponseWriter, credentials Credentials) (error, int) {
	expirationTime := time.Now().Add(time.Hour * 24) // JWT lasts 1 day
	claims := &Claims{
		Email: credentials.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return errors.New("error creating jwt"), http.StatusInternalServerError
	}

	http.SetCookie(w, 
		&http.Cookie{
			Name: "token",
			Value: tokenString,
			Path: "/",
			Expires: expirationTime,
			SameSite: http.SameSiteLaxMode,
			// Secure: true,
			HttpOnly: true,
		},
	)

	return nil, http.StatusOK
}

func CheckCookie(w http.ResponseWriter, r *http.Request) (*Claims, error, int) {
	// Declare claims so that it can be returned as an empty object if there is an error
	claims := &Claims{}

	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return claims, errors.New("no user logged in"), http.StatusBadRequest
		}
		return claims, errors.New("other cookie-related error"), http.StatusBadRequest
	}

	tokenStr := cookie.Value
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	// This error block should never hit, the HTTP cookie is set to expire at
	// the same time as the JWT, so the JWT parse should not return an
	// error for the JWT being expired
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return claims, errors.New("error parsing JWT"), http.StatusInternalServerError
		}

		return claims, errors.New("other JWT-related error"), http.StatusInternalServerError
	}

	if !tkn.Valid {
		// Since JWT is now invalid, delete the cookie
		DeleteCookie(w)

		return claims, errors.New("other JWT-related error"), http.StatusInternalServerError
	}

	return claims, nil, http.StatusOK
}

func DeleteCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // Signifies delete this cookie now
		HttpOnly: true,
	})
}