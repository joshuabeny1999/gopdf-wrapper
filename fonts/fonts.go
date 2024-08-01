// Package fonts contains the embedded fonts and utility functions.
package fonts

import (
	_ "embed"
)

//go:embed Lato-Heavy.ttf
var latoHeavy []byte

//go:embed Lato-Italic.ttf
var latoItalic []byte

//go:embed Lato-Regular.ttf
var latoRegular []byte

//go:embed LiberationSans-Bold.ttf
var liberationSansBold []byte

//go:embed LiberationSans-Italic.ttf
var liberationSansItalic []byte

//go:embed LiberationSans-Regular.ttf
var liberationSansRegular []byte

type FontFamily struct {
	Name   string
	Bold   []byte
	Italic []byte
	Normal []byte
}

func NewLatoFamily() (*FontFamily, error) {
	return &FontFamily{
		Name:   "lato",
		Bold:   latoHeavy,
		Italic: latoItalic,
		Normal: latoRegular,
	}, nil
}

func NewLiberationSansFamily() (*FontFamily, error) {
	return &FontFamily{
		Name:   "liberation-sans",
		Bold:   liberationSansBold,
		Italic: liberationSansItalic,
		Normal: liberationSansRegular,
	}, nil
}
