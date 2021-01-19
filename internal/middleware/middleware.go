package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

const (
	// Store this as environment variable in the future
	Secret = "the-todanni-secret"
	// This isn't used anywhere... yet
	Issuer = "todanni-user-service"
	// Encryption cost
	Cost = 14
)

type ToDanniClaims struct {
	jwt.StandardClaims
	UserInfo UserClaims `json:"user_info"`
}

type UserClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
}

type HTTPReqInfo struct {
	// GET etc.
	method    string
	uri       string
	referer   string
	userAgent string
	body      string
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqInfo := &HTTPReqInfo{
			method:    r.Method,
			uri:       r.URL.String(),
			referer:   r.Header.Get("Referer"),
			userAgent: r.Header.Get("User-Agent"),
		}
		log.Println("Testing logging")
		log.Println(reqInfo)
		next.ServeHTTP(w, r)
	})
}

// Middleware function, which will be called for each request
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Header.Get("Authorization")
		splitToken := strings.Split(tkn, "Bearer ")

		if len(splitToken) < 2 {
			http.Error(w, "Missing Auth Token", http.StatusUnauthorized)
			return
		}

		tkn = splitToken[1]
		if ctx, ok := IsValid(tkn, r.Context()); !ok {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		} else {
			// We found the token in our map
			log.Printf("Authenticated token %s", tkn)
			// TODO: Instead of the whole token, just send the user ID
			req := r.WithContext(ctx)
			next.ServeHTTP(w, req)
		}
	})
}

func IsValid(tokenString string, ctx context.Context) (context.Context, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &ToDanniClaims{}, keyFunc)
	if err != nil {
		log.Error(err)
		return ctx, false
	}

	if clms, ok := token.Claims.(*ToDanniClaims); ok && token.Valid {
		return context.WithValue(ctx, "user_id", clms.UserInfo.UserID), true
	}

	log.Error(err)
	return ctx, false
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	// TODO: later the "kid" can be used to check the version of the key used to sign the JWT
	// This will come in handy when key rotation is implemented.
	return []byte(Secret), nil
}
