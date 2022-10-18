package models

type ColourPalette struct {
	Colours []string `json:"colours"`
}

type RGBColour struct {
	Red   int
	Green int
	Blue  int
}
