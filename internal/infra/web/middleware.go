package web

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/dgrijalva/jwt-go"
)

func (app *Application) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			app.errorJson(w,errors.New("you must be logged in to access"), http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			app.errorJson(w,errors.New("no authorization header received"))
			return
		}

		token := headerParts[1]
		//marketplace-go nome do realm
		provider,err := oidc.NewProvider(r.Context(), os.Getenv("KEYCLOCK_URL"))
		if err != nil {
			app.errorJson(w,errors.New("error to connect to identity provider"), http.StatusInternalServerError)
			return
		}
		
		verifier := provider.Verifier(&oidc.Config{ClientID: "marketplace"})
		//verifier := provider.Verifier(&oidc.Config{SkipClientIDCheck: true})
		_,err = verifier.Verify(r.Context(),token)
		if err != nil {
			app.errorJson(w,errors.New("invalid token"), http.StatusUnauthorized)
			return
		}

		//pegar o email do token
		tokenPayload,_ := jwt.Parse(token, nil)
		claims := tokenPayload.Claims.(jwt.MapClaims)
		email := claims["email"]

		//inserir o novo valor no contexto da request
		ctx := context.WithValue(r.Context(), "email", email)

		//inseri o novo contexto nas requisicoes
		next.ServeHTTP(w,r.WithContext(ctx))
	})
}