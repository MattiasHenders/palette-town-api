package pkgs

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/cascax/colorthief-go"
	"github.com/gocolly/colly"

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

func GetWordPromptColourPalette(word string) (*models.ColourPalette, *errors.HTTPError) {

	// Get URL to Colorminds ML
	colormindURL := config.GetConfig().API.ColorMindURL

	//Search for images
	images := getSearch(word)
	if images.Count == 0 {
		return nil, errors.NewHTTPError(nil, http.StatusInternalServerError, "found no images from search term")
	}

	// Randomly get an image found
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	index := r1.Intn(images.Count)

	// Create file name
	raw := makeTimestamp()
	path := filepath.Join("bin", "images", raw)

	// Get Image from Google based on word
	downloadErr := downloadFile(images.Data[index], path)
	if downloadErr != nil {
		return nil, errors.NewHTTPError(downloadErr, http.StatusInternalServerError, downloadErr.Error())
	}

	// Get colours from image downloaded
	rawAlphaColours, parseErr := colorthief.GetPaletteFromFile(path, 5)
	if parseErr != nil {
		return nil, errors.NewHTTPError(parseErr, http.StatusInternalServerError, parseErr.Error())
	}

	// Parse prompt data
	colorsPromptData, inputErr := GetColourPaletteWordPromptData(rawAlphaColours)
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

func GetColourPaletteWordPromptData(colours []color.Color) (string, error) {

	// Convert each alpha color into RGB
	str := fmt.Sprint(colours)
	str = strings.ReplaceAll(str, "[{", "")
	str = strings.ReplaceAll(str, "}]", "")
	strArr := strings.Split(str, "} {")

	// Create the data to go into colorminds input
	var colorMindsInput string
	inputCount := 0
	for _, rgba := range strArr {
		rgb := strings.Split(rgba, " ")
		colorMindsInput += fmt.Sprintf("[%s,%s,%s]", rgb[0], rgb[1], rgb[2])
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

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("failed to download file properly")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func makeTimestamp() string {
	str := fmt.Sprintf("%x", time.Now().UnixNano()/int64(time.Millisecond))
	return str
}

func getSearch(searchQuery string) Images {

	searchString := strings.Replace(searchQuery, " ", "-", -1)

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:42.0) Gecko/20100101 Firefox/42.0"
	c.AllowURLRevisit = true
	c.DisableCookies()
	array := []string{}

	// Find and visit all links
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		src := e.Attr("src")
		if src != "" {
			array = append(array, e.Attr("src"))
		}
	})

	// Search query
	// pexelsQuery := strings.Replace(searchString, "-", "%20", -1)
	stocSnapQuery := strings.Replace(searchString, "-", "+", -1)

	// c.Visit("https://www.flickr.com/search/?text=" + pexelsQuery)
	c.Visit("http://www.google.com/images?q=" + stocSnapQuery)
	return Images{
		Count: len(array),
		Data:  array}
}

type Images struct {
	Count int      `json:"counts"`
	Data  []string `json:"data"`
}

func StringFromUInt32(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
