package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	errors "github.com/MattiasHenders/palette-town-api/src/internal/errors"
	h "github.com/MattiasHenders/palette-town-api/src/internal/server_helpers"
	p "github.com/MattiasHenders/palette-town-api/src/pkgs"
)

func GetRandomColorPaletteHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		// Get request parameters
		rawNumberOfColours := h.GetQueryParam(r, "numOfColours")
		numOfColours := 5
		if rawNumberOfColours != nil {
			parsedNums, err := strconv.Atoi(*rawNumberOfColours)
			if err != nil {
				return errors.NewHTTPError(err, http.StatusBadRequest, "Param numOfColours is not an integer")
			}
			if numOfColours > 5 {
				return errors.NewHTTPError(err, http.StatusBadRequest, "Param numOfColours can not be greater than 5")
			}
			numOfColours = parsedNums
		}

		// Call the function
		httpResp, httpErr := p.GetRandomColorPalette(numOfColours)
		if httpErr != nil {
			return httpErr
		}

		// If call was a success
		json.NewEncoder(w).Encode(httpResp)
		return nil
	}
}
