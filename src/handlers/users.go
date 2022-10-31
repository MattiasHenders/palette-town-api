package handlers

import (
	"encoding/json"
	"net/http"

	errors "github.com/MattiasHenders/palette-town-api/src/internal/errors"
	"github.com/MattiasHenders/palette-town-api/src/internal/server_helpers"
	"github.com/MattiasHenders/palette-town-api/src/pkgs"
)

func GetUserByIDHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		id := server_helpers.GetFormParam(r, "id")
		if id == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing id")
		}

		return nil
	}
}

func GetUserByEmailHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		email := server_helpers.GetFormParam(r, "email")
		if email == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing email")
		}

		user, userErr := pkgs.GetUserByEmail(*email)
		if userErr != nil {
			return userErr
		}

		json.NewEncoder(w).Encode(user)
		return nil
	}
}

func PostUserLoginHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		email := server_helpers.GetFormParam(r, "email")
		if email == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing email")
		}

		password := server_helpers.GetFormParam(r, "password")
		if password == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing password")
		}

		user, userErr := pkgs.UserLogin(*email, *password)
		if userErr != nil {
			return userErr
		}

		json.NewEncoder(w).Encode(user)
		return nil
	}
}

func PostUserSignupHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		email := server_helpers.GetFormParam(r, "email")
		if email == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing email")
		}

		password := server_helpers.GetFormParam(r, "password")
		if password == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing password")
		}

		// Optional params
		firstName := server_helpers.GetFormParam(r, "firstName")
		lastName := server_helpers.GetFormParam(r, "lastName")

		user, userErr := pkgs.UserSignup(firstName, lastName, *email, *password)
		if userErr != nil {
			return userErr
		}

		json.NewEncoder(w).Encode(user)
		return nil
	}
}

func PostUserLogoutHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		return nil
	}
}
