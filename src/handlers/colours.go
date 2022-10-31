package handlers

import (
	"encoding/json"
	"net/http"

	errors "github.com/MattiasHenders/palette-town-api/src/internal/errors"
	"github.com/MattiasHenders/palette-town-api/src/internal/server_helpers"
	"github.com/MattiasHenders/palette-town-api/src/models"
	p "github.com/MattiasHenders/palette-town-api/src/pkgs"
)

func GetRandomColourPaletteHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		httpResp, httpErr := p.GetRandomColourPalette()
		if httpErr != nil {
			return httpErr
		}

		// Generate a coolors link for easy viewing
		coolorsLink := server_helpers.GenerateCoolorsLink(httpResp)

		// If call was a success
		resp := models.ServerResponse{
			Message:     "Successfully got random colour palette",
			Code:        http.StatusOK,
			Data:        httpResp,
			CoolorsLink: coolorsLink,
		}
		json.NewEncoder(w).Encode(resp)
		return nil
	}
}

func GetColourPromptColourPaletteHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		colours := server_helpers.GetQueryParam(r, "colours")
		if colours == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing colours")
		}

		httpResp, httpErr := p.GetColourPromptColourPalette(*colours)
		if httpErr != nil {
			return httpErr
		}

		// Generate a coolors link for easy viewing
		coolorsLink := server_helpers.GenerateCoolorsLink(httpResp)

		// If call was a success
		resp := models.ServerResponse{
			Message:     "Successfully got colour palette from given colours",
			GivenInput:  colours,
			Code:        http.StatusOK,
			Data:        httpResp,
			CoolorsLink: coolorsLink,
		}
		json.NewEncoder(w).Encode(resp)
		return nil
	}
}

func GetWordPromptColourPaletteHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		words := server_helpers.GetQueryParam(r, "words")
		if words == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing words")
		}

		httpResp, httpErr := p.GetWordPromptColourPalette(*words)
		if httpErr != nil {
			return httpErr
		}

		// Generate a coolors link for easy viewing
		coolorsLink := server_helpers.GenerateCoolorsLink(httpResp)

		// If call was a success
		resp := models.ServerResponse{
			Message:     "Successfully got colour palette from given word",
			GivenInput:  words,
			Code:        http.StatusOK,
			Data:        httpResp,
			CoolorsLink: coolorsLink,
		}
		json.NewEncoder(w).Encode(resp)
		return nil
	}
}
