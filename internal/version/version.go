package version

import "github.com/common-nighthawk/go-figure"

// Version ...
var Row1 = "FSO_CLI"

// var Row2 = "developers"
var Row3 = "v0.0.2"

func GetVersion(font string) string {
	str := "\n" + figure.NewFigure(Row1, font, false).String()

	// str += figure.NewFigure(Row2, font, false).String()

	str += figure.NewFigure(Row3, font, false).String()
	str += "\n" + Row3
	return str
}
