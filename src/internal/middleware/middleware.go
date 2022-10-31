package middleware

import (
	"net/http"

	"github.com/MattiasHenders/palette-town-api/src/internal/auth"
	"github.com/MattiasHenders/palette-town-api/src/internal/errors"
	"github.com/MattiasHenders/palette-town-api/src/internal/server_helpers"
	"github.com/MattiasHenders/palette-town-api/src/pkgs"
)

const (
	rapidAPIHost = "rapidapi.com"
)

func AuthenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(server_helpers.Handler(func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		// Rapid API Auth is already handled
		if r.Host == rapidAPIHost {
			// Token is authenticated, pass it through
			next.ServeHTTP(w, r)
			return nil
		}

		tokenAuth, err := auth.ExtractTokenMetadata(r)
		if err != nil {
			return errors.NewHTTPError(nil, http.StatusUnauthorized, "Token is invalid")
		}
		user, err := auth.FetchAuth(tokenAuth)
		if err != nil || user.UserType != pkgs.UserTypeRegistered {
			return errors.NewHTTPError(nil, http.StatusUnauthorized, "Token is unauthorized")
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
		return nil
	}))
}

func AuthenticateAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(server_helpers.Handler(func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		tokenAuth, err := auth.ExtractTokenMetadata(r)
		if err != nil {
			return errors.NewHTTPError(nil, http.StatusUnauthorized, "Token is invalid")
		}
		user, err := auth.FetchAuth(tokenAuth)
		if err != nil || user.UserType != pkgs.UserTypeAdmin {
			return errors.NewHTTPError(nil, http.StatusUnauthorized, "Token is unauthorized")
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
		return nil
	}))
}
