package colors

import "github.com/fatih/color"

var (
	FgGreenUnderlined  = color.New(color.FgGreen, color.Underline).SprintfFunc()
	FgYellow           = color.New(color.FgYellow).SprintfFunc()
	FgYellowUnderlined = color.New(color.FgYellow, color.Underline).SprintfFunc()
	FgWhite            = color.New(color.FgWhite).SprintfFunc()
	FgWhiteBold        = color.New(color.FgWhite, color.Bold).SprintfFunc()
	FgGreen            = color.New(color.FgGreen).SprintfFunc()
	FgGreenBold        = color.New(color.FgGreen, color.Bold).SprintfFunc()
	FgMagenta          = color.New(color.FgMagenta).SprintfFunc()
	FgMagentaBold      = color.New(color.FgMagenta, color.Bold).SprintfFunc()
	FgBlue             = color.New(color.FgBlue).SprintfFunc()
	FgBlueBold         = color.New(color.FgBlue, color.Bold).SprintfFunc()
	FgRed              = color.New(color.FgRed).SprintfFunc()
	FgRedBold          = color.New(color.FgRed, color.Bold).SprintfFunc()
	BgYellow           = color.New(color.BgYellow).SprintfFunc()
	BgWhite            = color.New(color.BgWhite).SprintfFunc()
	BgGreen            = color.New(color.BgGreen).SprintfFunc()
	BgBlue             = color.New(color.BgBlue).SprintfFunc()
	BgHiBlue           = color.New(color.BgHiBlue).SprintfFunc()
	BgMagenta          = color.New(color.BgMagenta).SprintfFunc()
	BgRed              = color.New(color.BgRed).SprintfFunc()
	BgRedFGBlack       = color.New(color.BgRed, color.FgBlack).SprintfFunc()
	BgGreenFGBlack     = color.New(color.BgGreen, color.FgBlack).SprintfFunc()
)
