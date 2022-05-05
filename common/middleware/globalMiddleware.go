package middleware

import (
	"fmt"
	"net/http"
)

type GlobalMiddleware struct {
	secret string
}

func NewGlobalMiddleware(secret string) *GlobalMiddleware {
	return &GlobalMiddleware{
		secret: secret,
	}
}

func (m *GlobalMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Global middleware begin", m.secret)
		t := "hello " + m.secret
		next(w, r)

		fmt.Println("Global middleware end", t)
	}
}

// CommonJwtAuthMiddleware : with jwt on the verification, no jwt on the verification
/*
type CommonJwtAuthMiddleware struct {
	secret string
}

func NewCommonJwtAuthMiddleware(secret string) *CommonJwtAuthMiddleware {
	return &CommonJwtAuthMiddleware{
		secret: secret,
	}
}

func (m *CommonJwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("Authorization")) > 0 {
			//has jwt Authorization
			authHandler := handler.Authorize(m.secret)
			authHandler(next).ServeHTTP(w, r)
			return
		} else {
			//no jwt Authorization
			next(w, r)
		}
	}
}
*/
