package shared

import (
	"bytes"
	"fmt"
	"strconv"
	"text/tabwriter"
)

func rightAlignColumnize(value, unit string) string {
	w := new(tabwriter.Writer)
	bs := bytes.NewBuffer([]byte{})
	// Using tabwriter.Debug to output '|' which is the delimited in columnize
	w.Init(bs, 15, 0, 1, ' ', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintf(w, "%s\t%s", value, unit)
	w.Flush()
	return bs.String()
}

// func PrettifyAsset(w interfaces.EOSTAsset) string {
// 	const unit = 10000
// 	formatting := fmt.Sprintf("%%.%df", w.Precision())
// 	return rightAlignColumnize(fmt.Sprintf(formatting, float64(w.Amount())/float64(unit)), w.SymbolName())

// }

func PrettifyTime(micro int64) string {
	value := float64(micro)
	unit := "μs"
	if value > 1000000*60*60 {
		value /= float64(1000000 * 60 * 60)
		unit = "h"
	} else if value > 1000000*60 {
		value /= float64(1000000 * 60)
		unit = "m"
	} else if value > 1000000 {
		value /= float64(1000000)
		unit = "s"
	} else if value > 1000 {
		value /= float64(1000)
		unit = "ms"
	}

	precision := 3
	if value >= 100 {
		precision = 1
	} else if value >= 10 {
		precision = 2
	}

	return rightAlignColumnize(strconv.FormatFloat(value, 'f', precision, 64), unit)
}

func PrettifyBytes(b int64) string {
	const u = 1024
	if b < u {
		return rightAlignColumnize(fmt.Sprintf("%d", b), "bytes")
	}
	div, exp := int64(u), 0
	for n := b / u; n >= u; n /= u {
		div *= u
		exp++
	}
	value := float64(b) / float64(div)
	unit := fmt.Sprintf("%cB", "KMGTPE"[exp])

	precision := 3
	if value >= 100 {
		precision = 1
	} else if value >= 10 {
		precision = 2
	}

	return rightAlignColumnize(strconv.FormatFloat(value, 'f', precision, 64), unit)
}
