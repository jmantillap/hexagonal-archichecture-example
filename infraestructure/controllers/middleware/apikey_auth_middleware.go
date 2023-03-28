package middleware

import "net/http"

type apiKeyAuthMiddleware struct {
	next http.Handler
}


func NewApiKeyAuthMiddleware(next http.Handler) *apiKeyAuthMiddleware {
	return &apiKeyAuthMiddleware{next}
}


func (m *apiKeyAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	validApiKey := "my-secret-api-key"

	if apiKey == "" || apiKey != validApiKey {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	m.next.ServeHTTP(w, r)
}