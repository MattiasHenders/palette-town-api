package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/MattiasHenders/palette-town-api/config"
	"github.com/MattiasHenders/palette-town-api/src/db"
	"github.com/MattiasHenders/palette-town-api/src/models"
	"github.com/golang-jwt/jwt"
	"github.com/twinj/uuid"
)

func FetchAuth(authD *models.AccessDetails) (*models.User, error) {
	user, err := db.GetUserByEmail(authD.UserEmail)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateToken(email string) (*models.TokenDetails, error) {

	// Get the key
	access := config.GetConfig().JWT.Access
	var err error
	td := &models.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Hour * 2).Unix()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.AccessUuid = uuid.NewV4().String()
	td.RefreshUuid = uuid.NewV4().String()

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["email"] = email
	atClaims["exp"] = td.AtExpires
	atClaims["access_uuid"] = td.AccessUuid
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(access))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	refresh := config.GetConfig().JWT.Refresh
	rtClaims := jwt.MapClaims{}
	rtClaims["email"] = email
	rtClaims["exp"] = td.RtExpires
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(refresh))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	access := config.GetConfig().JWT.Access
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(access), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(r *http.Request) (*models.AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, fmt.Errorf("token is unauthorized - id")
		}
		email, ok := claims["email"].(string)
		if !ok {
			return nil, fmt.Errorf("token is unauthorized - email")
		}
		return &models.AccessDetails{
			AccessUuid: accessUuid,
			UserEmail:  email,
		}, nil
	}
	return nil, err
}
