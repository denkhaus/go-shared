package format

import "github.com/denkhaus/go-shared/colors"

func ColoredFloatPosNeg(format string, v float64, args ...any) string {
	if v > 0 {
		return colors.FgGreen(format, append([]interface{}{v}, args...)...)
	} else if v < 0 {
		return colors.FgRed(format, append([]interface{}{v}, args...)...)
	}

	return colors.FgWhite(format, append([]interface{}{v}, args...)...)
}
