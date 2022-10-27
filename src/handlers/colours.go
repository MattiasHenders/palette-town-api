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

		// Call the function
		httpResp, httpErr := p.GetRandomColourPalette()
		if httpErr != nil {
			return httpErr
		}

		// If call was a success
		resp := models.ServerResponse{
			Message: "Successfully got random colour palette",
			Code:    http.StatusOK,
			Data:    httpResp,
		}
		json.NewEncoder(w).Encode(resp)
		return nil
	}
}

func GetColourPromptColourPaletteHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		// Get the colours from query
		colours := server_helpers.GetQueryParam(r, "colours")
		if colours == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing colours")
		}

		// Call the function
		httpResp, httpErr := p.GetColourPromptColourPalette(*colours)
		if httpErr != nil {
			return httpErr
		}

		// If call was a success
		resp := models.ServerResponse{
			Message:    "Successfully got colour palette from given colours",
			GivenInput: colours,
			Code:       http.StatusOK,
			Data:       httpResp,
		}
		json.NewEncoder(w).Encode(resp)
		return nil
	}
}

func GetWordPromptColourPaletteHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		// Get the colours from query
		word := server_helpers.GetQueryParam(r, "word")
		if word == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing word")
		}

		// Call the function
		httpResp, httpErr := p.GetWordPromptColourPalette(*word)
		if httpErr != nil {
			return httpErr
		}

		// If call was a success
		resp := models.ServerResponse{
			Message:    "Successfully got colour palette from given word",
			GivenInput: word,
			Code:       http.StatusOK,
			Data:       httpResp,
		}
		json.NewEncoder(w).Encode(resp)
		return nil
	}
}
