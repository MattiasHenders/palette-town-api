package pkgs

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	errors "github.com/MattiasHenders/palette-town-api/src/internal/errors"
)

func GetRandomColorPalette(numOfColours int) (string, *errors.HTTPError) {

	client := &http.Client{}
	var data = strings.NewReader(`{"model":"default"}`)
	req, err := http.NewRequest("POST", "http://colormind.io/api/", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

	return "done", nil
}
