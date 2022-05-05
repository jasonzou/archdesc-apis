package middleware

import (
	"fmt"
	"net/http"
)

type TaxonomyMiddleware struct {
}

func NewTaxonomyMiddleware() *TaxonomyMiddleware {
	return &TaxonomyMiddleware{}
}

func (m *TaxonomyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		fmt.Println("taxonomy Middleware in")
		// Passthrough to next handler if need
		next(w, r)

		fmt.Println(("taxonomy middleware out"))
	}
}
