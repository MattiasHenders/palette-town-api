package pkgs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/MattiasHenders/palette-town-api/config"

	errors "github.com/MattiasHenders/palette-town-api/src/internal/errors"
	"github.com/MattiasHenders/palette-town-api/src/internal/server_helpers"
	"github.com/MattiasHenders/palette-town-api/src/models"
)

func GetRandomColourPalette() (*models.ColourPalette, *errors.HTTPError) {

	// Get URL to Colorminds ML
	colormindURL := config.GetConfig().API.ColorMindURL

	// Make internal request to get raw palette
	colourBytes, httpErr := server_helpers.MakeInternalRequest("POST", colormindURL, `{"model":"default"}`)
	if httpErr != nil {
		return nil, httpErr
	}

	// Parse ColorMindsRequest
	colourString := string(colourBytes)
	colorMindsResp := models.ColorMindsResponse{}
	err := json.Unmarshal([]byte(colourString), &colorMindsResp)
	if err != nil {
		return nil, errors.NewHTTPError(err, http.StatusInternalServerError, "Error generating ColorMindsRequest in GetRandomColourPalette")
	}

	// Generate RGB string array from Colorminds string
	rawColourArr := GetRGBCodes(&colorMindsResp)

	//Generate ColourPalette
	hexCodes := ConvertRGBArrayIntoHexCodeArray(rawColourArr)
	palette := CreateColourPalette(hexCodes)

	// Return the found colours
	return palette, nil
}

func GetColourPromptColourPalette(colours string) (*models.ColourPalette, *errors.HTTPError) {

	// Get URL to Colorminds ML
	colormindURL := config.GetConfig().API.ColorMindURL

	// Parse prompt data
	colorsPromptData, inputErr := GetColourPalettePromptData(colours)
	if inputErr != nil {
		return nil, errors.NewHTTPError(inputErr, http.StatusBadRequest, inputErr.Error())
	}

	// Make internal request to get raw palette
	colourBytes, httpErr := server_helpers.MakeInternalRequest("POST", colormindURL, colorsPromptData)
	if httpErr != nil {
		return nil, httpErr
	}

	// Parse ColorMindsRequest
	colourString := string(colourBytes)
	colorMindsResp := models.ColorMindsResponse{}
	err := json.Unmarshal([]byte(colourString), &colorMindsResp)
	if err != nil {
		return nil, errors.NewHTTPError(err, http.StatusInternalServerError, "Error generating ColorMindsRequest in GetRandomColourPalette")
	}

	// Generate RGB string array from Colorminds string
	rawColourArr := GetRGBCodes(&colorMindsResp)

	//Generate ColourPalette
	hexCodes := ConvertRGBArrayIntoHexCodeArray(rawColourArr)
	palette := CreateColourPalette(hexCodes)

	// Return the found colours
	return palette, nil
}

func GetRGBCodes(colourMindsResponse *models.ColorMindsResponse) []models.RGBColour {

	var rgbCodes []models.RGBColour
	for _, colour := range colourMindsResponse.Result {
		rgbCodes = append(rgbCodes, models.RGBColour{
			Red:   colour[0],
			Green: colour[1],
			Blue:  colour[2],
		})
	}
	return rgbCodes
}

func ConvertRGBArrayIntoHexCodeArray(rgbColours []models.RGBColour) []string {
	var hexCodes []string
	for _, colour := range rgbColours {
		hexCodes = append(hexCodes, ConvertRGBIntoHexCode(colour))
	}
	return hexCodes
}

func ConvertRGBIntoHexCode(rgbColour models.RGBColour) string {
	return fmt.Sprintf("#%02x%02x%02x", rgbColour.Red, rgbColour.Green, rgbColour.Blue)
}

func CreateColourPalette(hexCodeArray []string) *models.ColourPalette {
	return &models.ColourPalette{Colours: hexCodeArray}
}

func GetColourPalettePromptData(rawColours string) (string, error) {

	colours := strings.Split(rawColours, ",")

	//Check we have 5 or less colours
	if len(colours) > 5 {
		return "", fmt.Errorf("Max amount of colours allowed is 5, you provided %d", len(colours))
	}

	// Validate each colour
	for _, colour := range colours {
		err := ValidateHexCode(colour)
		if err {
			return "", fmt.Errorf("%s is not a valid hex code.", colour)
		}
	}

	// TODO: Convert each hex into RGB
	data := `{"input":[[44,43,44],[90,83,82],"N","N","N"],"model":"default"}`

	return data, nil
}

func ValidateHexCode(hexCode string) bool {
	match, _ := regexp.MatchString("^#(?:[0-9a-fA-F]{3}){1,2}$", hexCode)
	return match
}
