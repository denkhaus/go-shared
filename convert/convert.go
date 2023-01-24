package convert

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

func MustStringToFloat64(val string) float64 {
	if val == "" {
		return 0.0
	}

	ret, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic(errors.Errorf("unable to convert %s to float64", val))
	}
	return ret
}

func MustStringToInt32(val string) int {
	if val == "" {
		return 0
	}

	ret, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		panic(errors.Errorf("unable to convert %s to int32", val))
	}
	return int(ret)
}

func NormalizeFloat64(v float64, precision int) (float64, error) {
	unitsString := strconv.FormatFloat(v, 'f', precision, 64)
	return strconv.ParseFloat(unitsString, 64)
}

func MustNormalizeFloat64(v float64, precision int) float64 {
	normalized, err := NormalizeFloat64(v, precision)
	if err != nil {
		panic(fmt.Sprintf("normalize float64 %v", v))
	}

	return normalized
}

func InterfaceToStruct(in interface{}, out interface{}) error {
	if vString, ok := in.(string); ok {
		m, err := cast.ToStringMapE(vString)
		if err != nil {
			return errors.Wrap(err, "ToStringMapE")
		}
		in = m
	}

	buf, err := json.Marshal(in)
	if err != nil {
		return errors.Wrap(err, "Marshal")
	}

	if err := json.Unmarshal(buf, out); err != nil {
		return errors.Wrap(err, "Unmarshal")
	}

	return nil
}
