package handlers

import (
	"encoding/json"
	"net/http"

	errors "github.com/MattiasHenders/palette-town-api/src/internal/errors"
	h "github.com/MattiasHenders/palette-town-api/src/internal/server_helpers"
)

func GetRandomColorPaletteHandler() func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	return func(w http.ResponseWriter, r *http.Request) *errors.HTTPError {

		externalCampaignId := h.GetFormParam(r, "externalCampaignId")
		if externalCampaignId == nil {
			return errors.NewHTTPError(nil, http.StatusBadRequest, "Missing externalCampaignId")
		}

		httpErr := errors.NewHTTPError(nil, 500, "")

		if httpErr != nil {
			return httpErr
		}

		json.NewEncoder(w).Encode("success")
		return nil
	}
}
