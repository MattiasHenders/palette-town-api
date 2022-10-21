package pkgs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/MattiasHenders/palette-town-api/config"

	errors "github.com/MattiasHenders/palette-town-api/src/internal/errors"
	"github.com/MattiasHenders/palette-town-api/src/internal/server_helpers"
	"github.com/MattiasHenders/palette-town-api/src/models"
)

const (
	MAX_AMOUNT_COLOURS = 5
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

func ConvertHexIntoRGBCode(hex string) (models.RGBColour, error) {
	var rgb models.RGBColour
	fmt.Println(hex)
	values, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return models.RGBColour{}, err
	}

	rgb = models.RGBColour{
		Red:   int(uint8(values >> 16)),
		Green: int(uint8((values >> 8) & 0xFF)),
		Blue:  int(uint8(values & 0xFF)),
	}
	return rgb, nil
}

func CreateColourPalette(hexCodeArray []string) *models.ColourPalette {
	return &models.ColourPalette{Colours: hexCodeArray}
}

func GetColourPalettePromptData(rawColours string) (string, error) {

	colours := strings.Split(strings.ReplaceAll(rawColours, " ", ""), ",")

	//Check we have MAX_AMOUNT_COLOURS or less colours
	if len(colours) > MAX_AMOUNT_COLOURS {
		return "", fmt.Errorf("max amount of colours allowed is 5, you provided %d", len(colours))
	}

	// Validate each colour
	for _, colour := range colours {
		isValid := ValidateHexCode(colour)
		if !isValid {
			return "", fmt.Errorf("%s is not a valid hex code", colour)
		}
	}

	// Convert each hex into RGB
	var parsedColours []models.RGBColour
	for _, colour := range colours {
		parsedColour, parseErr := ConvertHexIntoRGBCode(strings.ReplaceAll(colour, "#", ""))
		if parseErr != nil {
			return "", fmt.Errorf("error parsing %s into rgb: %x", colour, parseErr)
		}
		parsedColours = append(parsedColours, parsedColour)
	}

	// Create the data to go into colorminds input
	var colorMindsInput string
	inputCount := 0
	for _, rgb := range parsedColours {
		colorMindsInput += fmt.Sprintf("[%d,%d,%d]", rgb.Red, rgb.Green, rgb.Blue)
		inputCount++

		// Check if its not our last input
		if inputCount != MAX_AMOUNT_COLOURS {
			colorMindsInput += ","
		}
	}

	// Dont go over max amount of colours
	for i := inputCount; i < 5; i++ {
		colorMindsInput += `"N"`
		if i != MAX_AMOUNT_COLOURS-1 {
			colorMindsInput += ","
		}
	}

	//Put the data in proper form
	data := fmt.Sprintf(`{"input":[%s],"model":"default"}`, colorMindsInput)

	return data, nil
}

func ValidateHexCode(hexCode string) bool {
	r := regexp.MustCompile(`^#(([0-9a-fA-F]{2}){3}|([0-9a-fA-F]){3})$`)
	match := r.MatchString(hexCode)
	return match
}
