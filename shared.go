package shared

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

// range specification, note that min <= max
type Int64Range struct {
	Min, Max int64
}

// get next random value within the interval including min and max
func (ir *Int64Range) NextRandom(r *rand.Rand) int64 {
	return r.Int63n(ir.Max-ir.Min+1) + ir.Min
}

// get next median value within the interval including min and max
func (ir *Int64Range) NextMedian() int64 {
	return (ir.Min + ir.Max) / 2
}

// func ToUint64Slice(in []int) []eos.Uint64 {
// 	out := make([]eos.Uint64, len(in))
// 	for idx, val := range in {
// 		out[idx] = eos.Uint64(val)
// 	}

// 	return out
// }

func GetMD5ShortHash(text string) string {
	hash := md5.Sum([]byte(text))
	h := hex.EncodeToString(hash[:])
	return h[:12]
}
