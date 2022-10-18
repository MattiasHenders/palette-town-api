package server_helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"

	errors "github.com/MattiasHenders/palette-town-api/src/internal/errors"
)

func Handler(h func(w http.ResponseWriter, r *http.Request) *errors.HTTPError) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		httpError := h(w, r)
		if httpError == nil {
			return
		}

		body, err := httpError.ResponseBody()
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(httpError.Status)
		w.Write(body)
	})
}

func GetURLParam(r *http.Request, field string) *string {
	param := chi.URLParam(r, field)
	if param == "" {
		return nil
	}

	return &param
}

func GetFormParam(r *http.Request, field string) *string {
	param := r.Form.Get(field)
	if strings.TrimSpace(param) == "" {
		return nil
	}

	return &param
}

func GetFormParamInt(r *http.Request, field string) *int {
	param := r.Form.Get(field)
	if param == "" {
		return nil
	}

	paramInt, err := strconv.Atoi(param)
	if err != nil {
		return nil
	}

	return &paramInt
}

func GetFormParamBool(r *http.Request, field string) *bool {
	param := r.Form.Get(field)
	if param == "" {
		return nil
	}

	paramBool, err := strconv.ParseBool(param)
	if err != nil {
		return nil
	}

	return &paramBool
}

func GetQueryParam(r *http.Request, field string) *string {
	param := r.URL.Query().Get(field)
	if param == "" {
		return nil
	}

	return &param
}

func MakeInternalRequest(method string, url string, rawData string) ([]byte, *errors.HTTPError) {

	// Build request
	client := &http.Client{}
	var data = strings.NewReader(rawData)
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, errors.NewHTTPError(err, http.StatusInternalServerError, fmt.Sprintf("Error building request to %s", url))
	}

	// Set headers and make request
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.NewHTTPError(err, http.StatusInternalServerError, fmt.Sprintf("Failed to make internal request to %s", url))
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.NewHTTPError(err, http.StatusInternalServerError, fmt.Sprintf("Failed to read response from internal request to %s", url))
	}

	// Return body text as byte array
	return bodyText, nil
}
