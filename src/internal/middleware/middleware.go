package middleware

import (
	"context"
	"net/http"

	"github.com/MattiasHenders/palette-town-api/src/internal/errors"
	"github.com/MattiasHenders/palette-town-api/src/models"

	"github.com/MattiasHenders/palette-town-api/src/internal/server_helpers"
	userPkgs "github.com/MattiasHenders/palette-town-api/src/pkgs"
	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

func Verifier(publicKeySet *jwk.Set) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			token, err := verifyRequest(publicKeySet, r, jwtauth.TokenFromHeader, jwtauth.TokenFromCookie)
			ctx = jwtauth.NewContext(ctx, token, err)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func HydrateAuthUser() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
			_, claims, err := jwtauth.FromContext(r.Context())
			if err != nil {
				return errors.NewHTTPError(err, http.StatusUnauthorized, "Failed to get claims for auth user")
			}

			id, ok := claims["id"].(string)
			if !ok {
				return errors.NewHTTPError(nil, http.StatusUnauthorized, "Claims did not include username")
			}

			user, httpErr := userPkgs.GetUserByID(id)
			if httpErr != nil {
				return httpErr
			}

			ctx := context.WithValue(r.Context(), userCtxKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
			return nil
		}

		return http.HandlerFunc(server_helpers.Handler(fn))
	}
}

func VerifyAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(server_helpers.Handler(func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
		authUser := r.Context().Value(userCtxKey).(*models.User)

		if authUser.UserType != userPkgs.UserTypeAdmin {
			return errors.NewHTTPError(nil, http.StatusUnauthorized, "Auth user must be an admin")
		}

		next.ServeHTTP(w, r)
		return nil
	}))
}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(server_helpers.Handler(func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil {
			return errors.NewHTTPError(err, http.StatusUnauthorized, jwtauth.ErrorReason(err).Error())
		}

		if token == nil || jwt.Validate(token) != nil {
			return errors.NewHTTPError(nil, http.StatusUnauthorized, "Unauthorized")
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
		return nil
	}))
}

func verifyRequest(publicKeySet *jwk.Set, r *http.Request, findTokenFns ...func(r *http.Request) string) (jwt.Token, error) {
	var tokenString string

	for _, fn := range findTokenFns {
		tokenString = fn(r)
		if tokenString != "" {
			break
		}
	}
	if tokenString == "" {
		return nil, jwtauth.ErrNoTokenFound
	}

	return verifyToken(tokenString, publicKeySet)
}

func verifyToken(tokenString string, publicKeySet *jwk.Set) (jwt.Token, error) {
	token, err := jwt.Parse([]byte(tokenString), jwt.WithKeySet(*publicKeySet))
	if err != nil {
		return token, err
	}

	if err := jwt.Validate(token); err != nil {
		return token, err
	}

	return token, nil
}
