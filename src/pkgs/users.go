package pkgs

import (
	"net/http"

	"github.com/MattiasHenders/palette-town-api/src/db"
	errors "github.com/MattiasHenders/palette-town-api/src/internal/errors"
	"github.com/MattiasHenders/palette-town-api/src/internal/utils"
	"github.com/MattiasHenders/palette-town-api/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	UserTypeGuest      = "guest"
	UserTypeRegistered = "user"
	UserTypeAdmin      = "admin"
)

func GetUserByID(id string) (*models.User, *errors.HTTPError) {

	user, findErr := db.GetUserByID(id)
	if findErr != nil {
		return nil, findErr
	}

	return user, nil
}

func GetUserByEmail(email string) (*models.User, *errors.HTTPError) {

	user, findErr := db.GetUserByEmail(email)
	if findErr != nil {
		return nil, findErr
	}

	return user, nil
}

func UserLogin(email string, password string) (*models.User, *errors.HTTPError) {

	user, findErr := db.GetUserByEmail(email)
	if findErr != nil {
		return nil, findErr
	}

	expectedPass := utils.HashPassword(password)
	if expectedPass != *user.Password {
		return nil, errors.NewHTTPError(nil, http.StatusUnauthorized, "Incorrect password")
	}

	return user, nil
}

func UserSignup(first *string, last *string, email string, password string) (*models.User, *errors.HTTPError) {

	hash := utils.HashPassword(password)

	insertUser := models.User{
		ID:        primitive.NewObjectID(),
		FirstName: first,
		LastName:  last,
		Email:     email,
		Password:  &hash,
		UserType:  UserTypeRegistered,
	}

	user, insertErr := db.CreateNewUser(insertUser)
	if insertErr != nil {
		return nil, insertErr
	}

	return user, nil
}

func UserLogout() *errors.HTTPError {

	return nil
}
